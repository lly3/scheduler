package usecase

import (
	"errors"
	"scheduler/entities"
	"scheduler/utils"
	"time"
)

func (uc *UseCase) CreateRecord(scheduleId string, nowDoing string) (string, error) {

	// check, is scheduleId exist?
	if _, err := uc.ScheduleRepo.GetScheduleById(scheduleId); err != nil {
		return "", err
	}

	recordId, rxerr := utils.RandomHex(16)
	if(rxerr != nil) {
		return "", rxerr
	}

	// check, is latest record exist?
	record, glrerr := uc.RecordRepo.GetLatestRecord()
	if glrerr == nil {
		record.Items[len(record.Items)-1].End = time.Now()
		if err := uc.RecordRepo.Update(record); err != nil {
			return "", err
		}
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
