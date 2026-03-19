package thirdpart

import (
	"crypto/tls"
	"net/http"
	"time"
)

// sharedTransport 全局共享HTTP传输层
// 禁用HTTP/2 + DNS缓存拨号器
// HTTP/2会将所有并发请求复用到同一TCP连接（只到达1台服务器），HTTP/1.1下每个并发请求使用独立连接，分散到不同后端IP
var SharedTransport = &http.Transport{
	ForceAttemptHTTP2: false,
	TLSNextProto:      make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),

	MaxIdleConns:          500,
	MaxIdleConnsPerHost:   100,
	MaxConnsPerHost:       500,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	DisableCompression:    true,

	DialContext: Global.DialContext,
}
