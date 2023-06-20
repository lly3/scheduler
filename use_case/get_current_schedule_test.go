package usecase

import (
	"scheduler/entities"
	"testing"
	"time"
)

func TestGetCurrentSchedule(t *testing.T) {
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
	wantErr := false

	t.Run("should get current schedule", func(t *testing.T) {
		remainSchedule, err := uc.GetCurrentSchedule()
		if (err != nil) != wantErr {
			t.Errorf("GetCurrentSchedule() error = %v, wantErr = %v", err, wantErr)
		}

		assertEqual(t, remainSchedule.RemainItems[0].Title, "valid_todo_1")
		assertEqual(t, remainSchedule.RemainItems[0].Remain, time.Hour*4-time.Now().Sub(time.Date(2023, 6, 20, 10, 10, 10, 10, time.Local)))
	})
}
