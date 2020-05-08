package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wuduozhi/gorm-demo/models"
	"github.com/wuduozhi/gorm-demo/pool"
	"github.com/wuduozhi/gorm-demo/utils"
	"strings"
	"time"
)

var singleDb *gorm.DB
var mycatDb *gorm.DB

func initSingleDb() {
	dbUserName := "root"
	dbPassword := "123456"
	dbIP := "localhost"
	dbPort := "3306"
	dbName := "cbdata"

	path := strings.Join([]string{dbUserName, ":", dbPassword, "@(", dbIP, ":", dbPort, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
	var err error
	singleDb, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}

	initExtraDb(singleDb)
}

func initMyCatDb() {
	dbUserName := "root"
	dbPassword := "123456"
	dbIP := "120.79.214.246"
	dbPort := "8066"
	dbName := "TESTDB"

	path := strings.Join([]string{dbUserName, ":", dbPassword, "@(", dbIP, ":", dbPort, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
	var err error
	mycatDb, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}

	initExtraDb(mycatDb)
}

func initExtraDb(db *gorm.DB) {
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(2000)
	// 启用Logger，显示详细日志
	db.LogMode(false)
}

func init() {
	initSingleDb()
	initMyCatDb()
}

func CreateBenchCompany(count int, poolSize int, db *gorm.DB) ([]int, time.Duration) {
	prefix := "company-"
	companyIDs := make([]int, 0, count)

	doChan := make(chan struct{}, count)

	doFunc := func(req interface{}) interface{} {
		company := req.(models.Company)
		models.CreateCompany(company, db)
		doChan <- struct{}{}
		return nil
	}

	p := pool.NewFunc(poolSize, doFunc)
	defer p.Close()
	fmt.Println("start company...")
	startTime := time.Now()
	for i := 0; i < count; i++ {
		company := models.Company{
			Id:          i,
			CompanyCode: utils.GenValidateCode(prefix, 10),
			Name:        utils.GenValidateCode(prefix, 4),
			Province:    utils.GenValidateCode(prefix, 2),
			City:        utils.GenValidateCode(prefix, 2),
			CreateTime:  utils.GetRandomTime(),
			Status:      utils.RandomInt(0, 2),
		}
		go p.Process(company)
		//models.CreateCompany(company, db)
		companyIDs = append(companyIDs, company.Id)
	}

	for i := 0; i < count; i++ {
		<-doChan
	}

	useTime := time.Now().Sub(startTime)

	utils.WriteCreateLog(models.CreateLog{CompanyIDs: companyIDs})

	return companyIDs, useTime
}

func CreateBenchCct(companyIDs []int, count int, poolSize int, db *gorm.DB) (cctNoS []string, useTime time.Duration) {
	prefix := "cct-"

	createLog := utils.ReadCreateLog()
	if len(companyIDs) == 0 {
		companyIDs = createLog.CompanyIDs
	}

	doChan := make(chan struct{}, count)

	doFunc := func(req interface{}) interface{} {
		cct := req.(models.Cct)
		models.CreateCct(cct, db)
		doChan <- struct{}{}
		return nil
	}

	p := pool.NewFunc(poolSize, doFunc)
	defer p.Close()
	cctCompanyMap := make(map[string]int, count)

	fmt.Println("start cct...")

	startTime := time.Now()
	for i := 0; i < count; i++ {
		cct := models.Cct{
			CctNo:     utils.GenValidateCode(prefix, 10),
			CctName:   utils.GenValidateCode(prefix, 10),
			UserID:    utils.GetRandomIntItem(companyIDs),
			CctType:   "",
			CctStatus: utils.RandomInt(0, 2),
			AreaNo:    "",
		}
		go p.Process(cct)
		cctNoS = append(cctNoS, cct.CctNo)
		if _, ok := cctCompanyMap[cct.CctNo]; !ok {
			cctCompanyMap[cct.CctNo] = cct.UserID
		}
	}

	for i := 0; i < count; i++ {
		<-doChan
	}
	useTime = time.Now().Sub(startTime)

	createLog.CctNoS = cctNoS
	createLog.CctCompanyMap = cctCompanyMap
	utils.WriteCreateLog(createLog)

	return
}

func CreateBenchMeter(count, poolSize int, db *gorm.DB) (useTime time.Duration) {
	prefix := "meter-"

	createLog := utils.ReadCreateLog()
	cctNos := createLog.CctNoS
	cctCompanyMap := createLog.CctCompanyMap

	doChan := make(chan struct{}, count)

	doFunc := func(req interface{}) interface{} {
		meter := req.(models.Meter)
		models.CreateMeter(meter, db)
		doChan <- struct{}{}
		return nil
	}

	p := pool.NewFunc(poolSize, doFunc)
	defer p.Close()
	var meterNos []string
	meterCctMap := make(map[string]string, count)
	fmt.Println("start meter...")

	startTime := time.Now()
	for i := 0; i < count; i++ {
		cctNo := utils.GetRandomStringItem(cctNos)
		userID := cctCompanyMap[cctNo]
		meter := models.Meter{
			Id:       utils.RandomInt(0, 20*count),
			MeterNo:  utils.GenValidateCode(prefix, 10),
			CustomNo: utils.GenValidateCode(prefix, 3),
			AreaNo:   0,
			CctNo:    cctNo,
			SnrNo:    utils.RandomInt(0, 20),
			UserID:   userID,
		}
		go p.Process(meter)
		meterNos = append(meterNos, meter.MeterNo)
		if _, ok := meterCctMap[meter.MeterNo]; !ok {
			meterCctMap[meter.MeterNo] = meter.CctNo
		}
	}

	for i := 0; i < count; i++ {
		<-doChan
	}
	useTime = time.Now().Sub(startTime)

	createLog.MeterNos = meterNos
	createLog.MeterCctMap = meterCctMap
	utils.WriteCreateLog(createLog)

	return

}

func CreateBenchMeterData(count, poolSize int, db *gorm.DB) (useTime time.Duration) {
	createLog := utils.ReadCreateLog()
	meterNos := createLog.MeterNos
	meterCctMap := createLog.MeterCctMap

	doChan := make(chan struct{}, count)

	doFunc := func(req interface{}) interface{} {
		meterData := req.(models.MeterData)
		models.CreateMeterData(meterData, db)
		doChan <- struct{}{}
		return nil
	}

	p := pool.NewFunc(poolSize, doFunc)
	defer p.Close()
	fmt.Println("start meter data...")

	startTime := time.Now()
	for i := 0; i < count; i++ {
		meterNo := utils.GetRandomStringItem(meterNos)
		cctNo := meterCctMap[meterNo]
		meterData := models.MeterData{
			MeterNo:      meterNo,
			CctNo:        cctNo,
			LastTotalAll: utils.RandomFloat(),
			JsrToTalAll:  utils.RandomFloat(),
			CreateTime:   utils.GetRandomTime(),
		}
		go p.Process(meterData)
	}

	for i := 0; i < count; i++ {
		<-doChan
	}
	useTime = time.Now().Sub(startTime)

	return
}

func CreateBenchHisMeterData(count, poolSize int, db *gorm.DB) (useTime time.Duration) {
	createLog := utils.ReadCreateLog()
	meterNos := createLog.MeterNos
	meterCctMap := createLog.MeterCctMap
	cctCompanyMap := createLog.CctCompanyMap

	doChan := make(chan struct{}, count)

	doFunc := func(req interface{}) interface{} {
		//hisMeterData := req.(models.HisMeterData)
		//models.CreateHisMeterData(hisMeterData, db)
		doChan <- struct{}{}
		return nil
	}

	p := pool.NewFunc(poolSize, doFunc)
	defer p.Close()
	fmt.Println("start his meter data...")

	startTime := time.Now()
	for i := 0; i < count; i++ {
		meterNo := utils.GetRandomStringItem(meterNos)
		cctNo := meterCctMap[meterNo]
		userID := cctCompanyMap[cctNo]

		hisMeterData := models.HisMeterData{
			MeterNo:     meterNo,
			CustomNo:    utils.GenValidateCode("his-meter-data", 3),
			CctNo:       cctNo,
			SnrNo:       utils.RandomInt(0, 20),
			UserID:      userID,
			JsrToTalAll: utils.RandomFloat(),
			UpdateTime:  utils.GetRandomTime(),
		}
		go p.Process(hisMeterData)
	}

	for i := 0; i < count; i++ {
		<-doChan
	}
	useTime = time.Now().Sub(startTime)

	return
}

func SelectBenchHisMeterData(count, poolSize int, db *gorm.DB) (useTime time.Duration) {

	doChan := make(chan struct{}, count)

	doFunc := func(req interface{}) interface{} {
		id := req.(int)
		models.GetHisMeterData(id,db)
		doChan <- struct{}{}
		return nil
	}

	p := pool.NewFunc(poolSize, doFunc)
	defer p.Close()
	fmt.Println("start his meter data...")

	startTime := time.Now()
	for i := 0; i < count; i++ {
		id := utils.RandomInt(100,5000000)
		go p.Process(id)
	}

	for i := 0; i < count; i++ {
		<-doChan
	}
	useTime = time.Now().Sub(startTime)

	return
}

func PrintFormat(tableName, testDB string, poolSize, count int, useTime time.Duration) {
	fmt.Printf("Insert into %v %v records use %v seconds by %v goroutines.%v\n",
		tableName, count, useTime.Seconds(), poolSize, testDB)
}

func main() {
	poolSize := 8*2*2
	var useTime time.Duration

	companyCount := 100
	_, useTime = CreateBenchCompany(companyCount, poolSize, mycatDb)
	PrintFormat("comapny", "mycat", poolSize, companyCount, useTime)

	cctCount := 100
	_, useTime = CreateBenchCct(nil, cctCount, poolSize, mycatDb)
	PrintFormat("cct", "mycat", poolSize, cctCount, useTime)

	meterCount := 600
	useTime = CreateBenchMeter(meterCount, poolSize, mycatDb)
	PrintFormat("meter", "mycat", poolSize, cctCount, useTime)

	meterDataCount := 1000
	useTime = CreateBenchMeterData(meterDataCount, poolSize, mycatDb)
	PrintFormat("meter-data", "mycat", poolSize, meterDataCount, useTime)

	hisMeterDataCount := 2000
	useTime = CreateBenchHisMeterData(hisMeterDataCount, poolSize, mycatDb)
	PrintFormat("his-meter-data", "mycat", poolSize, hisMeterDataCount, useTime)

	//insertHisMeterDataCount := 1600
	//useTime = SelectBenchHisMeterData(insertHisMeterDataCount, poolSize, mycatDb)
	//PrintFormat("his-meter-data", "mycat", poolSize, insertHisMeterDataCount, useTime)
}
