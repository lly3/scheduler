package usecase

import (
	"reflect"
	"scheduler/entities"
	"testing"
)

func TestGetSchedule(t *testing.T) {
	uc := &UseCase{
		newMockScheduleRepo(),
		newMockRecordRepo(),
	}
	type args struct {
		recordId string
	}
	tests := []struct {
		name string
		args args
		want bool
		wantErr bool
	} {
		{
			name: "test get schedule with correct recordId",
			args: args{"0"},
			want: true,
			wantErr: false,
		},
		{
			name: "test get schedule with incorrect recordId",
			args: args{"1"},
			want: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.GetSchedule(tt.args.recordId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchedule() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if (!reflect.DeepEqual(got, entities.RemainSchedule{})) != tt.want {
				t.Errorf("GetSchedule() = %v, want = %v", got, tt.want)
			}
		})
	}
}
