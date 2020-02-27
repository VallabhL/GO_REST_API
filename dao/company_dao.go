package dao

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"package/model"
)

func GetAllCompanies() ([]model.Company, error){
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	companies := []model.Company{}
	if err!=nil{
		log.Println("Connection Failed to Open -")
		return companies, err
	} else{
		db.Find(&companies)
		return companies, nil
	}
}

func GetCompanyDetails(companyId int) (model.Company, error){
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	company := model.Company{}
	if err!=nil{
		return company, err
	} else{
		err = db.Where("id = ?",companyId).Find(&company).Error
		if err != nil {
			fmt.Println("Error in DAO: ",err)	
			return company, err
		}
		return company, nil
	}
}

func DeleteCompany(companyId int) error{
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	company := model.Company{}
	if err!=nil{
		return err
	} else{
		err = db.Where("id = ?",companyId).Delete(&company).Error
		if err != nil {
			return err
		}
		return nil
	}
}

func SaveCompanyDetails(company model.Company) (model.Company, error) {
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	if err!=nil{
		return company, err
	}else{
		err := db.Create(&company).Error
		if err != nil{
			return company, err
		}
		return company, nil
	} 
}