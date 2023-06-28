package usecase

import (
	"scheduler/entities"
	"testing"
	"time"
)

func TestTerminateRecord(t *testing.T) {
	t.Run("should terminate record when latest record exist", func(t *testing.T) {
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
		uc.TerminateRecord()
		record, err := uc.RecordRepo.GetLatestRecord()
		if err != nil {
			t.Errorf("failed to GetRecordById(): %v", err)
		}
		assertEqual(t, record.Items[len(record.Items)-1].End.IsZero(), false)
	})

	t.Run("should fail to terminate record when latest record not exist", func(t *testing.T) {
		uc := &UseCase{
			newMockScheduleRepo(),
			newMockRecordRepo(),
		}
		uc.TerminateRecord()
		_, err := uc.RecordRepo.GetLatestRecord()
		if err == nil {
			t.Errorf("should failed to terminate record due latest record doesn't exist")
		}
	})
}
