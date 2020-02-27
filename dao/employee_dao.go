package dao

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"package/model"
)

func SaveEmployeeDetails(employee model.Employee) (model.Employee, error) {
	log.Println("Save method called")
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	if err!=nil{
		log.Println("Connection Failed to Open",err)
		return employee, err
	}else{
		log.Println("Connection Established",employee.CompanyId, employee.Address)
		err := db.Create(&employee).Error
		if err != nil{
			log.Println("err",err)
			return employee, err
		}
		return employee, nil
	} 
}

func GetEmployeeDetails(employeeId, companyId int) (model.Employee, error){
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	employee := model.Employee{}
	if err!=nil{
		return employee, err
	} else{
		err = db.Where("id = ? AND company_id = ?",employeeId, companyId).Find(&employee).Error
		if err != nil {
			return employee, err
		}
		return employee, nil
	}
}

func GetAllEmployees(companyId int) ([]model.Employee, error){
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	employees := []model.Employee{}
	if err!=nil{
		return employees, err
	} else{
		db.Where("company_id = ?",companyId).Find(&employees)
		return employees, nil
	}
}

func DeleteEmployee(employeeId , companyId int) (error){
	db, err := gorm.Open("mysql", "root:password@/GSLAB?charset=utf8&parseTime=True")
	defer db.Close()
	employee := model.Employee{}
	if err!=nil{
		return err
	} else{
		err = db.Where("id = ? AND company_id= ?",employeeId, companyId).Delete(&employee).Error
		if err != nil {
			return err
		}
		return nil
	}
}