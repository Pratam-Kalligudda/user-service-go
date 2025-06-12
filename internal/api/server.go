package api

import (
	"log"

	"github.com/Pratam-Kalligudda/user-service-go/config"
	"github.com/Pratam-Kalligudda/user-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/user-service-go/internal/api/rest/handler"
	"github.com/Pratam-Kalligudda/user-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/user-service-go/internal/helper"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupServer(config config.Config) {
	app := fiber.New()
	app.Use(logger.New())
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("couldnt connect to database " + config.DSN + "\nerr : " + err.Error())
		return
	}
	log.Println("connection to database is succesfully")

	if err = db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalln("automirgration failed")
		return
	}

	log.Println("automigration succeded")

	auth := helper.NewAuthHelper(config.Secret)

	setupRoutes(&rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
	})

	app.Listen(config.Host)

}

func setupRoutes(rh *rest.RestHandler) {
	handler.SetupUserHandler(rh)
}
