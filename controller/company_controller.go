
package controller

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"package/service"
	"package/model"
	"strconv"
	"github.com/gorilla/mux"
	"package/responses"
)

func GetCompany(writer http.ResponseWriter, request *http.Request){
	fmt.Println("In controller")
	vars := mux.Vars(request)
	companyId, err := strconv.Atoi(vars["companyId"])
	if err !=nil{
		responses.JSON(writer, http.StatusBadRequest, responses.ResponseJSON{Message: "Unable to parse company id.", Timestamp: time.Now()})
	}else{
		companyDTO, err := service.GetCompanyDetails(companyId)
		if err!=nil{
			responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
		} else{
			responses.JSON(writer, http.StatusOK, responses.ResponseJSONWithData{Message: "Company details ", Timestamp: time.Now(), Data: companyDTO})
		}
	}
}

func GetAllCompanies(writer http.ResponseWriter, request *http.Request){
	companyDTOList, err := service.GetAllCompanies()
	if err!=nil{
		responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
	} else{
		responses.JSON(writer, http.StatusOK, responses.ResponseJSONWithData{Message: "Listed Companies ", Timestamp: time.Now(), Data: companyDTOList})
	}
}

func SaveCompany(writer http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	var company model.Company	
	err := decoder.Decode(&company)
	
    if err != nil {
        panic(err)
    }
	companyDTO, err := service.SaveCompany(company)
	if err!=nil{		
		responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
	} else{
		responses.JSON(writer, http.StatusOK, responses.ResponseJSONWithData{Message: "Company added to our records!", Timestamp: time.Now(), Data: companyDTO})
	}
}

func DeleteCompany(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	companyId, err := strconv.Atoi(vars["companyId"])
	if err !=nil{
		responses.JSON(writer, http.StatusBadRequest, responses.ResponseJSON{Message: "Unable to parse company id.", Timestamp: time.Now()})
	}else{
		err := service.DeleteCompany(companyId)
		if err!=nil{
			responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
		} else{
			responses.JSON(writer, http.StatusOK, responses.ResponseJSON{Message: "Company deleted.", Timestamp: time.Now()})
		}
	}
}