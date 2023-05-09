package usecase

func (uc *UseCase) GetLatestRecordId() (string, error) {

	recordId, err := uc.RecordRepo.GetLatestRecordId()
	if err != nil {
		return "", err
	}

	return recordId, nil

}
