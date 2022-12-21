package service

import (
	"object-uploader/cos"
	"object-uploader/utils"
	"os"

	"github.com/spf13/viper"
)

// Upload 上传总体逻辑
func Upload() error {
	var (
		src         = viper.GetString("src")
		cosDestName = viper.GetString("dest")
	)

	dest, err := os.CreateTemp("", "object-uploader-*.tar.gz")
	if err != nil {
		return err
	}
	defer dest.Close()
	defer os.Remove(dest.Name())

	// 将需要上传的对象加密打包
	if err := utils.Compress(src, dest); err != nil {
		return err
	}

	// 上传 cos
	return cos.Upload(cosDestName+".tar.gz", dest.Name())
}
