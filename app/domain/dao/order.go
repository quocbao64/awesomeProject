package dao

import "time"

type Order struct {
	BaseModel
	OrderDate  time.Time `gorm:"not null" json:"order_date"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
	CustomerID uint      `gorm:"not null" json:"customer_id"`
	OrderItems []OrderItem
	ShipmentID uint `json:"shipment_id"`
	PaymentID  uint `json:"payment_id"`
}
