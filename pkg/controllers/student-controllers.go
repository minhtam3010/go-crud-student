package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/minhtam3010/student/pkg/models"
	"github.com/minhtam3010/student/pkg/utils"
)

func WriteResponse(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	students := models.GetAllStudent()
	if body, err := json.Marshal(students); err == nil {
		WriteResponse(w, body)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	studentDetails := models.GetStudentById(ID)
	if res, err := json.Marshal(studentDetails); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	createStudent := &models.Student{}
	utils.ParseBody(r, createStudent)
	student := createStudent.CreateStudent()
	if res, err := json.Marshal(student); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func Update(studentDetails *models.Student, updateStudent *models.Student) {
	if updateStudent.Name != "" {
		studentDetails.Name = updateStudent.Name
	}
	if updateStudent.Year > 0 && updateStudent.Year <= 4 {
		studentDetails.Year = updateStudent.Year
	}
	if updateStudent.Major != "" {
		studentDetails.Major = updateStudent.Major
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	updateStudent := &models.Student{}
	utils.ParseBody(r, updateStudent)
	studentDetails := models.GetStudentById(ID)
	Update(studentDetails, updateStudent)
	student := models.UpdateStudent(studentDetails)
	if res, err := json.Marshal(student); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	student := models.DeleteStudent(ID)
	if res, err := json.Marshal(student); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}
