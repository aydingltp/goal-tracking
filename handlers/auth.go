package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"goal-tracking/config"
	"goal-tracking/database"
	"goal-tracking/models/viewModels"
	"goal-tracking/service"
	"strings"
	"time"
)

func Login(c *fiber.Ctx) error {
	var model viewModels.LoginVm
	if err := c.BodyParser(&model); err != nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	model.EMail = strings.TrimSpace(strings.ToLower(model.EMail))
	model.Password = strings.TrimSpace(strings.ToLower(model.Password))

	db := database.DB()
	ss := service.NewUserService(db)

	user, err := ss.Login(&model)
	if user == nil {
		return c.JSON(fiber.Map{"is_success": false, "message": err, "data": nil})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["role"] = user.Role
	claims["email"] = user.Email
	claims["username"] = user.Name + " " + user.Surname
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "token": t})
}
