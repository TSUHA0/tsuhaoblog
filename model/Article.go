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
func GetArtList(pageSize int, pageNum int) ([]Article, int64,int) {
	var art []Article
	var total int64
	err = db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").
		Find(&art).Error
	// 单独计数
	db.Model(&art).Count(&total)
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return art,total, errmsg.SUCCSE
}

//搜索文章标题
func SearchArt(title string, pageSize int, pageNum int) ([]Article, int64, int) {
	var articleList []Article
	var err error
	var total int64
	err = db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").
		Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Joins("Category").
		Where("title LIKE ?", title+"%").Find(&articleList).Count(&total).Error

	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return articleList, total, errmsg.SUCCSE
}