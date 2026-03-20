package thirdpart

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"net"
	"net/url"
	"sync"
	"time"
)

// ---------------------------------------------------------------------------
// DNS 缓存层
// 预解析并缓存目标主机的所有IP地址，消除关键请求时刻的DNS查询延迟（5-50ms）
// 配合自定义DialContext，每次TCP连接随机选择一个IP，实现客户端侧负载均衡
// ---------------------------------------------------------------------------

// DNSCache DNS解析结果缓存，线程安全
type DNSCache struct {
	mu    sync.RWMutex
	cache map[string][]string
}

// Global 全局DNS缓存实例，所有transport共享
var Global = &DNSCache{
	cache: make(map[string][]string),
}

var defaultDialer = &net.Dialer{
	Timeout:   5 * time.Second,
	KeepAlive: 30 * time.Second,
}

// PreResolve 预解析主机名并缓存所有IP地址
// 应在初始化阶段调用，确保关键时刻DNS结果已就绪
func (d *DNSCache) PreResolve(host string) []string {
	addrs, err := net.LookupHost(host)
	if err != nil || len(addrs) == 0 {
		slog.Warn("[DNSCache] resolve failed", "host", host, "error", err)
		return nil
	}
	var v4addrs []string
	for _, addr := range addrs {
		if net.ParseIP(addr).To4() != nil {
			v4addrs = append(v4addrs, addr)
		}
	}
	if len(v4addrs) == 0 {
		slog.Warn("[DNSCache] no IPv4 address found", "host", host, "addrs", addrs)
		return nil
	}
	d.mu.Lock()
	d.cache[host] = v4addrs
	d.mu.Unlock()
	slog.Debug("[DNSCache]", "host", host, "addrs", v4addrs, "count", len(v4addrs))
	return v4addrs
}

// Resolve 获取缓存IP列表，缓存未命中则触发预解析
func (d *DNSCache) Resolve(host string) []string {
	d.mu.RLock()
	addrs := d.cache[host]
	d.mu.RUnlock()
	if len(addrs) > 0 {
		return addrs
	}
	return d.PreResolve(host)
}

// CachedIPCount 返回指定主机的已缓存IP数量
func (d *DNSCache) CachedIPCount(host string) int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return len(d.cache[host])
}

// DialContext 自定义拨号器，使用缓存的DNS结果直连IP
//
// 工作原理：
//  1. 从缓存中获取目标主机的所有IP
//  2. 随机选择一个IP建立TCP连接（客户端侧负载均衡）
//  3. 配合HTTP/1.1独立连接模式，每个并发请求可能被路由到不同后端服务器
//
// 这在目标服务器使用DNS轮询负载均衡时尤为关键：
// 并发N个请求 × M个后端IP = 最多N×M种竞争路径
func (d *DNSCache) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return defaultDialer.DialContext(ctx, network, addr)
	}
	addrs := d.Resolve(host)
	if len(addrs) == 0 {
		return defaultDialer.DialContext(ctx, network, addr)
	}
	// 随机选择IP，确保并发请求分散到不同后端
	ip := addrs[rand.IntN(len(addrs))]
	return defaultDialer.DialContext(ctx, network, net.JoinHostPort(ip, port))
}

// PreResolveURL 从URL中提取主机名并预解析DNS
// 返回解析到的所有IP地址
func PreResolveURL(rawURL string) []string {
	u, err := url.Parse(rawURL)
	if err != nil || u.Host == "" {
		return nil
	}
	return Global.PreResolve(u.Hostname())
}
