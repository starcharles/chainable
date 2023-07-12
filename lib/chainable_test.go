package chaninable_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	chaninable "github.com/starcharles/chainable/lib"
)

func Test_Filter(t *testing.T) {

	type args struct {
		f      func(int) bool
		values []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test Fileter 1",
			args: args{
				f:      func(v int) bool { return v%2 == 0 },
				values: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{2, 4, 6},
		},
		{
			name: "test Filter 2",
			args: args{
				f:      func(v int) bool { return v%3 == 0 },
				values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{3, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := chaninable.NewChaninable(tt.args.values).Filter(tt.args.f)

			assert.Equal(t, result.Values(), tt.want)
		})
	}
}

// func TestMap(t *testing.T) {
// 	t.Skip()

// 	type args struct {
// 		f      func(int) int
// 		values []int
// 	}

// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		{
// 			name: "test Map",
// 			args: args{
// 				f:      func(v int) int { return v * 2 },
// 				values: []int{1, 2, 3, 4, 5},
// 			},
// 			want: []int{2, 4, 6, 8, 10},
// 		},
// 		{
// 			name: "test Map 2",
// 			args: args{
// 				f:      func(v int) int { return v * 3 },
// 				values: []int{1, 2, 3, 4, 5},
// 			},
// 			want: []int{3, 6, 9, 12, 15},
// 		},
// 	}

// 	// for _, tt := range tests {
// 	// 	t.Run(tt.name, func(t *testing.T) {
// 	// 		result := chaninable.NewChaninable[int](tt.args.values).Map(func(v int) int {
// 	// 			return v * 2
// 	// 		})

// 	// 		if got := result.Map(tt.args.f); !reflect.DeepEqual(got, tt.want) {
// 	// 			t.Errorf("Set.Map() = %v, want %v", got, tt.want)
// 	// 		}
// 	// 	})
// 	// }
// }
