package server

import (
	"backend/internal/database"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var validate = validator.New()

func (s *FiberServer) CreateItemHandler(c *fiber.Ctx) error {
	input := database.CreateOrUpdateItemInput{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := s.db.CreateItem(&input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": res})
}

func (s *FiberServer) GetItemsHandler(c *fiber.Ctx) error {
	res, err := s.db.GetAllItems()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": res})
}

func (s *FiberServer) GetItemHandler(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id", -1)

	if err != nil || ID < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := s.db.GetItem(uint(ID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": res})
}

func (s *FiberServer) UpdateItemHandler(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id", -1)

	if err != nil || ID < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	input := database.CreateOrUpdateItemInput{}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := s.db.UpdateItem(uint(ID), &input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": ID})
}

func (s *FiberServer) DeleteItemHandler(c *fiber.Ctx) error {
	ID, err := c.ParamsInt("id", -1)

	if err != nil || ID < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := s.db.DeleteItem(uint(ID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "data": ID})
}
