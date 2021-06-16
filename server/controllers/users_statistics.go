package controllers

import (
	"context"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

func IndexPage(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetAllUsers(c *fiber.Ctx) error {
	ctx := c.Locals("ctx").(context.Context)

	db := ctx.Value("db").(*gorm.DB)
	users := database.NewUserRepo(db).GetAll()

	return c.JSON(users)
}

func GetTelegramId(c *fiber.Ctx, userId string) error {
	id, _ := strconv.ParseInt(userId, 10, 64)

	ctx := c.Locals("ctx").(context.Context)

	db := ctx.Value("db").(*gorm.DB)
	user := database.NewUserRepo(db).FindById(id)

	return c.JSON(user)
}

func GetUserHistory(c *fiber.Ctx, userId string) error {

	ctx := c.Locals("ctx").(context.Context)

	db := ctx.Value("db").(*gorm.DB)
	userHistory := database.NewUserHistoryRepo(db).FindAllRecordsByUsersID(userId)

	return c.JSON(userHistory)
}
