package migration

import (
	"fmt"
	"log"
	"mydj_server/src/entity"

	"gorm.io/gorm"
)

var err error

func Migration(db *gorm.DB) {
	if err = db.AutoMigrate(&entity.User{}); err != nil {
		log.Println("Error", fmt.Errorf("").Error())
	}
}
