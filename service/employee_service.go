package service

import (
	"package/dao"
	"package/dto"
	"package/model"
)

func GetAllEmployees(companyId int) ([]dto.EmployeeDTO, error) {
	
	employeeList, err := dao.GetAllEmployees(companyId)
	employeeDTOList := make([]dto.EmployeeDTO, len(employeeList))
	if err != nil{
		return employeeDTOList, err
	}else{
		for index, employee := range employeeList{
			employeeDTOList[index] = dto.EmployeeDTO{employee.Name, employee.Address, employee.CompanyId}
		}
		return employeeDTOList, nil
	}
	
}

func GetEmployeeDetails(employeeId, companyId int) (dto.EmployeeDTO, error){
	employeeDTO := dto.EmployeeDTO{}
	employee, err := dao.GetEmployeeDetails(employeeId, companyId)
	if err != nil{
		return employeeDTO, err
	}else{
		employeeDTO = dto.EmployeeDTO{Name: employee.Name, Address: employee.Address, CompanyId: employee.CompanyId}
		return employeeDTO, nil
	}
}

func DeleteEmployee(employeeId , companyId int) (error){
	return dao.DeleteEmployee(employeeId, companyId);
}

func SaveEmployee(employee model.Employee, companyId int) (dto.EmployeeDTO, error){
	employee.SetCompanyId(companyId)
	employeeDTO := dto.EmployeeDTO{}
	
	employee, err := dao.SaveEmployeeDetails(employee)
	if err != nil{
		return employeeDTO, err
	}else{
		employeeDTO = dto.EmployeeDTO{Name: employee.Name, Address: employee.Address, CompanyId: employee.CompanyId}
		return employeeDTO, nil
	}
}