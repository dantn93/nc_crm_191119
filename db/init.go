package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang191119/nc_crm/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
)

var db1 *gorm.DB

func Loc() *time.Location {
	loc, _ := time.LoadLocation("UTC")
	return loc
}

func GetDB() *gorm.DB {
	var once sync.Once

	once.Do(func() {
		db1 = new()
	})

	return db1
}

func new() *gorm.DB {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC().Truncate(1000 * time.Nanosecond)
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Config.Db.Host, config.Config.Db.User, config.Config.Db.Name, config.Config.Db.Password))
	if err != nil {
		log.Infof("Connect not success to postgres database at host:%s with user:%s and db:%s and %s",
			config.Config.Db.Host, config.Config.Db.User, config.Config.Db.Name, config.Config.Db.Password)
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().Ping()
	if config.Config.Db.Debug {
		db.LogMode(true)
	}
	return db
}
