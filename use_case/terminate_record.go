package usecase

import "time"

func (uc *UseCase) TerminateRecord() error {
	record, err := uc.RecordRepo.GetLatestRecord()
	if err != nil {
		return err
	}

	record.Items[len(record.Items)-1].End = time.Now()

	if err := uc.RecordRepo.Update(record); err != nil {
		return err
	}

	return nil
}
