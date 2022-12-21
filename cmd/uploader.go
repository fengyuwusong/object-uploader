package cmd

import (
	"fmt"
	"object-uploader/config"
	"object-uploader/service"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var src, dest, cfgFile string

var uploaderCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传对象",
	Run: func(cmd *cobra.Command, args []string) {
		// 初始化配置
		viper.SetConfigFile(cfgFile)
		config.InitConfig()
		if err := service.Upload(); err != nil {
			fmt.Println("service.Upload error:", err)
			os.Exit(-1)
		}
	},
}

func init() {
	// TODO: 增加顺序
	uploaderCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.object-uploader.yaml)")
	uploaderCmd.PersistentFlags().StringVar(&src, "src", "", "需要被上传的路径")
	uploaderCmd.PersistentFlags().StringVar(&dest, "dest", "", "需要上传的对象存储地址")
	viper.BindPFlag("src", uploaderCmd.PersistentFlags().Lookup("src"))
	viper.BindPFlag("dest", uploaderCmd.PersistentFlags().Lookup("dest"))
}
