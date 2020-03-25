package models

type Meter struct {
	Id       int    `gorm:"column:id"`
	MeterNo  string `gorm:"column:meterno"`
	CustomNo string `gorm:"column:custno"`
	AreaNo   int    `gorm:"column:areano"`
	CctNo    string `gorm:"column:cct_no"`
	SnrNo    int    `gorm:"column:snr_no"`
	UserID   int `gorm:"column:user_id"`
}

func (Meter) TableName() string {
	return "t_meterinfo"
}
