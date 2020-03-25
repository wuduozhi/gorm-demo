package models

import "time"

type Company struct {
	Id          int       `gorm:"column:id"`
	CompanyCode string    `gorm:"column:companyCode"`
	Name        string    `gorm:"column:Name"`
	Province    string    `gorm:"column:Province"`
	City        string    `gorm:"column:city"`
	Status      int       `gorm:"column:Satus"`
	CreateTime  time.Time `gorm:"column:createTime"`
	ExportRule  string    `gorm:"column:exportRule"`
}

func (Company) TableName() string {
	return "t_company_info"
}
