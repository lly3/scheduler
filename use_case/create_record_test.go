package usecase

import (
	"scheduler/entities"
	"testing"
	"time"
)

func TestCreateRecord(t *testing.T) {
	t.Run("should create record when there is no record yet", func(t *testing.T) {
		uc := &UseCase{
			ScheduleRepo: newMockScheduleRepo(),
			RecordRepo:   newMockRecordRepo(),
		}
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
		scheduleId, nowDoing := "valid_scheduleId", "valid_todo_2"
		wantErr := false

		recordId, err := uc.CreateRecord(scheduleId, nowDoing)
		if (err != nil) != wantErr {
			t.Errorf("failed to CreateRecord(): %v, wantErr %v", err, wantErr)
		}

		record, err := uc.RecordRepo.GetRecordById(recordId)
		if err != nil {
			t.Errorf("failed to GetRecordById(): %v", err)
		}

		assertEqual(t, record.ScheduleId, scheduleId)
		assertEqual(t, record.Items[len(record.Items)-1].Title, nowDoing)
		assertEqual(t, record.Items[len(record.Items)-1].End.IsZero(), true)
	})

	t.Run("should create record when records exist", func(t *testing.T) {
		uc := &UseCase{
			newMockScheduleRepo(),
			newMockRecordRepo(),
		}
		uc.RecordRepo.Insert(entities.Record{
			Id:         "valid_recordId",
			ScheduleId: "valid_scheduleId",
			Items: []entities.RecordItem{
				{
					Title: "Gaming",
					Start: time.Now(),
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
		scheduleId, nowDoing := "valid_scheduleId", "valid_todo_2"
		wantErr := false

		latestRecord, err := uc.RecordRepo.GetLatestRecord()
		if err != nil {
			t.Errorf("failed to GetLatestRecord(): %v", err)
		}

		recordId, err := uc.CreateRecord(scheduleId, nowDoing)
		if (err != nil) != wantErr {
			t.Errorf("failed to CreateRecord(): %v, wantErr %v", err, wantErr)
		}

		record, err := uc.RecordRepo.GetRecordById(recordId)
		if err != nil {
			t.Errorf("failed to GetRecordById(): %v", err)
		}

		assertEqual(t, record.ScheduleId, scheduleId)
		assertEqual(t, record.Items[len(record.Items)-1].Title, nowDoing)
		assertEqual(t, latestRecord.Items[len(latestRecord.Items)-1].End, record.Items[0].Start)
		assertEqual(t, record.Items[len(record.Items)-1].End.IsZero(), true)
	})

	t.Run("should fail when given invalid input", func(t *testing.T) {
		type args struct {
			scheduleId string
			nowDoing   string
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "invalid scheuldId",
				args: args{
					scheduleId: "invalid_scheduleId",
					nowDoing:   "valid_todo_2",
				},
				wantErr: true,
			},
			{
				name: "invalid Todo",
				args: args{
					scheduleId: "valid_scheduleId",
					nowDoing:   "invalid_todo",
				},
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
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
							Start: time.Now(),
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
				_, err := uc.CreateRecord(tt.args.scheduleId, tt.args.nowDoing)
				if (err != nil) != tt.wantErr {
					t.Errorf("failed to CreateRecord(): %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})
}
