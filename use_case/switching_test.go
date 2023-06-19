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
		switchTo string
	}
	tests := []struct {
		name string
		args args
		want bool
		wantErr bool
	} {
		{
			name: "test switching",
			args: args{"Gaming"},
			want: true,
			wantErr: false,
		},
		{
			name: "test switching with correct todo",
			args: args{"Relax"},
			want: true,
			wantErr: false,
		},
		{
			name: "test switching with incorrect todo",
			args: args{"Cooking"},
			want: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := uc.Switching(tt.args.switchTo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Switching() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}
