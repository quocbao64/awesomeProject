package dao

type WishList struct {
	BaseModel
	ProductID  uint `gorm:"not null" json:"product_id"`
	CustomerID uint `gorm:"not null" json:"customer_id"`
}
