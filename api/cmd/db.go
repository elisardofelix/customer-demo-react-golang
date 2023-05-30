package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLConnect() *gorm.DB {
	dsn := "root:ttpass@tcp(host.docker.internal:3306)/cusdb"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
