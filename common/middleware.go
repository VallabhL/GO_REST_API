package common

import (
	"net/http"
	"package/responses"
	"fmt"
	"package/service"
	"time"
	"os"
	"strings"
	"github.com/casbin/casbin"
)

/*func LoggingMiddlewareGenerator(e *casbin.Enforcer) ( func (next http.Handler) http.Handler){
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		
		user_email, err := TokenValid(request)
		
		if err != nil {
			responses.JSON(writer, http.StatusUnauthorized, responses.ResponseJSON{Message: "Authentication failed.", Timestamp: time.Now()})
		}else{
			fmt.Println("user_email: ",user_email)
			role := service.GetRole(user_email)
			//casbin rule enforcing
			res, err := e.EnforceSafe(role, request.URL.Path, request.Method)
			if err != nil {
				responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: "Authorization failed.", Timestamp: time.Now()})
			}			
		}
	
		//next.ServeHTTP(writer, request)
    })
}*/

var authEnforcer *casbin.Enforcer

func init(){
	var err error
	authEnforcer, err = casbin.NewEnforcerSafe("authorization/auth_model.conf", "authorization/policy.csv")  
	if err!=nil{
		pwd, _ := os.Getwd()
		fmt.Println("Present working directory: ",pwd)
		panic("Failed to load authorization policies.");
	}  
}

// func Test1() http.Handler{
// 	return http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request){
// 		fmt.Println("Test 1 middleware")
// 		next.ServeHTTP(writer, request)
// 	})
// }

// func Test2(data string) http.Handler{
// 	return http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request){
// 		fmt.Println("Test 2 middleware",data)
// 		next.ServeHTTP(writer, request)
// 	})
// }

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		
		user_email, err := TokenValid(request)
		
		if err != nil {
			responses.JSON(writer, http.StatusUnauthorized, responses.ResponseJSON{Message: "Authentication failed.", Timestamp: time.Now()})
			return 
		}else{
			role := service.GetRole(user_email)
			//casbin rule enforcing
			res, err := authEnforcer.EnforceSafe(strings.TrimSpace(role), request.URL.Path, request.Method)
			if err != nil {
				responses.JSON(writer, http.StatusInternalServerError, responses.ResponseJSON{Message: "Authorization failed.", Timestamp: time.Now()})
			}
			if res {
				next.ServeHTTP(writer, request)

            } else {
				responses.JSON(writer, http.StatusForbidden, responses.ResponseJSON{Message: "Authorization failed.", Timestamp: time.Now()})
			}
			return	
		}
    })
}
