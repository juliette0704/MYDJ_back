package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("./mydj_backend.db"), &gorm.Config{})
    if err != nil {
        log.Println(("err"))
        return nil, err
    }
    return db, nil
}

func ReturnDB() (*gorm.DB, error) {
    return ConnectDB()
}