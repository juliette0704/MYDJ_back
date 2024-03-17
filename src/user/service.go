package user

import (
	"mydj_server/config"
	"mydj_server/src/entity"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserLoginService(email string, password string) (uuid.UUID, bool, error) {
	var err error
	var users []entity.User

	db, err := config.ReturnDB()

	db.Where("email = ? ", email).Find(&users)
	if len(users) == 0 {
		return uuid.UUID{}, false, err
	}
	if bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(password)) == nil {
		return users[0].UUID, true, nil
	}
	return uuid.UUID{}, false, err
}

func BeforeSaveService(u *entity.User, db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func SaveUserService(u *entity.User, db *gorm.DB) (*entity.User, error) {
	err := db.Omit("uid").Create(&u).Error
	if err != nil {
		return &entity.User{}, err
	}

	return u, nil
}
