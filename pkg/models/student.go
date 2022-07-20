package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/minhtam3010/student/pkg/config"
)

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Major string `json:"major"`
}

var db *gorm.DB

var Students []Student

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{}, &Parent{}, &Student_Parent{})
}

func GetAllStudent() []Student {
	err := db.Find(&Students).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return Students
}

func GetStudentById(ID int64) *Student {
	var student Student
	fmt.Println(student)
	err := db.Where("ID=?", ID).Find(&student).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	return &student
}

func (s *Student) CreateStudent() *Student {
	db.NewRecord(s)

	err := db.Create(&s).Error
	if err != nil {
		log.Printf("Error problem with create record: %v\n", err)
	}
	return s
}

func UpdateStudent(studentDetails *Student) *Student {
	err := db.Save(&studentDetails).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	return studentDetails
}

func DeleteStudent(ID int64) Student {
	var student Student
	err := db.Where("ID=?", ID).Delete(student).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	return student
}
