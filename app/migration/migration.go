package migration

import (
	"awesomeProject/app/domain/dao"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&dao.Customer{},
		&dao.Shipment{},
		&dao.Payment{},
		&dao.Category{},
		&dao.Order{},
		&dao.Product{},
		&dao.OrderItem{},
		&dao.Cart{},
		&dao.WishList{})
	if err != nil {
		return
	}
}
