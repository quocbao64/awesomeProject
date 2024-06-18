package repository

import (
	"awesomeProject/app/domain/dao"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetAll() ([]dao.Customer, error)
	GetByID(id int) (dao.Customer, error)
	GetByEmail(email string) (dao.Customer, error)
	Create(customer dao.Customer) (dao.Customer, error)
	Update(customer dao.Customer) (dao.Customer, error)
	Delete(id int) error
}

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func (cus CustomerRepositoryImpl) GetAll() ([]dao.Customer, error) {
	var customers []dao.Customer
	err := cus.db.Find(&customers).Error

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (cus CustomerRepositoryImpl) GetByID(id int) (dao.Customer, error) {
	var customer dao.Customer
	err := cus.db.First(&customer, id).Error

	if err != nil {
		return dao.Customer{}, err
	}

	return customer, nil
}

func (cus CustomerRepositoryImpl) GetByEmail(email string) (dao.Customer, error) {
	var customer dao.Customer
	err := cus.db.Where("email = ?", email).First(&customer).Error

	if err != nil {
		return dao.Customer{}, err
	}

	return customer, nil
}

func (cus CustomerRepositoryImpl) Create(customer dao.Customer) (dao.Customer, error) {
	err := cus.db.Create(&customer).Error

	if err != nil {
		return dao.Customer{}, err
	}

	return customer, nil
}

func (cus CustomerRepositoryImpl) Update(customer dao.Customer) (dao.Customer, error) {
	err := cus.db.Save(&customer).Error

	if err != nil {
		return dao.Customer{}, err
	}

	return customer, nil
}

func (cus CustomerRepositoryImpl) Delete(id int) error {
	err := cus.db.Delete(&dao.Customer{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func CustomerRepositoryInit(db *gorm.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{
		db: db,
	}
}
