
package controller

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"package/service"
	"package/model"
	"strconv"
	"github.com/gorilla/mux"
	"package/responses"
)

func SaveEmployee(writer http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	var employee model.Employee	
	decoder.Decode(&employee)
	
	vars := mux.Vars(request)
	companyId, err := strconv.Atoi(vars["companyId"])
	if err !=nil{
		responses.JSON(writer, http.StatusBadRequest, responses.ResponseJSON{Message: "Unable to parse company id.", Timestamp: time.Now()})
	}else{
		fmt.Println("Called once")
		_, err := service.SaveEmployee(employee, companyId)
		if err!=nil{		
			responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
		} else{
			responses.JSON(writer, http.StatusOK, responses.ResponseJSON{Message: "Employee added to our records!", Timestamp: time.Now()})
		}
	}
}

func GetEmployee(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	employeeId, err1 := strconv.Atoi(vars["employeeId"])
	companyId, err2 := strconv.Atoi(vars["companyId"])

	if (err1 !=nil || err2 !=nil){
		responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: "Unable to parse given input parameters.", Timestamp: time.Now()})
	}else{
		employee, err := service.GetEmployeeDetails(employeeId, companyId)
		if err!=nil{
			responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
		} else{
			responses.JSON(writer, http.StatusOK, responses.ResponseJSONWithData{Message: "Employee details ", Timestamp: time.Now(), Data: employee})
		}
	}
}

func GetAllEmployees(writer http.ResponseWriter, request *http.Request){
	
	vars := mux.Vars(request)
	companyId, err := strconv.Atoi(vars["companyId"])

	if (err !=nil){
		responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: "Unable to parse given input parameters.", Timestamp: time.Now()})
	} else{
		employeeList, err := service.GetAllEmployees(companyId);
		if err != nil {
			responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
		}else{
			responses.JSON(writer, http.StatusOK, responses.ResponseJSONWithData{Message: "Listed Employees ", Timestamp: time.Now(), Data: employeeList})
		}
	}
}



func DeleteEmployee(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	employeeId, err1 := strconv.Atoi(vars["employeeId"])
	companyId, err2 := strconv.Atoi(vars["companyId"])

	if (err1 !=nil || err2 !=nil){
		responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: "Unable to parse given input parameters.", Timestamp: time.Now()})
	}else{
		err := service.DeleteEmployee(employeeId, companyId)
		if err!=nil{
			responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
		} else{
			responses.JSON(writer, http.StatusOK, responses.ResponseJSON{Message: "Employee deleted.", Timestamp: time.Now()})
		}
	}
}
