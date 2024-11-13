package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"scan-table/config"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDBInstance() *sql.DB {
	once.Do(func() {
		config := config.GetConfig()
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Dbname)

		var err error
		db, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
	})

	return db
}
