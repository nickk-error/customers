package db

import (
	"customer/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var log = logger.NewLogger()

var sqLitePool *gorm.DB

func ConnectSQLite(transID string) *gorm.DB {

	var err error
	sqLitePool, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Error(transID, "SQL Exception")
	}

	return sqLitePool
}
