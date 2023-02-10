package models

import (
	"errors"
	//"gorm.io/gorm"
)

func (db dbHandle) GetCompany(name string) (*Company, error) {

	if name == "" {
		return nil, errors.New("Empty company name")
	}

	var dbCompany Company
	err := db.First(&dbCompany,"name = ?", name).Error
	if err != nil {
		return nil, errors.New("Failed to get company details")
	}
	return &dbCompany, nil
}

func (db dbHandle) CreateCompany(c1 *Company) (*Company, error) {

	err := db.Create(c1).Error
	if err != nil {
		return nil, errors.New("Failed to create company details")
	}

	return c1, nil
}



func (db dbHandle) DeleteCompany(name string) ( error) {

	if name == "" {
		return  errors.New("Empty company name")
	}

	var dbCompany Company
	err := db.Delete(&dbCompany,"name = ?", name).Error
	if err != nil {
		return  errors.New("Failed to delete company details")
	}
	return  nil
}

func (db dbHandle) UpdateCompany(c1 *Company) (*Company, error) {

	err := db.Updates(c1).Error
	if err != nil {
		return nil, errors.New("Failed to update company details")
	}

	return c1, nil
}
