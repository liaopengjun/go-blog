package global

import (
	"gin-blog/common/config"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}
