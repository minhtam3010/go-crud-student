package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/minhtam3010/student/pkg/models"
	"github.com/minhtam3010/student/pkg/utils"
)

func CreateEnroll(w http.ResponseWriter, r *http.Request) {
	createEnroll := &models.Enroll{}
	utils.ParseBody(r, createEnroll)
	enroll := createEnroll.CreateEnroll()

	// Create struct Student_Parent to take the student_id also parent_id
	student_parent := &models.Student_Parent{}
	student_parent.Student_id  = enroll.Student.ID
	student_parent.Parent_id = enroll.Parent.ID

	// Define function to fetch the student_id and also parent_id
	student_parent.CreateStudentParent()

	if res, err := json.Marshal(enroll); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}
