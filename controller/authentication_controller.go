package controller

import (
	"package/common"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"package/model"
	"net/http"
	"encoding/json"
	"package/responses"
	"time"
	// "fmt"
)

func Login(writer http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var userInput model.User	
	err := decoder.Decode(&userInput)
	
    if err != nil {
        panic(err)
    }

	user := model.User{}

	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()

	err = db.Where("email = ?", userInput.Email).Take(&user).Error
	if err != nil {
		responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
	}
	err = model.VerifyPassword(user.Password, userInput.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		responses.JSON(writer, http.StatusUnauthorized, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
	}
	token, err := common.CreateToken(user.Email)
	
	if err!=nil{
		responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: err.Error(), Timestamp: time.Now()})
	}else{
		responses.JSON(writer, http.StatusOK, responses.ResponseJSONWithData{Message: "User logged in successfully.", Timestamp: time.Now(), Data: token})	
    }
}


func Logout(writer http.ResponseWriter, request *http.Request) {
		responses.JSON(writer, http.StatusOK, responses.ResponseJSON{Message: "User logged out.", Timestamp: time.Now()})
}