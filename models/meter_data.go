package models

import "time"

type MeterData struct {
	MeterNo      string    `gorm:"column:meterno"`
	CctNo        string    `gorm:"column:cctno"`
	LastTotalAll float32   `gorm:"column:Last_TotalAll"`
	JsrToTalAll  float32   `gorm:"column:Jsr_TotalAll"`
	CreateTime   time.Time `gorm:"column:Createtime"`
}

func (MeterData) TableName() string {
	return "t_meterdata"
}
