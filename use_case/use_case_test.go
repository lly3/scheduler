package usecase

import (
	"errors"
	"scheduler/entities"
	"time"
)

var schedule = []entities.Schedule{
	{
		Id: "0",
		Todos: []entities.ScheduleItem{
			{
				Title: "Relax",
				Duration: time.Hour * 4,
			},
			{
				Title: "Gaming",
				Duration: time.Hour * 2,
			},
		},
	},
}

var records = []entities.Record{
	{
		Id: "0",
		ScheduleId: "0",
		Items: []entities.RecordItem{
			{
				Title: "Gaming",	
				Start: time.Now(),
				End: time.Now(),
			},
			{
				Title: "Relax",	
				Start: time.Now(),
				End: time.Time{},
			},
		},
	},
}

type mockScheduleRepo struct {}
type mockRecordRepo struct {}

func (s mockScheduleRepo) GetScheduleById(scheduleId string) (entities.Schedule, error) {
	for _, s := range schedule {
		if s.Id == scheduleId {
			return s, nil
		}
	}
	return entities.Schedule{}, errors.New("Can't find scheduleId: " + scheduleId)
}

func (s mockScheduleRepo) GetAllSchedule() ([]entities.Schedule, error) {
	return schedule, nil
}

func (r mockRecordRepo) GetRecordById(recordId string) (entities.Record, error) {
	for _, r := range records {
		if r.Id == recordId {
			return r, nil
		}
	}
	return entities.Record{}, errors.New("Can't find recordId: " + recordId)
}

func (r mockRecordRepo) Insert(record entities.Record) error {
	records = append(records, record)
	return nil
}

func (r mockRecordRepo) Update(record entities.Record) error {
	for i, r := range records {
		if r.Id == record.Id {
			records = append(records[:i], records[i+1:]...)
		}
	}	
	records = append(records, record)
	return nil
}

func newMockScheduleRepo() ScheduleRepository {
	return &mockScheduleRepo{}
}

func newMockRecordRepo() RecordRepository {
	return &mockRecordRepo{}
}
