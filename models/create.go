package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CreateLog struct {
	CompanyIDs    []int
	CctNoS        []string
	CctCompanyMap map[string]int
	MeterNos      []string
	MeterCctMap   map[string]string
}

func CreateCompany(model Company, db *gorm.DB) {
	db.Create(&model)
}

func CreateCct(model Cct, db *gorm.DB) {
	db.Create(&model)
}

func CreateMeter(model Meter, db *gorm.DB) {
	db.Create(&model)
}

func CreateMeterData(model MeterData, db *gorm.DB) {
	db.Create(&model)
}

func CreateHisMeterData(model HisMeterData, db *gorm.DB) {
	db.Create(&model)
}

func GetHisMeterData(id int, db *gorm.DB) HisMeterData {
	var hisMeterData HisMeterData
	db.Where("id = ?", id).First(&hisMeterData)
	return hisMeterData
}

func createCompany(db *gorm.DB) {
	c := Company{
		Id:          11,
		CompanyCode: "gorm",
		Name:        "gorm-123",
		Province:    "111",
		City:        "111",
		Status:      0,
		CreateTime:  time.Now(),
	}
	db.Create(&c)
}

func createCct(db *gorm.DB) {
	c := Cct{
		CctNo:   "gorm-cct",
		CctName: "gorm-cct",
		UserID:  11,
	}
	db.Create(&c)
}

func createMeter(db *gorm.DB) {
	m := Meter{
		Id:      123456800,
		MeterNo: "gorm-meter",
		CctNo:   "gorm-cct",
		SnrNo:   0,
		UserID:  11,
	}
	db.Create(&m)
}

func createMeterData(db *gorm.DB) {
	meterData := MeterData{
		MeterNo:      "gorm-meter",
		CctNo:        "gorm-cct",
		LastTotalAll: 0,
		JsrToTalAll:  0,
		CreateTime:   time.Time{},
	}
	db.Create(&meterData)
}

func createHisMeterData(db *gorm.DB) {
	hisMeterData := HisMeterData{
		MeterNo:     "gorm-meter",
		CctNo:       "gorm-cct",
		CustomNo:    "",
		AreaNo:      0,
		SnrNo:       0,
		UserID:      0,
		JsrToTalAll: 0,
		UpdateTime:  time.Now(),
	}
	db.Create(&hisMeterData)
}
