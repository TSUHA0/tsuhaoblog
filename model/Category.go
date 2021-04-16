package model

import (
	"gorm.io/gorm"
	"tsuhaoblog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//添加分类
func AddCategory(data *Category) int {

	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}

//查询分类是否存在
func CheckCate(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

//查询单个分类
func GetCate(id int) (Category, int) {
	var cate Category
	err := db.Where("id=?", id).First(&cate).Error
	if err != nil {
		return cate, errmsg.ERROR
	}
	return cate, errmsg.SUCCSE
}

//查询分类列表
func GetCateList(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

//查询单个分类下的文章
func GetCateArt(Cid int, pageSize int, pageNum int) ([]Article, int64, int) {
	var arts []Article
	var totalcnt int64
	err := db.Preload("Category").Limit(pageSize).Offset(pageSize*(pageNum-1)).
		Where("Cid = ?", Cid).Find(&arts).Error
	db.Model(&arts).Where("Cid = ?", Cid).Count(&totalcnt)
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return arts, totalcnt, errmsg.SUCCSE
}

//编辑分类
func EditCate(id int, data *Category) int {
	var cate,other Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	db.Select("id").Where("name = ?", data.Name).First(&other)
	if other.ID >0 && other.ID != uint(id){
		return errmsg.ERROR_CATENAME_USED
	}
	err := db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除分类
func DeleteCate(id int) int {
	err := db.Where("id=?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
