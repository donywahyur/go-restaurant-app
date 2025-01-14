package database

import (
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"

	"gorm.io/gorm"
)

func dbSeed(db *gorm.DB) {
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})

	menu := []model.MenuItem{
		{Name: "Hamburger", OrderCode: "HAMBURGER", Price: 10000, Type: constant.MenuTypeFood},
		{Name: "Coke", OrderCode: "COKE", Price: 5000, Type: constant.MenuTypeDrink},
		{Name: "Ice Tea", OrderCode: "ICE-TEA", Price: 5000, Type: constant.MenuTypeDrink},
		{Name: "Pepsi", OrderCode: "PEPSI", Price: 5000, Type: constant.MenuTypeDrink},
		{Name: "Sprite", OrderCode: "SPRITE", Price: 5000, Type: constant.MenuTypeDrink},
		{Name: "Spagetti", OrderCode: "SPAGETTI", Price: 10000, Type: constant.MenuTypeFood},
		{Name: "Fried Chicken", OrderCode: "FRIED-CHICKEN", Price: 15000, Type: constant.MenuTypeFood},
		{Name: "Hot Dog", OrderCode: "HOT-DOG", Price: 10000, Type: constant.MenuTypeFood},
		{Name: "Fried Rice", OrderCode: "FRIED-RICE", Price: 10000, Type: constant.MenuTypeFood},
	}

	err := db.First(&model.MenuItem{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Create(&menu)
		}
	}
}
