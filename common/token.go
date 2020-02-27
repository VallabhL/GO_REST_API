package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"
	"strconv"
	"strings"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(user_email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte("V@11@bh"))
	if err != nil{
		return "",err
	}else{
		return result, nil
	}
}

func TokenValid(r *http.Request) (string, error) {
	tokenString := ExtractToken(r)
	var user_email string;
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("V@11@bh"), nil
	})
	if err != nil {
		return user_email, err
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_email = claims["user_id"].(string)
		//Pretty(claims)
	}
	return user_email, nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(r *http.Request) (uint32, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("V@11@bh"), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
