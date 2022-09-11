package controllers

import (
	"Cafeteria/pkg/common/models/web"
	"Cafeteria/pkg/common/services/orchestrator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type UserController struct {
	UseCases *orchestrator.Orchestrator
}

func (cc *UserController) Init(router fiber.Router) {
	//router.Get("users/get_by_id", cc.GetByID)
	//router.Post("users/create", cc.Create)
	//router.Post("users/delete", cc.DeleteMark)
	router.Post("users/login", cc.Login)
}

func (cc *UserController) Login(ctx *fiber.Ctx) error {
	bodyData := web.User{}
	err := ctx.BodyParser(&bodyData)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	user := bodyData.User
	password := bodyData.Password

	dataStore, err := cc.UseCases.UserOrchestrator.GetByName(ctx.Context(), user, password)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  dataStore.User,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{"token": t})
}
