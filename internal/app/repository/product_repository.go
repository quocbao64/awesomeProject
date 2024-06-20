package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(limit int, offset int) ([]dao.Product, error)
	GetByID(id int) (dao.Product, error)
	Create(product dao.Product) (dao.Product, error)
	Update(product dao.Product) (dao.Product, error)
	Delete(id int) error
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func (p ProductRepositoryImpl) GetAll(limit int, offset int) ([]dao.Product, error) {
	var products []dao.Product
	err := p.db.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (p ProductRepositoryImpl) GetByID(id int) (dao.Product, error) {
	var product dao.Product
	err := p.db.First(&product, id).Error
	return product, err
}

func (p ProductRepositoryImpl) Create(product dao.Product) (dao.Product, error) {
	err := p.db.Create(&product).Error
	return product, err
}

func (p ProductRepositoryImpl) Update(product dao.Product) (dao.Product, error) {
	err := p.db.Save(&product).Error
	return product, err
}

func (p ProductRepositoryImpl) Delete(id int) error {
	return p.db.Delete(&dao.Product{}, id).Error
}

func ProductRepositoryInit(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}
