package utils

import (
	"reflect"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name    string
		dataSet [][]bool
		want    [][]bool
	}{
		{
			name: "test expectation for example",
			dataSet: [][]bool{
				{false, false, false, false, false, false},
				{false, false, false, false, false, false},
				{false, false, true, true, true, false},
				{false, true, true, true, false, false},
				{false, false, false, false, false, false},
				{false, false, false, false, false, false},
			},
			want: [][]bool{
				{false, false, false, false, false, false},
				{false, false, false, true, false, false},
				{false, true, false, false, true, false},
				{false, true, false, false, true, false},
				{false, false, true, false, false, false},
				{false, false, false, false, false, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Evaluate(tt.dataSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
