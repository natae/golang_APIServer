package server

import (
	"apiserver/server/define"
	"apiserver/server/logger"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func RecoverAndStackTraceMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			stackTraceBuf := make([]byte, 1<<10)
			runtime.Stack(stackTraceBuf, true)

			logger.LogErrorF("PrintPanicStack: %v\n %s", r, stackTraceBuf)

			c.Status(fiber.StatusInternalServerError).JSON(
				define.BaseResponse{
					ResultCode: define.Panic,
				})
		}
	}()

	return c.Next()
}

func RequestPrintMiddleware(c *fiber.Ctx) error {
	logger.LogInfoF("Request: %s", string(c.Request().Body()))
	return c.Next()
}

func ResponsePrintMiddleware(c *fiber.Ctx) error {
	c.Next()
	logger.LogInfoF("Response: %s", string(c.Response().Body()))
	return nil
}
