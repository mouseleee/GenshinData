package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fix = &cobra.Command{
	Use:   "fix",
	Short: "仅限开发使用",
	Run: func(cmd *cobra.Command, args []string) {
		// db.InitConn(viper.GetString("data"))
		// defer db.CloseConn()
		FetchPage()
	},
}

func FetchPage() {
	u := launcher.New().Bin(viper.GetString("browser")).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose()

	url := "https://bbs.mihoyo.com/ys/obc/channel/map/189/25?bbs_presentation_style=no_header"
	p := browser.MustPage(url).MustWaitLoad()
	wait := p.WaitRequestIdle(3*time.Second, []string{}, []string{})
	wait()
	data := p.MustHTML()

	f, err := os.OpenFile("./assets/html/roles.html", os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, "保存角色索引，打开文件错误", err)
		return
	}

	_, err = io.WriteString(f, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "保存角色索引页面错误", err)
		return
	}
}

// func ParseRoleIndex() {
// 	f, _ := os.Open("assets/html/roles.html")
// 	n, err := html.Parse(f)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "解析html失败", err)
// 	}

// }

func init() {
	fix.Flags().String("data", "./gs.db", "数据库存储路径")
	fix.Flags().String("browser", "", "浏览器路径")
	viper.BindPFlag("data", fix.Flags().Lookup("data"))
	viper.BindPFlag("browser", fix.Flags().Lookup("browser"))
}
