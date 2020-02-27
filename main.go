package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"database/sql"
	"package/controller"
	"package/common"
	// "github.com/casbin/casbin"
)

var db *sql.DB

func main(){

	//Gorilla mux to create a router, register end-points and handlers
	router := mux.NewRouter()
	// router.HandleFunc("/signup",controller.Signup).Methods("POST")
	router.HandleFunc("/login",controller.Login).Methods("POST")
	
	companyRouter := router.PathPrefix("/company").Subrouter()
	// companyRouter.Use(common.Test1, common.Test2("abc"), common.LoggingMiddleware)
	companyRouter.Use(common.LoggingMiddleware)

	companyRouter.HandleFunc("/",controller.GetAllCompanies).Methods("GET")
	companyRouter.HandleFunc("/",controller.SaveCompany).Methods("POST")
	companyRouter.HandleFunc("/{companyId:[0-9]+}/",controller.GetCompany).Methods("GET")
	companyRouter.HandleFunc("/{companyId:[0-9]+}/",controller.DeleteCompany).Methods("DELETE")

	employeeRouter := companyRouter.PathPrefix("/{companyId:[0-9]+}/employee").Subrouter()
	employeeRouter.HandleFunc("/",controller.GetAllEmployees).Methods("GET")
	employeeRouter.HandleFunc("/{employeeId:[0-9]+}/",controller.GetEmployee).Methods("GET")
	employeeRouter.HandleFunc("/",controller.SaveEmployee).Methods("POST")	
	employeeRouter.HandleFunc("/{employeeId:[0-9]+}/",controller.DeleteEmployee).Methods("DELETE")
	
	http.ListenAndServe(":8080",router)
}

func logFatal(err error) {
    if err != nil {
        log.Fatal(err)
    }
}