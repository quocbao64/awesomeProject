package repository

import (
	"awesomeProject/internal/app/domain/dao"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll(limit int, offset int) ([]dao.Category, error)
	GetByID(id int) (dao.Category, error)
	Create(category dao.Category) (dao.Category, error)
	Update(category dao.Category) (dao.Category, error)
	Delete(id int) error
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func (cat CategoryRepositoryImpl) GetAll(limit int, offset int) ([]dao.Category, error) {
	var categories []dao.Category
	err := cat.db.Limit(limit).Offset(offset).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cat CategoryRepositoryImpl) GetByID(id int) (dao.Category, error) {
	var category dao.Category
	err := cat.db.First(&category, id).Error

	if err != nil {
		return dao.Category{}, err
	}

	return category, nil
}

func (cat CategoryRepositoryImpl) Create(category dao.Category) (dao.Category, error) {
	err := cat.db.Create(&category).Error

	if err != nil {
		return dao.Category{}, err
	}

	return category, nil
}

func (cat CategoryRepositoryImpl) Update(category dao.Category) (dao.Category, error) {
	err := cat.db.Save(&category).Error

	if err != nil {
		return dao.Category{}, err
	}

	return category, nil
}

func (cat CategoryRepositoryImpl) Delete(id int) error {
	var category dao.Category
	err := cat.db.Delete(&category, id).Error

	if err != nil {
		return err
	}

	return nil
}

func CategoryRepositoryInit(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
		db: db,
	}
}
