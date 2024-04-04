package controller

import (
	"github.com/gofiber/fiber/v2"
	"simple_banking_application/internal/common/types"
	"simple_banking_application/internal/service"
)

type ITransactionController interface {
	Payment(ctx *fiber.Ctx) error
	RegisterRoutes(router fiber.Router)
}

type transactionController struct {
	transactionService service.ITransactionService
}

func NewTransactionController() ITransactionController {
	return &transactionController{transactionService: service.NewTransactionService()}
}

func (ctl *transactionController) RegisterRoutes(router fiber.Router) {
	transaction := router.Group("/payment")

	transaction.Post("", ctl.Payment)
}

func (ctl *transactionController) Payment(ctx *fiber.Ctx) error {
	body := new(types.PaymentRequest)

}
