package Models

import "github.com/jinzhu/gorm"

type Topic struct {
	gorm.Model
	Reply      []Reply
	Title      string
	Body       string `gorm:"type:text"`
	UserId     int64
	CategoryId int64
}
