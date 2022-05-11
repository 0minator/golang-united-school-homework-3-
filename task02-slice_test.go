package homework

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		input []int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int64
	}{
		{
			name: "",
			args: struct{ input []int64 }{
				input: []int64{4, 2, 3, 5},
			},
			wantResult: []int64{5, 3, 2, 4},
		},
		{
			name: "",
			args: struct{ input []int64 }{
				input: []int64{4, 3, 2, 1, 0, 5, 10, 9, 8, 7, 6},
			},
			wantResult: []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name: "",
			args: struct{ input []int64 }{
				input: []int64{5, 1, 4, 1, 3, 9, 5, 3, 5, 6, 2},
			},
			wantResult: []int64{5, 3, 5, 6, 2, 9, 5, 1, 4, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := reverse(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("reverse() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
