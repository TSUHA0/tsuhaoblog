package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tsuhaoblog/middleware"
	"tsuhaoblog/model"
	"tsuhaoblog/utils/errmsg"
)

func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	code = model.CheckLogin(user.Username, user.Password)
	if code == errmsg.SUCCSE {
		setToken(c, user)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    user.Username,
			"id":      user.ID,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "tsuhao",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.SUCCSE,
			"message": errmsg.GetErrMsg(errmsg.SUCCSE),
			"token":   token,
		})
	}

}
