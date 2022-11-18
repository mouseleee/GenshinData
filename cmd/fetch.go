package cmd

import "github.com/spf13/cobra"

var fetch = &cobra.Command{
	Use: "fetch",
	Run: func(cmd *cobra.Command, args []string) {
		println("这个现在也不能用")
	},
}
