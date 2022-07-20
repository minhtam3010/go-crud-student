package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Parent struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Parents []Parent

func GetAllParent() []Parent {
	err := db.Find(&Parents).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return Parents
}

func GetParentById(ID int64) *Parent {
	var parent Parent
	err := db.Where("ID=?", ID).Find(&parent).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	return &parent
}

func (p *Parent) CreateParent() *Parent {
	db.NewRecord(p)

	err := db.Create(&p).Error
	if err != nil {
		log.Printf("Error problem with create record: %v\n", err)
	}
	return p
}

func UpdateParent(parentDetails *Parent) *Parent {
	err := db.Save(&parentDetails).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return parentDetails
}

func DeleteParent(ID int64) Parent {
	var parent Parent
	err := db.Where("ID=?", ID).Delete(parent).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	return parent
}
