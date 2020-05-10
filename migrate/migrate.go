package migrate

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wuduozhi/gorm-demo/models"
)

func DoMigrateHisMeter(limit int, old, new *gorm.DB) {
	var (
		id           = 1000000
		offset       = 0
		sum          = 0
		checkMeterNo = make(map[string]struct{})
	)
	for {
		oldHisMeterDatas := models.GetHisMeterDataWithLimit(id, 0, limit, old)
		for i, m := range oldHisMeterDatas {
			//fmt.Println(m.Id)
			if i%7 == 0 {
				checkMeterNo[m.MeterNo] = struct{}{}
			}
			m.Id = 0
			models.CreateHisMeterData(m, new)
		}

		fmt.Printf("Migrate.Get %v records from old db.\n", len(oldHisMeterDatas))
		if len(oldHisMeterDatas) > 1 {
			id = oldHisMeterDatas[len(oldHisMeterDatas)-1].Id
		}
		offset += limit - 1
		sum += len(oldHisMeterDatas)

		if len(oldHisMeterDatas) < limit {
			fmt.Printf("Migrate done.All %v records.\n", sum)
			break
		}
	}

	MigrateCheckHisMeter(checkMeterNo, old, new)
}

func MigrateCheckHisMeter(checkMeterNo map[string]struct{}, old, new *gorm.DB) {
	fmt.Println("Begin check...")
	sum := 0
	for meterNo := range checkMeterNo {
		oldMeters := models.GetHisMeterByMeterNo(meterNo, old)
		newMeters := models.GetHisMeterByMeterNo(meterNo, new)
		if !checkHisMeter(oldMeters, newMeters) {
			sum++
			fmt.Printf("Migrate errror.MeterNo:%v\n", meterNo)
		}
	}
	fmt.Printf("MigrateCheckHisMeter done.Error Record:%v", sum)
}

func checkHisMeter(olds, news []models.HisMeterData) bool {
	if len(olds) != len(news) {
		return false
	}

	for i := 0; i < len(olds); i++ {
		old := olds[i]
		new := news[i]
		if old.MeterNo != new.MeterNo || old.CctNo != new.CctNo || old.UserID != new.UserID {
			return false
		}
	}
	return true
}
