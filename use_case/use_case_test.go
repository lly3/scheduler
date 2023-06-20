package usecase

import (
	"errors"
	"scheduler/entities"
	"testing"
)

type mockScheduleRepo struct {
	schedules []entities.Schedule
}

func (s *mockScheduleRepo) GetScheduleById(scheduleId string) (entities.Schedule, error) {
	for _, s := range s.schedules {
		if s.Id == scheduleId {
			return s, nil
		}
	}
	return entities.Schedule{}, errors.New("Can't find scheduleId: " + scheduleId)
}

func (s *mockScheduleRepo) GetAllSchedule() ([]entities.Schedule, error) {
	return s.schedules, nil
}

func (s *mockScheduleRepo) Insert(sc entities.Schedule) error {
	s.schedules = append(s.schedules, sc)
	return nil
}

type mockRecordRepo struct {
	records []entities.Record
}

func (r *mockRecordRepo) GetRecordById(recordId string) (entities.Record, error) {
	for _, r := range r.records {
		if r.Id == recordId {
			return r, nil
		}
	}
	return entities.Record{}, errors.New("Can't find recordId: " + recordId)
}

func (r *mockRecordRepo) GetLatestRecordId() (string, error) {
	return r.records[len(r.records)-1].Id, nil
}

func (r *mockRecordRepo) GetLatestRecord() (entities.Record, error) {
	if len(r.records) == 0 {
		return entities.Record{}, errors.New("There is no record yet")
	}
	return r.records[len(r.records)-1], nil
}

func (r *mockRecordRepo) Insert(record entities.Record) error {
	r.records = append(r.records, record)
	return nil
}

func (r *mockRecordRepo) Update(record entities.Record) error {
	for i, rec := range r.records {
		if rec.Id == record.Id {
			r.records = append(r.records[:i], r.records[i+1:]...)
		}
	}
	r.records = append(r.records, record)
	return nil
}

func newMockScheduleRepo() ScheduleRepository {
	return &mockScheduleRepo{}
}

func newMockRecordRepo() RecordRepository {
	return &mockRecordRepo{}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%v != %v", a, b)
	}
}
