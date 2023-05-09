package usecase

import "scheduler/entities"

func (uc *UseCase) GetAllSchedule() ([]entities.Schedule, error) {
	
	schedule, err := uc.ScheduleRepo.GetAllSchedule()
	if err != nil {
		return []entities.Schedule{}, err
	}

	return schedule, nil
}
