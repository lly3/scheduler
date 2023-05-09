package fiber

import (
	"scheduler/use_case"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	Uc *usecase.UseCase
	Server *fiber.App
}

type ServerConfig struct {
	Port string
}

func New(cfg *ServerConfig, uc *usecase.UseCase) *FiberServer {
	serv := fiber.New()

	f := &FiberServer{
		uc,
		serv,
	}

	f.CreateRecordRoutes()
	f.CreateScheduleRoutes()

	return f
}

func (f *FiberServer) Start(cfg *ServerConfig) {

	f.Server.Listen(":"+cfg.Port)
}
