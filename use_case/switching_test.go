package usecase

import (
	"testing"
)

func TestSwitching(t *testing.T) {
	uc := &UseCase{
		newMockScheduleRepo(),
		newMockRecordRepo(),
	}
	type args struct {
		recordId string
		switchTo string
	}
	tests := []struct {
		name string
		args args
		want bool
		wantErr bool
	} {
		{
			name: "test switching with correct recordId",
			args: args{"0", "Gaming"},
			want: true,
			wantErr: false,
		},
		{
			name: "test switching with correct recordId and correct todo",
			args: args{"0", "Relax"},
			want: true,
			wantErr: false,
		},
		{
			name: "test switching with incorrect recordId",
			args: args{"1", "Gaming"},
			want: false,
			wantErr: true,
		},
		{
			name: "test switching with correct recordId and incorrect todo",
			args: args{"0", "Cooking"},
			want: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := uc.Switching(tt.args.recordId, tt.args.switchTo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchedule() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}
