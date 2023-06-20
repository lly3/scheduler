package fiber

import "github.com/gofiber/fiber/v2"

type CreateRecordRequest struct {
	ScheduleId string `json:"schedule_id"`
	NowDoing   string `json:"now_doing"`
}

type SwitchRecordRequest struct {
	SwitchTo string `json:"switch_to"`
}

func (f *FiberServer) CreateRecordRoutes() {
	recordRoutes := f.Server.Group("/record")

	recordRoutes.Post("/", func(c *fiber.Ctx) error {
		request := CreateRecordRequest{}

		c.BodyParser(&request)

		recordId, err := f.Uc.CreateRecord(request.ScheduleId, request.NowDoing)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK).SendString(recordId)
		return nil
	})

	recordRoutes.Post("/switch", func(c *fiber.Ctx) error {
		request := SwitchRecordRequest{}

		c.BodyParser(&request)

		err := f.Uc.Switching(request.SwitchTo)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK)
		return nil
	})

	recordRoutes.Get("/remain", func(c *fiber.Ctx) error {

		remainSchedule, err := f.Uc.GetCurrentSchedule()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK).SendString(remainSchedule.String())
		return nil
	})

	recordRoutes.Get("/latest", func(c *fiber.Ctx) error {

		remainSchedule, err := f.Uc.GetCurrentSchedule()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK).SendString(remainSchedule.String())
		return nil
	})
}
