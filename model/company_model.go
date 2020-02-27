package model

type Company struct{

	Id int
	Name string
	Address string
}

func (Company) TableName() string {
	return "Company"
  }
