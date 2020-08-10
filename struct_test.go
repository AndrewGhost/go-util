package util

import "testing"

// go test -v ./struct_test.go struct.go -test.run TestSetStructFields
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
			name: "set-struct-fields",
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
			t.Logf("new struct = %+v", tt.args.s)
		})
	}
}
