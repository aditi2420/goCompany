package models

import (
	//"github.com/gofrs/uuid"
	"github.com/google/uuid"
	//uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID          string `gorm:"primary_key"`
	Name        string `gorm:"not null;unique"`
	Description string
	Amount      uint   `gorm:"not null"`
	Registered  bool   `gorm:"not null"`
	Type        string `gorm:"not null"`
}

func (base *Company) BeforeCreate(tx *gorm.DB) error {
	//uuid := uuid.new(type)
	base.ID = uuid.NewString()
	//tx.Statement.SetColumn("ID", uuid)
	return nil
}

// Create table for `User`
//.Migrator().CreateTable(&User{})


