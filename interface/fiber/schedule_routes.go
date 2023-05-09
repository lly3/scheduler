package fiber

import "github.com/gofiber/fiber/v2"

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
}
