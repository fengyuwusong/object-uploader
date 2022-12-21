package cos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"object-uploader/utils"
	"time"

	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func Upload(destPath, srcPath string) error {
	bucket := utils.MustString(viper.GetString("cos.bucket"))
	region := utils.MustString(viper.GetString("cos.region"))
	appID := utils.MustString(viper.GetString("cos.appID"))
	secretKey := utils.MustString(viper.GetString("cos.secretKey"))
	rawURL := fmt.Sprintf("http://%s.cos.%s.myqcloud.com", bucket, region)
	u, _ := url.Parse(rawURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  appID,
			SecretKey: secretKey,
		},
	})
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType:              "text/html",
			XCosServerSideEncryption: "AES256",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{},
	}
	_, err := c.Object.PutFromFile(context.Background(), destPath, srcPath, opt)
	return err
}
