package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO 后续自动构建变量载入
		fmt.Println("object uploader version: 0.0.1\nauthor: fengyuwusong\ndate: 2022-12-21 16:15:35")
	},
}
