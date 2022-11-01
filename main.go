package main

import (
	"os"
	"time"

	"github.com/go-rod/rod"
	mlib "github.com/mouseleee/mouselib"
	"github.com/rs/zerolog"
	"golang.org/x/net/html"
)

const browserBinPath = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"

var log zerolog.Logger

type LocalBrowser struct {
	Browser *rod.Browser
}

func MatchAttr(o []html.Attribute, m []html.Attribute) bool {
	for _, attr := range m {
		match := false
		for _, cmp := range o {
			if cmp.Namespace == attr.Namespace && cmp.Key == attr.Key && cmp.Val == attr.Val {
				match = true
				break
			}
		}
		if !match {
			return match
		}
	}
	return true
}

func Find(n *html.Node, t html.NodeType, ele string, attrs []html.Attribute) []*html.Node {
	r := make([]*html.Node, 0)
	var cond bool
	switch t {
	case html.ElementNode:
		cond = n.Type == t && n.Data == ele && MatchAttr(n.Attr, attrs)
	case html.TextNode:
		cond = n.Type == t
	default:
		cond = false
	}

	if cond {
		r = append(r, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r = append(r, Find(c, t, ele, attrs)...)
	}
	return r
}

func NewLocalBrowser(b *rod.Browser) *LocalBrowser {
	return &LocalBrowser{
		Browser: b,
	}
}

func main() {
	var err error
	log, err = mlib.CommandLogger("debug")
	if err != nil {
		os.Exit(1)
	}

	// u := launcher.New().Bin(browserBinPath).MustLaunch()
	// browser := rod.New().ControlURL(u).MustConnect()
	// lb := NewLocalBrowser(browser)
	// defer lb.Browser.MustClose()

	// roles := "https://bbs.mihoyo.com/ys/obc/channel/map/189/25?bbs_presentation_style=no_header"
	// lb.SavePage(roles)
	FetchRole("./assets/html/roles.html")
}

func (b *LocalBrowser) SavePage(url string) {
	p := b.Browser.MustPage(url).MustWaitLoad()
	wait := p.WaitRequestIdle(3*time.Second, []string{}, []string{})
	wait()
	data := p.MustHTML()
	mlib.WriteFile("./assets/html/roles.html", []byte(data))
}

func FetchRole(p string) {
	f, err := os.Open(p)
	if err != nil {
		log.Err(err).Msg("打开html文件错误")
	}
	doc, err := html.Parse(f)
	if err != nil {
		log.Err(err).Msg("转换html文档树失败")
	}
	eles := Find(doc, html.ElementNode, "div", []html.Attribute{
		{
			Key: "class",
			Val: "collection-avatar",
		},
	})
	log.Info().Int("len", len(eles)).Msg("search result")

	as := Find(eles[1], html.ElementNode, "a", []html.Attribute{
		{
			Key: "class",
			Val: "collection-avatar__item",
		},
	})
	log.Info().Int("len", len(as)).Msg("search result")

	for _, a := range as {
		t := Find(a, html.TextNode, "", nil)
		log.Info().Int("len", len(t)).Str("text1", t[0].Data).Str("text2", t[1].Data).Msg("search result")
	}
}

// <a data-v-51c84696="" href="/ys/obc/content/4781/detail?bbs_presentation_style=no_header" target="_blank" class="collection-avatar__item"><div data-v-51c84696="" class="collection-avatar__icon" data-src="https://uploadstatic.mihoyo.com/ys-obc/2022/09/20/4328207/0587df6ac5144c9dd023b0f73ceaf8be_7837788427682725710.png?x-oss-process=image/quality,q_75/resize,s_120" lazy="loaded" style="background-image: url(&quot;https://uploadstatic.mihoyo.com/ys-obc/2022/09/20/4328207/0587df6ac5144c9dd023b0f73ceaf8be_7837788427682725710.png?x-oss-process=image/quality,q_75/resize,s_120&quot;);"><div data-v-51c84696="" class="red-point"><!----></div></div> <div data-v-51c84696="" class="collection-avatar__title">坎蒂丝</div></a>
