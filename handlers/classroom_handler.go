package handlers

import (
	"github.com/gofiber/fiber/v2"
	"goal-tracking/database"
	"goal-tracking/models"
	"goal-tracking/service"
	"strconv"
)

func ClassroomGetAll(c *fiber.Ctx) error {
	db := database.DB()
	ss := service.NewClassroomService(db)
	data, err := ss.GetAll()
	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"is_success": true, "message": "success", "data": data})
}

func ClassroomCreate(c *fiber.Ctx) error {
	var requestBody models.Classroom
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	db := database.DB()
	ss := service.NewClassroomService(db)
	err = ss.Create(&requestBody)
	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"is_success": true, "message": "successfully created!", "data": nil})
}

func ClassroomGetById(c *fiber.Ctx) error {
	idS := c.Params("id")
	id, _ := strconv.Atoi(idS)

	db := database.DB()
	ss := service.NewClassroomService(db)

	data, err := ss.GetById(int64(id))

	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	return c.JSON(data)
}

func ClassroomDelete(c *fiber.Ctx) error {
	idS := c.Params("id")
	id, _ := strconv.Atoi(idS)

	db := database.DB()
	ss := service.NewClassroomService(db)
	err := ss.Delete(int64(id))

	if err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	return c.JSON(fiber.Map{"is_success": true, "message": "user deleted", "data": nil})
}
