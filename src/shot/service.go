// Dans le package shot/service.go

package shot

import (
	"mydj_server/src/entity"
	"time"

	"gorm.io/gorm"
	// "github.com/google/uuid"
)

// AddShotService ajoute un nouveau shot à la base de données
func AddShotService(newShot *entity.Shot, db *gorm.DB) (*entity.Shot, error) {
	// Générez un identifiant UUID pour le nouveau shot
	// newUUID := uuid.New()

	// Créez un nouvel enregistrement de shot avec les données fournies
	shot := &entity.Shot{
		ID:           newShot.ID,
		Name:         newShot.Name,
		Percentage:   newShot.Percentage,
		AlreadyTake: newShot.AlreadyTake,
		// AlreadyTaken: newShot.AlreadyTaken,
		Price:     newShot.Price,
		Points:    newShot.Points,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Ajoutez le nouveau shot à la base de données
	if err := db.Create(shot).Error; err != nil {
		return nil, err
	}

	return shot, nil
}
