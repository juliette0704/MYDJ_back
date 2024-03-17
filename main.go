package main

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mydj_server/config"
	"mydj_server/src/migration"
	"mydj_server/src/route"
)

func main() {
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

	route.RoutingUser(router, db)
    route.RoutingShot(router, db)

	err = router.Run(": 8080")
	if err != nil {
		panic("Failed to start server")
	}
}
