package usecase

import (
	"reflect"
	"scheduler/entities"
	"testing"
)

func TestGetCurrentSchedule(t *testing.T) {
	uc := &UseCase{
		newMockScheduleRepo(),
		newMockRecordRepo(),
	}
	tests := []struct {
		name string
		want bool
		wantErr bool
	} {
		{
			name: "test get current schedule",
			want: true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.GetCurrentSchedule()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentSchedule() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if (!reflect.DeepEqual(got, entities.RemainSchedule{})) != tt.want {
				t.Errorf("GetCurrentSchedule() = %v, want = %v", got, tt.want)
			}
		})
	}
}
