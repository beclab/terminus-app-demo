package apiserver

import (
	"github.com/beclab/terminus-app-demo/pkg/apiserver/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type server struct {
	address string
	app     *fiber.App
}

func New(address string) *server {
	return &server{address: address, app: fiber.New()}
}

func (s *server) Start() {
	s.app.Use(logger.New())

	handlers := handlers.New()

	api := s.app.Group("api")

	api.Get("/list", handlers.ListNotesGet)
	api.Post("/create", handlers.CreateNotesPost)
	api.Post("/delete", handlers.DeleteNotesPost)
	api.Post("/update", handlers.UpdateNotesPost)

	err := s.app.Listen(s.address)
	if err != nil {
		panic(err)
	}

}

func (s *server) Stop() {
	s.app.Shutdown()
}
