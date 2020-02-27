package model

type Employee struct{
	// gorm.Model
	Id int
	Name string
	Address string
	CompanyId int

}

func (Employee) TableName() string {
	return "Employee"
}

func (employee *Employee) SetCompanyId(id int){
	employee.CompanyId = id
}