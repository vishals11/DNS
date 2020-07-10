package location

import (
	"testing"

	"DroneNavigationSystem/location/model"
)

func Test_service_GetLocation(t *testing.T) {
	type fields struct {
		SectorID float64
	}
	type args struct {
		request model.LocateDataBankRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "Test #1",
			fields:fields{SectorID:54},
			args:args{request: model.LocateDataBankRequest{X: "12", Y: "30", Z: "40", Vel: "120"}},
			want: 4548,
		},
		{
			name: "Test #2",
			fields:fields{SectorID:0},
			args:args{request: model.LocateDataBankRequest{X: "12", Y: "30", Z: "40", Vel: "120"}},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				SectorID: tt.fields.SectorID,
			}
			if got, _ := s.GetLocation(tt.args.request); got != tt.want {
				t.Errorf("GetLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}