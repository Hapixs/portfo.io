package entities

import (
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestEntity struct {
	gorm.Model
	Name        string
	Age         int
	Version     int
	Description int
}

func SetupMySql(config *Config) {
	println("Connection to Mysql")
	dbi := config.Database
	dsn := dbi.User + ":" + dbi.Password + "@tcp(" + dbi.Host + ":" + strconv.Itoa(dbi.Port) + ")/" + dbi.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&TestEntity{})

}
