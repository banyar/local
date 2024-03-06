package services

import (
	"log"
	"os"
	"strconv"
	"strings"

	"git.frontiir.net/sa-dev/rtdatacore/pkg/driven/rtdb"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

}

// GetCustonFieldNames
func GetCustonFieldNames() []string {
	return strings.Split(os.Getenv("CUSTOM_FIELD_NAMES"), ",")
}

// GetCustonFieldIDs
func GetCustonFieldIDs() []uint {
	customFieldIDs := strings.Split(os.Getenv("CUSTOM_FIELD_IDS"), ",")
	var uintArr = []uint{}
	for _, i := range customFieldIDs {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		uintArr = append(uintArr, uint(j))
	}
	return uintArr
}

// ConnectDatabase
func ConnectDatabase() *gorm.DB {
	DB, err := rtdb.ConnectMySQL(GetDSNMySQL())
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err.Error())
	}
	return DB
}

// GetDSNMySQL
func GetDSNMySQL() *rtdb.DSNMySQL {
	DSN := rtdb.DSNMySQL{
		Host:     os.Getenv("MYSQL_DB_HOST"),
		Port:     os.Getenv("MYSQL_DB_PORT"),
		Username: os.Getenv("MYSQL_DB_USERNAME"),
		Password: os.Getenv("MYSQL_DB_PASSWORD"),
		Database: os.Getenv("MYSQL_DB_DATABASE"),
	}
	return &DSN
}

func GetApplyMiddleWare() bool {
	boolValue, err := strconv.ParseBool(os.Getenv("APPLY_MIDDLEWARE"))
	if err != nil {
		log.Fatal(err)
	}

	return boolValue

}
