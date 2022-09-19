package controller

import (
	"apiserver/server/define"
	"apiserver/server/logger"
	"apiserver/server/service"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	service *service.AccountService
}

func InitAccountController(app *fiber.App, db *sql.DB) {
	_accountController = &AccountController{
		service: service.NewAccountService(db),
	}

	app.Post(define.LoginPath, _accountController.login)
	app.Post(define.JoinPah, _accountController.join)
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
			ResultCode: define.DBFail,
		})
	}

	if account == nil {
		logger.LogError("Not exist account")
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

func (c *AccountController) join(ctx *fiber.Ctx) error {
	req := define.JoinRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(define.BaseResponse{
			ResultCode: define.FailBodyParser,
		})
	}

	lastInsertId, err := c.service.CreateAccount(req.Nickname)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(define.BaseResponse{
			ResultCode: define.DBFail,
		})
	}

	return ctx.JSON(define.BaseResponse{
		Response: define.JoinReponse{
			Id: lastInsertId,
		},
	})
}

var _accountController *AccountController
