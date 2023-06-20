package usecase

import (
	"fmt"
	"scheduler/entities"
	"scheduler/utils"
	"time"
)

func (uc *UseCase) CreateRecord(scheduleId string, nowDoing string) (string, error) {

	// check, is scheduleId exist?
	schedule, err := uc.ScheduleRepo.GetScheduleById(scheduleId)
	if err != nil {
		return "", ErrScheduleNotFound
	}

	// check, is latest record exist?
	if record, err := uc.RecordRepo.GetLatestRecord(); err == nil {
		record.Items[len(record.Items)-1].End = time.Now()
		if err := uc.RecordRepo.Update(record); err != nil {
			return "", err
		}
	}

	recordId, err := utils.RandomHex(16)
	if err != nil {
		return "", err
	}

	newRecord := entities.Record{
		Id:         recordId,
		ScheduleId: scheduleId,
		Items: []entities.RecordItem{
			{
				Title: nowDoing,
				Start: time.Now(),
				End:   time.Time{},
			},
		},
	}

	if isExist := isSwitchToExistInTodos(schedule.Todos, nowDoing); !isExist {
		return "", fmt.Errorf("This todo %s is not exist in schedule", nowDoing)
	}

	if err := uc.RecordRepo.Insert(newRecord); err != nil {
		return "", err
	}

	return recordId, nil
}
