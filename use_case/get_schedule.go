package usecase

import "scheduler/entities"

func (uc *UseCase) GetScheduleById(scheduleId string) (entities.Schedule, error) {
	schedule, err := uc.ScheduleRepo.GetScheduleById(scheduleId)
	if err != nil {
		return entities.Schedule{}, ErrScheduleNotFound
	}

	return schedule, nil
}
