package upload

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"tsuhaoblog/utils"
	"tsuhaoblog/utils/errmsg"
)

func Upload(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: utils.Bucket,
	}
	mac := qbox.NewMac(utils.AccessKey, utils.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	putExtra := storage.PutExtra{}

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", errmsg.ERROR
	}
	return ret.Key, errmsg.SUCCSE
}
