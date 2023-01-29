package model

import "github.com/jinzhu/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductId uint `gorm:"not null"`
}
