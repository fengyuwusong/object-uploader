package logic

import (
	"object-uploader/utils"
	"os"

	"github.com/spf13/viper"
)

func Compress(src string, dest *os.File) error {
	// TODO 判断是否需要加密
	if viper.GetBool("encrypt.use") {

	}

	// 压缩 TODO: 后续支持多种压缩类型
	return utils.Compress(src, dest)
}
