package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tsuhaoblog/model"
	"tsuhaoblog/utils/errmsg"
)

//添加文章
func AddArt(c *gin.Context) {
	var art model.Article
	c.ShouldBindJSON(&art)
	code = model.AddArt(&art)
	c.JSON(http.StatusOK, gin.H{
		"data":    art,
		"status":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个文章
func GetArt(c *gin.Context) {
	var art model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	art, code = model.GetArt(id)
	c.JSON(http.StatusOK, gin.H{
		"data":    art,
		"status":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArtList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	art,total,code:=model.GetArtList(pageSize,pageNum)
	c.JSON(http.StatusOK, gin.H{
		"data":    art,
		"status":    code,
		"total":	total,
		"message": errmsg.GetErrMsg(code),
	})

}

//编辑文章
func EditArt(c *gin.Context) {
	var art model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&art)
	code = model.EditArt(id, &art)
	c.JSON(http.StatusOK, gin.H{
		"data":    art,
		"status":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章z
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
