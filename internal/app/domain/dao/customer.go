package dao

type Customer struct {
	BaseModel
	Name        *string    `json:"name"`
	Address     string     `gorm:"not null" json:"address"`
	Email       string     `gorm:"not null;unique" json:"email"`
	Password    string     `gorm:"not null;" json:"password,omitempty"`
	PhoneNumber string     `gorm:"not null" json:"phone_number"`
	Orders      []Order    `gorm:"foreignKey:CustomerID" json:"orders"`
	Payments    []Payment  `gorm:"foreignKey:CustomerID" json:"payments"`
	Carts       []Cart     `gorm:"foreignKey:CustomerID" json:"carts"`
	WishList    []WishList `gorm:"foreignKey:CustomerID" json:"wish_list"`
	Shipment    []Shipment `gorm:"foreignKey:CustomerID" json:"shipment"`
}
