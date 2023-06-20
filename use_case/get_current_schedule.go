package usecase

import (
	"scheduler/entities"
)

func (uc *UseCase) GetCurrentSchedule() (entities.RemainSchedule, error) {

	record, err := uc.RecordRepo.GetLatestRecord()
	if err != nil {
		return entities.RemainSchedule{}, err
	}

	schedule, err := uc.ScheduleRepo.GetScheduleById(record.ScheduleId)
	if err != nil {
		return entities.RemainSchedule{}, ErrScheduleNotFound
	}

	remainSchedule := calculateRemainingTime(schedule, record)

	return remainSchedule, nil
}
