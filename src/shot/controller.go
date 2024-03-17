package shot

import (
	"net/http"

	"mydj_server/config"
	"mydj_server/src/entity"

	"github.com/gin-gonic/gin"
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

