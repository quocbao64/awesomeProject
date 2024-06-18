package dao

type Cart struct {
	ID         uint `gorm:"primary_key" json:"id"`
	Quantity   int  `gorm:"not null" json:"quantity"`
	CustomerID uint `gorm:"not null" json:"customer_id"`
	ProductID  uint `gorm:"not null" json:"product_id"`
}
