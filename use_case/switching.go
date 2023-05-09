package usecase

import (
	"errors"
	"scheduler/entities"
	"time"
)

func (uc *UseCase) Switching(recordId string, switchTo string) error {

	record, err := uc.RecordRepo.GetRecordById(recordId)
	if(err != nil) {
		return err
	}

	schedule, err := uc.ScheduleRepo.GetScheduleById(record.ScheduleId)
	if(err != nil) {
		return err
	}
	if isExist := isSwitchToExistInTodos(schedule.Todos, switchTo); !isExist {
		return errors.New("This todo: '" + switchTo + "' is not exist in schedule")
	}

	record.Items[len(record.Items)-1].End = time.Now()
	record.Items = append(record.Items, entities.RecordItem{
		Title: switchTo,
		Start: time.Now(),
		End: time.Time{},
	})
	if err := uc.RecordRepo.Update(record); err != nil {
		return err
	}

	return nil
}
