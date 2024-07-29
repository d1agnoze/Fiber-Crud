package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	s.App.Get("/health", s.healthHandler)

	itemGroup := s.App.Group("/item")
	itemGroup.Get("/list", s.GetItemsHandler)
	itemGroup.Get("/:id<int;min(0)>", s.GetItemHandler)
	itemGroup.Patch("/update/:id<int;min(0)>", s.UpdateItemHandler)
	itemGroup.Delete("/del/:id<int;min(0)>", s.DeleteItemHandler)
	itemGroup.Post("/create", s.CreateItemHandler)
}

func (*FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{"message": "Hello World"}
	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
