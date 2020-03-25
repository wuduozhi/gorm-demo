package models

type Cct struct {
	Id        int    `gorm:"column:cct_id"`
	CctNo     string `gorm:"column:cct_no"`
	CctName   string `gorm:"column:cct_name"`
	UserID    int `gorm:"column:User_id"`
	CctType   string `gorm:"column:cct_type"`
	CctStatus int    `gorm:"column:cct_status"`
	AreaNo    string `gorm:"column:areano"`
}

func (Cct) TableName() string {
	return "t_cctinfo"
}
