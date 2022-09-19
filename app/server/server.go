package server

import (
	"apiserver/server/controller"
	"apiserver/server/logger"
	"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	db   *sql.DB
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
	s.registerDB()
	s.registerMiddleware()
	s.registerController()

	err := s.app.Listen(fmt.Sprintf(":%d", s.port))
	return err
}

func (s *Server) registerDB() {
	db, err := sql.Open("mysql", "root:root@tcp(db-mysql:3306)/COMMON?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	s.db = db
}

func (s *Server) registerController() {
	controller.InitAccountController(s.app, s.db)
}

func (s *Server) registerMiddleware() {
	s.app.Use(RecoverAndStackTraceMiddleware)
	s.app.Use(RequestPrintMiddleware)
	s.app.Use(ResponsePrintMiddleware)
}
