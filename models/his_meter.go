package models

import "time"

type HisMeterData struct {
	Id       int    `gorm:"column:id"`
	MeterNo  string `gorm:"column:meterno"`
	CustomNo string `gorm:"column:custno"`
	AreaNo   int    `gorm:"column:areano"`
	CctNo    string `gorm:"column:cct_no"`
	SnrNo    int    `gorm:"column:snr_no"`
	UserID   int `gorm:"column:user_id"`
	JsrToTalAll float32 `gorm:"column:Jsr_TotalAll"`
	UpdateTime time.Time `gorm:"column:updateTime"`
}

func (HisMeterData) TableName() string {
	return "t_his_meter"
}
