package parser

import (
	"reflect"
	"testing"
	"time"
)

func TestParseUntisTime(t *testing.T) {
	type args struct {
		untisTime string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"Valid date", args{untisTime: "7.2.2019 8:09"}, time.Date(2019, 2, 7, 8, 9, 0, 0, time.UTC), false},
		{"Zeroed date", args{untisTime: "07.02.2019 08:09"}, time.Date(2019, 2, 7, 8, 9, 0, 0, time.UTC), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUntisTime(tt.args.untisTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUntisTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUntisTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUntisDate(t *testing.T) {
	type args struct {
		untisDate string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"Valid date", args{untisDate: "Vertretungen  7.2. / Donnerstag"}, time.Date(2019, 2, 7, 0, 0, 0, 0, time.UTC), false},
		{"Valid date 2", args{untisDate: "Vertretungen  8.2. / Freitag"}, time.Date(2019, 2, 8, 0, 0, 0, 0, time.UTC), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUntisDate(tt.args.untisDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUntisDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUntisDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
