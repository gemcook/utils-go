package timeutil

import (
	"reflect"
	"testing"
	"time"
)

func TestParseDBDatetime(t *testing.T) {
	type args struct {
		dbTimestamp string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"1", args{"2000-01-01 01:02:03"}, time.Date(2000, 1, 1, 1, 2, 3, 0, time.Local), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDBDatetime(tt.args.dbTimestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDBDatetime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDBDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustParseDBDatetime(t *testing.T) {
	type args struct {
		dbTimestamp string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"1", args{"2000-01-01 01:02:03"}, time.Date(2000, 1, 1, 1, 2, 3, 0, time.Local)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustParseDBDatetime(tt.args.dbTimestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustParseDBDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}
