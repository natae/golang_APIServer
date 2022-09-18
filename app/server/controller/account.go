package controller

import (
	"apiserver/server/define"
	"apiserver/server/logger"
	"apiserver/server/service"

	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	service *service.AccountService
}

func InitAccountController(app *fiber.App) {
	_accountController = &AccountController{service: &service.AccountService{}}

	app.Post(define.LoginPath, _accountController.login)
}

func (c *AccountController) login(ctx *fiber.Ctx) error {
	req := define.LoginRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(define.BaseResponse{
			ResultCode: define.FailBodyParser,
		})
	}

	account, err := c.service.GetAccount(req.Nickname)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(define.BaseResponse{
			ResultCode: define.NotExistUser,
		})
	}

	return ctx.JSON(define.BaseResponse{
		Response: define.LoginResponse{
			Id: account.Id,
		},
	})
}

var _accountController *AccountController
