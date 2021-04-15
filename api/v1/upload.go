package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tsuhaoblog/upload"
	"tsuhaoblog/utils/errmsg"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"msg":  errmsg.GetErrMsg(errmsg.ERROR),
		})
		return
	}
	fileSize := fileHeader.Size
	url, code := upload.Upload(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":  errmsg.GetErrMsg(code),
		"url":  url,
	})
}
