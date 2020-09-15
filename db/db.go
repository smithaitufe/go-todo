package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func Connect(conf Config) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	// cnstring := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Database, conf.Password)
	cnstring := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", conf.User, conf.Password, conf.Host, conf.Database)

	for {
		db, err = gorm.Open("postgres", cnstring)
		if err == nil {
			break
		}
		log.Printf("Could not open connection to db %v\n", err)
		time.Sleep(time.Second * 5)
	}
	return db
}
