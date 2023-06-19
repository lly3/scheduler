package usecase

import "testing"

func TestCreateRecord(t *testing.T) {
	uc := &UseCase{
		newMockScheduleRepo(),
		newMockRecordRepo(),
	}
	type args struct {
		scheduleId string
		nowDoing string
	}
	tests := []struct {
		name string
		args args
		want bool
		wantErr bool
	} {
		{
			name: "test create record with correct scheduleId",
			args: args{"0", "Gaming"},
			want: true,
			wantErr: false,
		},
		{
			name: "test create record with incorrect todo",
			args: args{"0", "Cooking"},
			want: false,
			wantErr: true,
		},
		{
			name: "test create record with incorrect scheduleId",
			args: args{"1", "Gaming"},
			want: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.CreateRecord(tt.args.scheduleId, tt.args.nowDoing)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRecord() error = %v, wantErr = %v", err, tt.wantErr)
				return
			} 
			if (len(got) != 0) != tt.want {
				t.Errorf("CreateRecord() = %v, want = %v", got, tt.want)
			}
		})
	}
}
