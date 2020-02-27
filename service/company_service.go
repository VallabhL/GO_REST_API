package service

import (
	"package/dao"
	"package/dto"
	"package/model"
)

func GetAllCompanies() ([]dto.CompanyDTO, error) {
	companies, err := dao.GetAllCompanies()
	companyDTOList := make([]dto.CompanyDTO, len(companies))
	for index, company := range companies{
		companyDTOList[index] = dto.CompanyDTO{Name: company.Name, Address: company.Address}
	}
	return companyDTOList, err
}

func GetCompanyDetails(companyId int) (dto.CompanyDTO, error){
	company, err := dao.GetCompanyDetails(companyId)
	companyDTO := dto.CompanyDTO{}
	if err != nil{
		return companyDTO, err
	}else{
		companyDTO = dto.CompanyDTO{Name: company.Name, Address: company.Address}
	}
	return companyDTO, err
}

func DeleteCompany(companyId int) error{
	return dao.DeleteCompany(companyId);
}

func SaveCompany(company model.Company) (dto.CompanyDTO, error){
	company, err := dao.SaveCompanyDetails(company)
	companyDTO := dto.CompanyDTO{}
	if err != nil{
		return companyDTO, err
	}else{
		companyDTO = dto.CompanyDTO{Name: company.Name, Address: company.Address}
		return companyDTO, nil
	}
}
