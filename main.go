package main

import (
	// "os/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "gorm.io/gorm"

	"mydj_server/config"
	// "mydj_server/src/database"
	"mydj_server/src/migration"
	"mydj_server/src/route"
)

// func Migrate(db *gorm.DB) error {
// 	// Exécuter la migration pour créer la table "users"
// 	if err := db.AutoMigrate(&user.User{}); err != nil {
// 		return err
// 	}

// 	return nil
// }

func main() {
	// Initialise la configuration
	config.InitConfig("./config/config.yaml")

	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("Failed to get database connection")
		}
		err = sqlDB.Close()
		if err != nil {
			panic("Failed to close database connection")
		}
	}()

	// Migration de la base de données

	// database_migration.Migration(db) // Vous pouvez exécuter la migration ici si nécessaire
	// Migrate(db)
	// Configure Gin
	migration.Migration(db)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:     true,
		AllowPrivateNetwork: true,
		AllowMethods:        []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:        []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Token", "Access-Control-Allow-Origin"},
		AllowCredentials:    true,
		MaxAge:              12 * time.Hour,
	}))

	// Configurer les routes
	route.RoutingUser(router, db)
    route.RoutingShot(router, db)
	// route.RoutingGroup(router, db)
	// Ajoutez d'autres routes ici selon vos besoins

	// Exécute le serveur Gin
	err = router.Run(": 8080")
	if err != nil {
		panic("Failed to start server")
	}
}
