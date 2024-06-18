package dao

import "time"

type Payment struct {
	BaseModel
	PaymentDate   time.Time `gorm:"not null" json:"payment_date"`
	PaymentMethod string    `gorm:"not null" json:"payment_method"`
	Amount        float64   `gorm:"not null" json:"amount"`
	CustomerID    uint      `gorm:"not null" json:"customer_id"`
	Orders        []Order   `gorm:"foreignKey:PaymentID" json:"orders"`
}
