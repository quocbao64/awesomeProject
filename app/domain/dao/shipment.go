package dao

import "time"

type Shipment struct {
	BaseModel
	ShipmentDate time.Time `gorm:"not null" json:"shipment_date"`
	Address      string    `gorm:"not null" json:"address"`
	City         string    `gorm:"not null" json:"city"`
	State        string    `gorm:"not null" json:"state"`
	Country      string    `gorm:"not null" json:"country"`
	ZipCode      string    `gorm:"not null" json:"zip_code"`
	CustomerID   uint      `gorm:"not null" json:"customer_id"`
	Orders       []Order   `gorm:"foreignKey:ShipmentID" json:"orders"`
}
