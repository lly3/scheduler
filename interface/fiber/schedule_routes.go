package fiber

import (
	usecase "scheduler/use_case"

	"github.com/gofiber/fiber/v2"
)

type CreateScheduleRequest struct{
	Todos []ScheduleItem `json:"todos"`
}

type ScheduleItem struct{
	Title string `json:"title"`
	Duration string `json:"duration"`
}

func (f *FiberServer) CreateScheduleRoutes() {
	scheduleRoutes := f.Server.Group("/schedule")

	scheduleRoutes.Get("/", func(c *fiber.Ctx) error {

		res, err := f.Uc.GetAllSchedule()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		
		result := ""

		for _, v := range res {
			result += v.String()
			result += "\n"
		}

		c.Status(fiber.StatusOK).SendString(result)
		return nil
	})

	scheduleRoutes.Get("/:id", func(c *fiber.Ctx) error {
		scheduleId := c.Params("id")
		res, err := f.Uc.GetScheduleById(scheduleId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		c.Status(fiber.StatusOK).SendString(res.String())
		return nil
	})

	scheduleRoutes.Post("/", func(c *fiber.Ctx) error {
		var request CreateScheduleRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		
		scheduleBody := toScheduleBody(request)
		scheduleId, err := f.Uc.CreateSchedule(scheduleBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK).SendString(scheduleId)
		return nil
	})
}

func toScheduleBody(req CreateScheduleRequest) usecase.ScheduleBody {
	var scheduleBodyItems []usecase.ScheduleBodyItem

	for _,v := range req.Todos {
		scheduleBodyItems = append(scheduleBodyItems, usecase.ScheduleBodyItem {
			Title: v.Title,
			Duration: v.Duration,
		})
	}

	return usecase.ScheduleBody {
		Todos: scheduleBodyItems,
	}
}
