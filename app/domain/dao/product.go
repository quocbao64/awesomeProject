package dao

type Product struct {
	BaseModel
	SKU         string      `gorm:"unique_index"`
	Description string      `gorm:"not null" json:"description"`
	Price       float64     `gorm:"not null" json:"price"`
	Stock       int         `gorm:"not null" json:"stock"`
	Carts       []Cart      `gorm:"foreignKey:ProductID" json:"carts"`
	OrderItems  []OrderItem `gorm:"foreignKey:ProductID" json:"order_items"`
	WishList    []WishList  `gorm:"foreignKey:ProductID" json:"wish_list"`
	CategoryID  uint        `gorm:"not null" json:"category_id"`
}
