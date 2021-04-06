package model

import (
	"gorm.io/gorm"
	"tsuhaoblog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

//新建文章
func AddArt(data *Article) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除文章
func DeleteArt(id int) int {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//修改文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["Title"] = data.Title
	maps["Cid"] = data.Cid
	maps["Desc"] = data.Desc
	maps["Content"] = data.Content
	maps["Img"] = data.Img
	maps["CommentCount"] = data.CommentCount
	maps["ReadCount"] = data.ReadCount

	err := db.Model(&art).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询单个文章
func GetArt(id int) (Article, int) {
	var art Article
	err := db.Where("id = ?", id).Find(&art).Error
	if err != nil {
		return art, errmsg.ERROR
	}
	return art, errmsg.SUCCSE
}

//查询文章列表
func GetArtList(pageSize int, pageNum int) ([]Article, int) {
	var art []Article

	err := db.Limit(pageSize).Offset(pageSize * (pageNum - 1)).Find(&art).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return art, errmsg.SUCCSE
}
