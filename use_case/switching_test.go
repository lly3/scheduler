package usecase

import (
	"scheduler/entities"
	"testing"
	"time"
)

func TestSwitching(t *testing.T) {
	t.Run("should switch when given correct todo", func(t *testing.T) {
		uc := &UseCase{
			newMockScheduleRepo(),
			newMockRecordRepo(),
		}
		uc.RecordRepo.Insert(entities.Record{
			Id:         "valid_recordId",
			ScheduleId: "valid_scheduleId",
			Items: []entities.RecordItem{
				{
					Title: "valid_todo_1",
					Start: time.Date(2023, 6, 20, 10, 10, 10, 10, time.Local),
					End:   time.Time{},
				},
			}},
		)
		uc.ScheduleRepo.Insert(entities.Schedule{
			Id: "valid_scheduleId",
			Todos: []entities.ScheduleItem{
				{
					Title:    "valid_todo_1",
					Duration: time.Hour * 4,
				},
				{
					Title:    "valid_todo_2",
					Duration: time.Hour * 4,
				},
			}},
		)
		switchTo, wantErr := "valid_todo_2", false

		err := uc.Switching(switchTo)
		if (err != nil) != wantErr {
			t.Errorf("Switching() error = %v, wantErr = %v", err, wantErr)
		}

		latestRecord, err := uc.RecordRepo.GetLatestRecord()
		if err != nil {
			t.Errorf("failed to GetLatestRecord(): %v", err)
		}

		assertEqual(t, latestRecord.Items[len(latestRecord.Items)-1].Title, "valid_todo_2")
		isPassEstimatedTime := timeEstimator(latestRecord.Items[len(latestRecord.Items)-1].Start.Sub(latestRecord.Items[len(latestRecord.Items)-2].End), time.Second*5)
		assertEqual(t, isPassEstimatedTime, true)
	})

	t.Run("should failed switch when given incorrect todo", func(t *testing.T) {
		uc := &UseCase{
			newMockScheduleRepo(),
			newMockRecordRepo(),
		}
		uc.RecordRepo.Insert(entities.Record{
			Id:         "valid_recordId",
			ScheduleId: "valid_scheduleId",
			Items: []entities.RecordItem{
				{
					Title: "valid_todo_1",
					Start: time.Date(2023, 6, 20, 10, 10, 10, 10, time.Local),
					End:   time.Time{},
				},
			}},
		)
		uc.ScheduleRepo.Insert(entities.Schedule{
			Id: "valid_scheduleId",
			Todos: []entities.ScheduleItem{
				{
					Title:    "valid_todo_1",
					Duration: time.Hour * 4,
				},
				{
					Title:    "valid_todo_2",
					Duration: time.Hour * 4,
				},
			}},
		)
		switchTo, wantErr := "invalid_todo", true

		err := uc.Switching(switchTo)
		if (err != nil) != wantErr {
			t.Errorf("Switching() error = %v, wantErr = %v", err, wantErr)
		}
	})
}
