package shot

import (
	"errors"
	"net/http"
	"strconv"
	"mydj_server/config"
	"mydj_server/src/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddShotController(c *gin.Context) {
	db, _ := config.ReturnDB()
	var newShot entity.Shot
	if err := c.ShouldBindJSON(&newShot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shot, err := AddShotService(&newShot, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add shot"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Shot added successfully", "shot": shot})
}

func GetAllShotsController(c *gin.Context) {
	db, _ := config.ReturnDB()

	shots, err := GetAllShotsService(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shots"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shots": shots})
}

func FindUserByUUID(uuid string) (*entity.User, error) {
	db, _ := config.ReturnDB()
	var user entity.User
	result := db.Where("uuid = ?", uuid).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func SaveShot(shot *entity.Shot) error {
	db, _ := config.ReturnDB()

	result := db.Save(shot)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SaveUser(user *entity.User) error {
	db, _ := config.ReturnDB()

	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindShotByName(name string) (*entity.Shot, error) {
	db, _ := config.ReturnDB()

	var shot entity.Shot
	result := db.Where("name = ?", name).First(&shot)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("shot not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &shot, nil
}
func AddShotToUserWithNameController(c *gin.Context) {
	uuid := c.Param("uuid")

	user, err := FindUserByUUID(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	shotName := c.Query("shotname")
	if shotName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name parameter is required"})
		return
	}

	numShotsStr := c.Query("nbshot")
	numShots := 1
	if numShotsStr != "" {
		numShots, err = strconv.Atoi(numShotsStr)
		if err != nil || numShots <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of shots"})
			return
		}
	}

	shot, err := FindShotByName(shotName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shot not found"})
		return
	}
	println("shot = ", shot.Name)
	shot.AlreadyTake = true

	for i := 0; i < numShots; i++ {
		user.Shots = append(user.Shots, *shot)
	}

	err = SaveUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	err = SaveShot(shot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shot(s) added successfully"})
}
