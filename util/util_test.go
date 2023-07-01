package util

import "testing"

func TestGetRssiPercentage(t *testing.T) {
	type args struct {
		rssi int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				rssi: -90,
			},
			want: 38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRssiPercentage(tt.args.rssi); got != tt.want {
				t.Errorf("GetRssiPercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}
