
package model

type Role struct{
	Id int
	RoleName string 
}

func (Role) TableName() string {
	return "Role"
}