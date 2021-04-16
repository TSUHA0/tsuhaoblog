package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tsuhaoblog/model"
	"tsuhaoblog/utils/errmsg"
)

//添加分类
func AddCate(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCSE {
		model.AddCategory(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类
func GetCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cate, code := model.GetCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":    code,
		"message": errmsg.GetErrMsg(code),
		"data":    cate,
	})
}

//查询单个分类下的文章
func GetCateArt(c *gin.Context) {
	var art []model.Article
	var total int64
	cid, _ := strconv.Atoi(c.Param("id"))
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

	art, total, code = model.GetCateArt(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"data":    art,
		"status":    code,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

//查询分类列表
func GetCateList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetCateList(pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑分类
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.EditCate(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

//删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
