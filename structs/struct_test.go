package structs

import (
	"reflect"
	"testing"
)

// go test -v ./struct_test.go structs.go -test.run TestSetStructFields
func TestSetStructFields(t *testing.T) {
	type foo struct {
		Name string
		Age  int64
	}

	type args struct {
		s           interface{}
		fieldValues map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "set-structs-fields",
			args: args{
				s: &foo{
					Name: "Andrew",
					Age:  26,
				},
				fieldValues: map[string]interface{}{
					"Name": "Blue",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetStructFields(tt.args.s, tt.args.fieldValues); (err != nil) != tt.wantErr {
				t.Errorf("SetStructFields() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("new structs = %+v", tt.args.s)
		})
	}
}

// go test -v ./struct_test.go struct.go -test.run TestConvertToMap
func TestConvertToMap(t *testing.T) {
	type args struct {
		s   interface{}
		key string
	}
	type User struct {
		Id   uint64
		Name string
	}
	tests := []struct {
		name string
		args args
		want map[interface{}]interface{}
	}{
		{
			name: "convertMap",
			args: args{
				s: []User{
					{
						Id:   1,
						Name: "blue",
					},
					{
						Id:   2,
						Name: "crank",
					},
				},
				key: "Id",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToMap(tt.args.s, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Log(got[uint64(2)].(User))
				t.Errorf("convertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
