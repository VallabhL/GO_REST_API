package service

import (
	"package/dao"
)

func GetRole(email string) string {
	role, err := dao.GetRole(email)
	if err!=nil{
		return ""
	}else{
		return role
	}
}