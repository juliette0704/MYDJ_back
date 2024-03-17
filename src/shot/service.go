// Dans le package shot/service.go

package shot

import (
	"errors"
	"mydj_server/src/entity"
	"time"

	"gorm.io/gorm"
)

func AddShotService(newShot *entity.Shot, db *gorm.DB) (*entity.Shot, error) {
	var existingShot entity.Shot
	if err := db.Where("name = ?", newShot.Name).First(&existingShot).Error; err == nil {
		return nil, errors.New("shot with the same name already exists")
	}
	shot := &entity.Shot{
		ID:          newShot.ID,
		Name:        newShot.Name,
		Percentage:  newShot.Percentage,
		AlreadyTake: newShot.AlreadyTake,
		Price:       newShot.Price,
		Points:      newShot.Points,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.Create(shot).Error; err != nil {
		return nil, err
	}

	return shot, nil
}

func GetAllShotsService(db *gorm.DB) ([]entity.Shot, error) {
	var shots []entity.Shot
	if err := db.Find(&shots).Error; err != nil {
		return nil, err
	}
	return shots, nil
}
