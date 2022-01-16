package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"goal-tracking/handlers"
	"goal-tracking/models"
	"strings"
)

func SetupRotes(app *fiber.App) {

	api := app.Group("/api")
	api.Post("/login", handlers.Login)

	api.Use(jwtware.New(jwtware.Config{
		SigningKey:     []byte("secret"),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
		//TokenLookup:    "header:" + fiber.HeaderAuthorization + ",query:token",
	}))

	user := api.Group("/user")
	user.Get("/", handlers.UserGetAll)
	user.Get("/:id", handlers.UserGetById)
	user.Post("/", handlers.UserCreate)
	user.Delete("/:id", handlers.UserDelete)

	classroom := api.Group("/classroom")
	classroom.Get("/", handlers.ClassroomGetAll)
	classroom.Get("/:id", handlers.ClassroomGetById)
	classroom.Post("/", handlers.ClassroomCreate)
	classroom.Delete("/:id", handlers.ClassroomDelete)

}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}

func jwtSuccess(c *fiber.Ctx) error {
	user := models.User{}
	tokenByte := c.Request().Header.Peek("Authorization")
	tokenStr := strings.ReplaceAll(string(tokenByte), "Bearer ", "")

	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return err
	}
	if !jwtToken.Valid {
		return err
	}

	claims := jwtToken.Claims.(jwt.MapClaims)

	userID := claims["userID"].(float64)
	user.ID = uint(userID)

	user.Email = claims["email"].(string)

	role := claims["role"].(float64)
	user.Role = models.Role(role)

	c.Locals("user", user)
	return c.Next()

}
