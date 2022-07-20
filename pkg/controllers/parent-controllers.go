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


func GetAllParent(w http.ResponseWriter, r *http.Request) {
	parents := models.GetAllParent()
	if body, err := json.Marshal(parents); err == nil {
		WriteResponse(w, body)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func GetParentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parentId := vars["parentId"]
	ID, err := strconv.ParseInt(parentId, 0, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	parentDetails := models.GetParentById(ID)
	if res, err := json.Marshal(parentDetails); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func CreateParent(w http.ResponseWriter, r *http.Request) {
	createParent := &models.Parent{}
	utils.ParseBody(r, createParent)
	parent := createParent.CreateParent()
	if res, err := json.Marshal(parent); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func UpdateParentInfo(parentDetails *models.Parent, updateParent *models.Parent) {
	if updateParent.Name != "" {
		parentDetails.Name = updateParent.Name
	}
	if updateParent.Age > 36{
		parentDetails.Age = updateParent.Age
	}
}

func UpdateParent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parentId := vars["parentId"]
	ID, err := strconv.ParseInt(parentId, 0, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	updateParent := &models.Parent{}
	utils.ParseBody(r, updateParent)
	parentDetails := models.GetParentById(ID)
	UpdateParentInfo(parentDetails, updateParent)
	parent := models.UpdateParent(parentDetails)
	if res, err := json.Marshal(parent); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func DeleteParent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parentId := vars["parentId"]
	ID, err := strconv.ParseInt(parentId, 0, 0)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	parent := models.DeleteParent(ID)
	if res, err := json.Marshal(parent); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}
