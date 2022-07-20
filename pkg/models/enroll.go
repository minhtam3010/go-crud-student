package models

import "log"

type Enroll struct {
	Student Student `json:"student"`
	Parent  Parent  `json:"parent"`
}

type Student_Parent struct {
	Student_id uint `json:"student_id"`
	Parent_id  uint `json:"parent_id"`
}

func (e *Enroll) CreateEnroll() *Enroll {
	db.NewRecord(e.Student)
	db.NewRecord(e.Parent)

	err1 := db.Create(&e.Student).Error
	err2 := db.Create(&e.Parent).Error
	if err1 != nil || err2 != nil {
		log.Printf("Error problem with create record: %v or %v\n", err1, err2)
	}
	return e
}

func (sp *Student_Parent) CreateStudentParent() {
	db.NewRecord(sp)
	err := db.Create(&sp).Error
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
