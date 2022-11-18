package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "mgs",
	Short: "mgs是原神数据工具，基于米游社数据提供快速查询服务",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("现在还不能用")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, "执行错误", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(fetch)
	rootCmd.AddCommand(fix)
}

func initConfig() {
	viper.SetConfigFile("./config.yaml")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("加载配置文件：", viper.ConfigFileUsed())
	}
}
