package handlers

import (
	"github.com/gofiber/fiber/v2"
	"goal-tracking/database"
	"goal-tracking/models"
	"goal-tracking/service"
	"strconv"
)

func UserGetAll(c *fiber.Ctx) error {
	db := database.DB()
	ss := service.NewUserService(db)
	data, err := ss.GetAll()
	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"is_success": true, "message": "success", "data": data})
}

func UserCreate(c *fiber.Ctx) error {
	var requestBody models.User
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	db := database.DB()
	ss := service.NewUserService(db)
	err = ss.Create(&requestBody)
	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"is_success": true, "message": "successfully created!", "data": nil})
}

func UserGetById(c *fiber.Ctx) error {
	idS := c.Params("id")
	id, _ := strconv.Atoi(idS)

	db := database.DB()
	ss := service.NewUserService(db)

	data, err := ss.GetById(int64(id))

	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	return c.JSON(data)
}

func UserDelete(c *fiber.Ctx) error {
	idS := c.Params("id")
	id, _ := strconv.Atoi(idS)

	db := database.DB()
	ss := service.NewUserService(db)
	err := ss.Delete(int64(id))

	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"is_success": true, "message": "user deleted", "data": nil})
}
