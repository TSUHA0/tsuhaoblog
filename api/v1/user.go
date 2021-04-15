package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tsuhaoblog/model"
	"tsuhaoblog/utils/errmsg"
	"tsuhaoblog/utils/validator"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var vds string
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println("ShouldBindJSON Error")
		return
	}

	vds, code = validator.Validate(data)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"error":   vds,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, code := model.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":    code,
		"message": errmsg.GetErrMsg(code),
		"data":    user,
	})
}

//查询用户列表
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetUserList(username,pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

//编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code = model.EditUser(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"data":data,
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//更改密码
func EditPassword(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&user)

	code = model.EditPassword(id, &user)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

