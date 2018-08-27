package utils

import (
	"reflect"
	"testing"

	"github.com/gemcook/pagination-go"
)

func TestConvertQueryStringToStruct(t *testing.T) {

	type TestStruct struct {
		Name string `json:"name"`
	}

	type args struct {
		queryStringParams map[string]string
		s                 interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *TestStruct
	}{
		// TODO: Add test cases.
		{"name is me", args{
			queryStringParams: map[string]string{"name": "me"},
			s:                 &TestStruct{}},
			false, &TestStruct{Name: "me"}},

		{"non-pointer-param-gets-error", args{
			queryStringParams: map[string]string{"name": "me"},
			s:                 TestStruct{}},
			true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertQueryStringToStruct(tt.args.queryStringParams, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("ConvertQueryStringToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(tt.want, tt.args.s) {
				t.Errorf("ConvertQueryStringToStruct() want = %v, but got %v", tt.want, tt.args.s)
			}
		})
	}
}

func TestConvertQueryStringToPaginationQuery(t *testing.T) {
	type args struct {
		qs map[string]string
	}
	tests := []struct {
		name string
		args args
		want *pagination.Query
	}{
		// TODO: Add test cases.
		{"simple",
			args{map[string]string{"limit": "2", "page": "1"}},
			&pagination.Query{Limit: 2, Page: 1, Sort: []*pagination.Order{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertQueryStringToPaginationQuery(tt.args.qs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertQueryStringToPaginationQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
