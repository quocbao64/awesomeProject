package dao

type OrderItem struct {
	BaseModel
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
}
