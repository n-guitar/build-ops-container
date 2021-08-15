package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/n-guitar/build-ops-container/pkg/buildapi"
	"github.com/n-guitar/build-ops-container/pkg/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func topPage(c *fiber.Ctx) error {
	return c.SendString("Hello Go Builder!!")
}

func initDatabase() *gorm.DB {
	var err error
	// databaseに接続。ファイルが無ければ作成
	database.DBConn, err = gorm.Open(sqlite.Open("./data/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	// Migrate the schema
	database.DBConn.AutoMigrate(&buildapi.BuildData{})
	fmt.Println("Database Migrated")
	return database.DBConn
}

func setupRoutes(app *fiber.App) {
	app.Get("/", topPage)

	// database crud
	app.Get("/api/v1/build", buildapi.GetBuildDataSet)
	app.Get("/api/v1/build/:id", buildapi.GetBuildData)
	app.Post("/api/v1/build", buildapi.NewBuildData)
	app.Delete("/api/v1/build/:id", buildapi.DeleteBuildData)

	// git cmd
	app.Post("/api/v1/build/:id/gitclone", buildapi.GitCloneCmdAPi)
}

func main() {
	fmt.Println("go builder")

	app := fiber.New()
	// Default config
	// app.Use(cors.New())
	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	db := initDatabase()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	// app.Get("/", topPage)
	setupRoutes(app)

	app.Listen(":3009")

	// Close
	defer sqlDB.Close()
}
