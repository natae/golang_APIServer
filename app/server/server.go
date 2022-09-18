package server

import (
	"apiserver/server/controller"
	"apiserver/server/logger"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	port int
}

func NewServer(port int, logLevel logger.LogLevel) *Server {
	logger.InitLogger(logLevel)

	return &Server{
		app:  fiber.New(),
		port: port,
	}
}

func (s *Server) Start() error {
	s.registerMiddleware()
	s.registerController()

	err := s.app.Listen(fmt.Sprintf(":%d", s.port))
	return err
}

func (s *Server) registerController() {
	controller.InitAccountController(s.app)
}

func (s *Server) registerMiddleware() {
	//s.app.Use(recover.New())
	s.app.Use(RecoverAndStackTraceMiddleware)
	s.app.Use(RequestPrintMiddleware)
	s.app.Use(ResponsePrintMiddleware)
}
