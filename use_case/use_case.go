package usecase

import (
	"errors"
	"scheduler/entities"
	"time"
)

var (
	ErrScheduleNotFound = errors.New("schedule not found")
	ErrRecordNotFound   = errors.New("record not found")
)

type UseCase struct {
	ScheduleRepo ScheduleRepository
	RecordRepo   RecordRepository
}

type RecordRepository interface {
	GetRecordById(recordId string) (entities.Record, error)
	GetLatestRecordId() (string, error)
	GetLatestRecord() (entities.Record, error)
	Insert(record entities.Record) error
	Update(record entities.Record) error
}

type ScheduleRepository interface {
	GetScheduleById(scheduleId string) (entities.Schedule, error)
	GetAllSchedule() ([]entities.Schedule, error)
	Insert(schedule entities.Schedule) error
}

func isSwitchToExistInTodos(todos []entities.ScheduleItem, switchTo string) bool {
	for _, v := range todos {
		if v.Title == switchTo {
			return true
		}
	}
	return false
}

func calculateRemainingTime(schedule entities.Schedule, record entities.Record) entities.RemainSchedule {
	remainSchedule := entities.RemainSchedule{}

	for _, todo := range schedule.Todos {
		var total time.Duration
		for i, item := range record.Items {
			if todo.Title == item.Title {
				var duration time.Duration
				if i == len(record.Items)-1 && record.Items[len(record.Items)-1].End.IsZero() {
					duration = time.Now().Sub(item.Start)
				} else {
					duration = item.End.Sub(item.Start)
				}
				total += duration
			}
		}

		remainSchedule.RemainItems = append(remainSchedule.RemainItems, entities.RemainItem{
			Title:  todo.Title,
			Remain: todo.Duration - total,
		})
	}

	return remainSchedule
}
