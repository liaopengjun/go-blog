package models

import (
	"gin-blog/global"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	global.Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate GORM 回调方法
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// BeforeUpdate GORM 回调方法
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps map[string]interface{}) (tags []Tag) {
	global.Db.Model(&Tag{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	global.Db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func GetTagInfo(name string, id int) (tag *Tag, err error) {
	db := global.Db
	if name != "" {
		db.Where(" name =? ", name)
	}
	if id > 0 {
		db.Where(" id =? ", id)
	}
	err = db.First(&tag).Error
	return
}

func AddTag(name string, state int, createdBy string) bool {
	result := global.Db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) error {
	return global.Db.Where("id = ?", id).Delete(&Tag{}).Error
}

func EditTag(id int, data interface{}) error {
	return global.Db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
}
