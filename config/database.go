package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

func (config DbConfig) MysqlConnection() *gorm.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.Username, config.Password, config.Host, config.Port, config.DbName)

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
