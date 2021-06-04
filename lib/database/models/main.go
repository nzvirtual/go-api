package models

import (
	"fmt"
	"time"

	log "github.com/dhawton/log4g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var MaxAttempts = 10
var DelayBetweenAttempts = time.Minute * 1
var attempt = 1

func Connect(user string, pass string, hostname string, port string, database string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", user, pass, hostname, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Category("db").Error("Error connecting to database: " + err.Error())
		if attempt < MaxAttempts {
			log.Category("db").Info(fmt.Sprintf("Attempt %d/%d Failed. Waiting %s before trying again...", attempt, MaxAttempts, DelayBetweenAttempts.String()))
			time.Sleep(DelayBetweenAttempts)
			attempt += 1
			Connect(user, pass, hostname, port, database)
			return
		}
		panic("Max attempts occured. Aborting startup.")
	}

	db.AutoMigrate(&Airport{}, &Rank{}, &User{})

	DB = db
}
