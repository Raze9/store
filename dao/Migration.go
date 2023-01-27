package dao

import (
	"GOproject/GIT/model"
	"fmt"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Admin{},
		&model.Carousel{},
		&model.Category{},
		&model.Cart{},
		&model.Notice{},
		&model.Product{},
		&model.ProductImg{},
		&model.Favorite{},
		&model.Order{},
	)
	if err != nil {
		fmt.Println("err", err)
	}
	return
}
