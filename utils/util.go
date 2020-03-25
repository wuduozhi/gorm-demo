package utils

import (
	"encoding/json"
	"fmt"
	"github.com/wuduozhi/gorm-demo/models"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

const LogPath = "create_log.json"

func init(){
	rand.Seed(time.Now().UnixNano())
}

func GetRandomIntItem(ls []int) int {
	l := len(ls)
	i := RandomInt(0, l)
	return ls[i]
}

func GetRandomStringItem(ls []string) string {
	l := len(ls)
	i := RandomInt(0, l)
	return ls[i]
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomFloat() float32 {
	return rand.Float32()
}

func GenValidateCode(prefix string, width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return prefix + sb.String()
}

func GetRandomTime() time.Time {
	random := RandomInt(0, 2*30)
	sd, _ := time.ParseDuration("-24h")
	return time.Now().Add(sd * time.Duration(random))
}

func WriteCreateLog(creatLog models.CreateLog) {
	data, err := json.Marshal(creatLog)
	if err != nil {
		log.Fatal(err)
	}
	Write(data)
}

func ReadCreateLog() models.CreateLog {
	data := Read()
	var createLog models.CreateLog
	err := json.Unmarshal(data, &createLog)
	if err != nil {
		log.Fatal(err)
	}
	return createLog
}

func Write(data []byte) {
	err := ioutil.WriteFile(LogPath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Read() []byte {
	data, err := ioutil.ReadFile(LogPath)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
