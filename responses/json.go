package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseJSON struct{
	Message string
	Timestamp time.Time
}

type ResponseJSONWithData struct{
	Message string
	Timestamp time.Time
	Data interface{}
}

func JSON(writer http.ResponseWriter, statusCode int, data interface{}){
	writer.WriteHeader(statusCode)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil{
		fmt.Fprintf(writer, "%s", err.Error())
	}
}