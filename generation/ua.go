package generation

import (
	"fmt"
	"hash/fnv"
	"math/rand"
)

// BrowserProfile 模拟真实浏览器的完整指纹
// Chrome 110+ 实施了 UA Reduction，UA 字符串中的 OS 版本被冻结为固定值
// 但同时引入了 Client Hints (sec-ch-ua-*) 作为替代，这些头在每个请求中自动发送
// 缺少 Client Hints 的 "Chrome" 请求是最明显的机器人特征之一
type BrowserProfile struct {
	UserAgent       string // 完整UA字符串
	SecChUa         string // sec-ch-ua: 浏览器品牌和版本
	SecChUaMobile   string // sec-ch-ua-mobile: 是否移动端
	SecChUaPlatform string // sec-ch-ua-platform: 操作系统平台
	AcceptLanguage  string // accept-language: 语言偏好
}

// chromeVersions 近期 Chrome 版本池，用于生成多样化的浏览器指纹
// 避免所有实例使用完全相同的版本号
var chromeVersions = []int{143, 144, 145, 146}

// platforms 平台配置：UA中的OS字符串 + Client Hints中的平台名
// Chrome UA Reduction 后，macOS 固定为 10_15_7，Windows 固定为 NT 10.0
var platforms = []struct {
	uaOS       string // UA 中的 OS 部分
	chPlatform string // sec-ch-ua-platform 值
}{
	{"Macintosh; Intel Mac OS X 10_15_7", `"macOS"`},
	{"Windows NT 10.0; Win64; x64", `"Windows"`},
}

// languages 语言偏好池，针对韩国平台优化
// 使用韩语为主、英语为辅的组合，更符合 Weverse 用户画像
var languages = []string{
	"ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7",
	"ko-KR,ko;q=0.9,en;q=0.8",
	"ko,en-US;q=0.9,en;q=0.8",
	"en-US,en;q=0.9,ko-KR;q=0.8,ko;q=0.7",
}

// notABrandVariants Not-A-Brand 标签的变体，Chrome 每个大版本会轮换这个值
var notABrandVariants = []string{
	`"Not-A.Brand";v="99"`,
	`"Not/A)Brand";v="8"`,
	`"Not)A;Brand";v="99"`,
	`"Not:A-Brand";v="8"`,
}

// RandomBrowserProfile 生成一个随机但内部一致的浏览器指纹
// 每个实例应在初始化时调用一次，之后复用同一个 Profile
// 保证同一实例的所有请求指纹一致（真实浏览器就是这样）
func RandomBrowserProfile() BrowserProfile {
	ver := chromeVersions[rand.Intn(len(chromeVersions))]
	plat := platforms[rand.Intn(len(platforms))]
	lang := languages[rand.Intn(len(languages))]
	notABrand := notABrandVariants[ver%len(notABrandVariants)]

	return BrowserProfile{
		UserAgent: fmt.Sprintf(
			"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36",
			plat.uaOS, ver,
		),
		SecChUa: fmt.Sprintf(
			`"Chromium";v="%d", "Google Chrome";v="%d", %s`,
			ver, ver, notABrand,
		),
		SecChUaMobile:   "?0",
		SecChUaPlatform: plat.chPlatform,
		AcceptLanguage:  lang,
	}
}

func RandomBrowserProfileWithSalt(salt string) BrowserProfile {
	h := fnv.New64a()
	h.Write([]byte(salt))
	seed := int64(h.Sum64())

	var rng = rand.New(rand.NewSource(seed))
	ver := chromeVersions[rng.Intn(len(chromeVersions))]
	plat := platforms[rng.Intn(len(platforms))]
	lang := languages[rng.Intn(len(languages))]
	notABrand := notABrandVariants[ver%len(notABrandVariants)]

	return BrowserProfile{
		UserAgent: fmt.Sprintf(
			"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36",
			plat.uaOS, ver,
		),
		SecChUa: fmt.Sprintf(
			`"Chromium";v="%d", "Google Chrome";v="%d", %s`,
			ver, ver, notABrand,
		),
		SecChUaMobile:   "?0",
		SecChUaPlatform: plat.chPlatform,
		AcceptLanguage:  lang,
	}
}
