package server

import (
	"backend/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "backend",
			AppName:      "backend",
		}),

		db: database.New(),
	}

	server.App.Use(cors.New())
	server.db.Migrate()

	return server
}
