package shot

import (
	"net/http"

	"mydj_server/config"
	"mydj_server/src/entity"

	"github.com/gin-gonic/gin"
)

// AjouteShotController gère l'ajout d'un shot
func AddShotController(c *gin.Context) {
	db, _ := config.ReturnDB()

	// Créer une nouvelle instance de Shot
	var newShot entity.Shot

	// Lier les données de la requête à la structure Shot
	if err := c.ShouldBindJSON(&newShot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ajouter le shot en utilisant le service
	shot, err := AddShotService(&newShot, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add shot"})
		return
	}

	// Retourner une réponse réussie avec le shot ajouté
	c.JSON(http.StatusCreated, gin.H{"message": "Shot added successfully", "shot": shot})
}
