package routes

import (
	"github.com/gorilla/mux"
	"github.com/minhtam3010/student/pkg/controllers"
)

var RegisterStudentRoutes = func(routes *mux.Router) {

	// API for student
	routes.HandleFunc("/student/", controllers.GetAllStudent).Methods("GET")
	routes.HandleFunc("/student/{studentId}", controllers.GetStudentById).Methods("GET")
	routes.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")
	routes.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")
	routes.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")

	// API for parent
	routes.HandleFunc("/parent/", controllers.GetAllParent).Methods("GET")
	routes.HandleFunc("/parent/{parentId}", controllers.GetParentById).Methods("GET")
	routes.HandleFunc("/parent/", controllers.CreateParent).Methods("POST")
	routes.HandleFunc("/parent/{parentId}", controllers.UpdateParent).Methods("PUT")
	routes.HandleFunc("/parent/{parentId}", controllers.DeleteParent).Methods("DELETE")

	// API for enroll
	routes.HandleFunc("/enroll/", controllers.CreateEnroll).Methods("POST")
}
