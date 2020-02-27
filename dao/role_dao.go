package dao

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"package/model"
)

func GetRole(email string) (string, error){
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	role := model.Role{}
	if err!=nil{
		log.Println("Connection Failed to Open")
		return "", err
	} else{
		err := db.Table("Role").Select("role_name").Joins("JOIN UserRole on UserRole.role_id = Role.id").Joins("JOIN User on User.id = UserRole.user_id").Where("email = ?",email).Find(&role).Error
		if err != nil {
			return "",err
		}
		return role.RoleName, nil
	}
}