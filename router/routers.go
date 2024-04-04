package router

import (
	"github.com/gofiber/fiber/v2"
	"simple_banking_application/internal/common/config"
	"simple_banking_application/internal/controller"
)

func SetUpRoutes() {

	app := fiber.New()
	var (
		transactionController = controller.NewTransactionController()
	)

	v1 := app.Group("/v1")

	// *****************
	// TRANSACTION ROUTES
	// *****************
	transactionController.RegisterRoutes(v1)

	app.Listen(config.Instance.Port)
}
