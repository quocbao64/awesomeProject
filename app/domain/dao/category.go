package dao

type Category struct {
	BaseModel
	Name     string    `gorm:"not null" json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products"`
}
