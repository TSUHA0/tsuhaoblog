package model

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"tsuhaoblog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,gte=2"`
}

func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

func GetUser(id int) (User, int) {
	var user User
	err := db.Where("id=?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

func GetUserList(username string, pageSize int, pageNum int) ([]User, int64, int) {
	var users []User
	var total int64
	err := db.Model(User{}).Where("username LIKE ?", username+"%").Count(&total).
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, total, errmsg.ERROR
	}
	return users, total, errmsg.SUCCSE

}

func ScryptPW(password string) string {
	//func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error)
	keyLen := 10
	salt := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	fpw, err := scrypt.Key([]byte(password), salt, 2048, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(fpw)
}

func EditUser(id int, data *User) int {
	var user, other User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	db.Select("id").Where("username = ?", data.Username).First(&other)
	if other.ID > 0 && other.ID != uint(id) {
		return errmsg.ERROR_USERNAME_USED
	}
	err := db.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//更改密码
func EditPassword(id int, data *User) int {
	err := db.Select("password").Where("id = ?", id).Updates(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = ScryptPW(u.Password)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPW(u.Password)
	return nil
}

//后台登陆验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIT
	}
	if user.Role > 1 {
		return errmsg.ERROR_USER_NO_PERMISSION
	}
	if user.Password != ScryptPW(password) {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCSE
}
