package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type rootCmd struct {
	cobraCmd   *cobra.Command
	optionsCmd []*cobra.Command
}

func (rc *rootCmd) addCommands(cs ...*cobra.Command) {
	rc.optionsCmd = append(rc.optionsCmd, cs...)
	rc.cobraCmd.AddCommand(cs...)
}

// rootCmd 代表没有调用子命令时的基础命令
var rCmd = &rootCmd{
	cobraCmd: &cobra.Command{
		SilenceUsage: true,
		Short:        "object-uploader 是一个打包加密并上传对象存储的软件",
		Long:         `object-uploader 是一个打包加密并上传对象存储的软件`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Args: cobra.ExactArgs(0),
	},
}

// Execute 将所有子命令添加到root命令并适当设置标志。会被 main.main() 调用一次。
func Execute() {
	if err := rCmd.cobraCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// 去除自动补全
	rCmd.cobraCmd.CompletionOptions.DisableDefaultCmd = true
	rCmd.addCommands(versionCmd, uploaderCmd)
}
