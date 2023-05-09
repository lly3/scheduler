package usecase

import (
	"scheduler/entities"
	"scheduler/utils"
	"time"
)

type ScheduleBody struct {
	Todos []ScheduleBodyItem
}

type ScheduleBodyItem struct {
	Title string
	Duration string
}

func (uc *UseCase) CreateSchedule(scheduleBody ScheduleBody) (string, error) {

	scheduleId, err := utils.RandomHex(16)
	if err != nil {
		return "", err
	}

	schedule, err := toScheduleEntity(scheduleId, scheduleBody)
	if err != nil {
		return "", err
	}

	if err := uc.ScheduleRepo.Insert(schedule); err != nil {
		return "", err
	}

	return scheduleId, nil
}

func toScheduleEntity(scheduleId string, scheduleBody ScheduleBody) (entities.Schedule, error) {
	var scheduleItems []entities.ScheduleItem

	for _, v := range scheduleBody.Todos {
		duration, err := time.ParseDuration(v.Duration)
		if err != nil {
			return entities.Schedule{}, err
		}

		scheduleItems = append(scheduleItems, entities.ScheduleItem{
			Title: v.Title,
			Duration: duration,
		})
	}

	return entities.Schedule{
		Id: scheduleId,
		Todos: scheduleItems,
	}, nil
}
