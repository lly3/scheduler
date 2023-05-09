package usecase

import (
	"scheduler/entities"
)

func (uc *UseCase) GetSchedule(recordId string) (entities.RemainSchedule, error) {

	record, err := uc.RecordRepo.GetRecordById(recordId)
	if(err != nil) {
		return entities.RemainSchedule{}, err
	}

	schedule, err := uc.ScheduleRepo.GetScheduleById(record.ScheduleId)
	if(err != nil) {
		return entities.RemainSchedule{}, err
	}
	
	remainSchedule := calculateRemainingTime(schedule, record)

	return remainSchedule, nil
}
