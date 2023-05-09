package usecase

import (
	"errors"
	"scheduler/entities"
	"scheduler/utils"
	"time"
)

func (uc *UseCase) CreateRecord(prevRecord string, scheduleId string, nowDoing string) (string, error) {

	// check, is scheduleId exist?
	if _, err := uc.ScheduleRepo.GetScheduleById(scheduleId); err != nil {
		return "", err
	}

	recordId, err := utils.RandomHex(16)
	if(err != nil) {
		return "", err
	}
	
	if prevRecord != "init-record" {
		record, err := uc.RecordRepo.GetRecordById(prevRecord)
		if err != nil {
			return "", err
		}
		record.Items[len(record.Items)-1].End = time.Now()
	}

	newRecord := entities.Record{
		Id: recordId,
		ScheduleId: scheduleId,
		Items: []entities.RecordItem{
			{
				Title: nowDoing,
				Start: time.Now(),
				End: time.Time{},
			},
		},
	}

	schedule, err := uc.ScheduleRepo.GetScheduleById(newRecord.ScheduleId)
	if(err != nil) {
		return "", err
	}
	if isExist := isSwitchToExistInTodos(schedule.Todos, nowDoing); !isExist {
		return "", errors.New("This todo: '" + nowDoing + "' is not exist in schedule")
	}
	
	if err := uc.RecordRepo.Insert(newRecord); err != nil {
		return "", err
	}

	return recordId, nil
}
