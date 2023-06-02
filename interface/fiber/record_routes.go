package fiber

import "github.com/gofiber/fiber/v2"

type CreateRecordRequest struct{
	PrevRecordId string `json:"prev_record_id"`
	ScheduleId string `json:"schedule_id"`
	NowDoing string `json:"now_doing"`
}

type SwitchRecordRequest struct {
	RecordId string `json:"record_id"`
	SwitchTo string `json:"switch_to"`
}

func (f *FiberServer) CreateRecordRoutes() {
	recordRoutes := f.Server.Group("/record")

	recordRoutes.Post("/switch", func(c *fiber.Ctx) error {
		request := SwitchRecordRequest{}

		c.BodyParser(&request)

		err := f.Uc.Switching(request.RecordId, request.SwitchTo)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK)
		return nil
	})


	recordRoutes.Post("/", func(c *fiber.Ctx) error {
		request := CreateRecordRequest{}

		c.BodyParser(&request)

		res, err := f.Uc.CreateRecord(request.PrevRecordId, request.ScheduleId, request.NowDoing)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK).SendString(res)
		return nil
	})

	recordRoutes.Get("/remain/:recordId", func(c *fiber.Ctx) error {
		recordId := c.Params("recordId")

		res, err := f.Uc.GetSchedule(recordId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK).SendString(res.String())
		return nil
	})

	recordRoutes.Get("/latest", func(c *fiber.Ctx) error {

		res, err := f.Uc.GetLatestRecordId()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Status(fiber.StatusOK).SendString(res)
		return nil
	})
}
