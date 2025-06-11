package api

import (
	"log"

	"github.com/Pratam-Kalligudda/user-service-go/config"
	api "github.com/Pratam-Kalligudda/user-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/user-service-go/internal/api/rest/handler"
	"github.com/Pratam-Kalligudda/user-service-go/internal/domain"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupServer(config config.Config) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("couldnt connect to database")
		return
	}
	log.Println("connection to database is succesfully")

	if err = db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalln("automirgration failed")
		return
	}

	log.Println("automigration succeded")

	setupRoutes(&api.RestHandler{
		App: app,
		DB:  db,
	})

	app.Listen(config.Secret)

}

func setupRoutes(rh *api.RestHandler) {
	handler.SetupUserHandler(rh)
}
