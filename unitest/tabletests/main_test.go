package tabletests

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test2",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test3",
			args: args{
				arr: []int{3, 1, 2},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
