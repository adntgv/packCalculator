package main

import (
	"reflect"
	"testing"
)

var defaultPackSizes = []int{250, 500, 1000, 2000, 5000}

func Test_calculatePacks(t *testing.T) {
	type args struct {
		items     int
		packSizes []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "Test 1",
			args: args{
				items:     1,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{
				250: 1,
			},
		},
		{
			name: "Test 2",
			args: args{
				items:     251,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{
				500: 1,
			},
		},
		{
			name: "Test 3",
			args: args{
				items:     501,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			name: "Test 4",
			args: args{
				items:     12001,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
		{
			name: "Test 5 - No items",
			args: args{
				items:     0,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{},
		},
		{
			name: "Test 6 - Single pack size",
			args: args{
				items:     1000,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{
				1000: 1,
			},
		},
		{
			name: "Test 7 - Multiple packs of the same size",
			args: args{
				items:     10000,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{
				5000: 2,
			},
		},
		{
			name: "Test 9 - Smallest pack size",
			args: args{
				items:     5,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{
				250: 1,
			},
		},
		{
			name: "Test 10 - Very large pack size",
			args: args{
				items:     100000000,
				packSizes: defaultPackSizes,
			},
			want: map[int]int{
				5000: 20000,
			},
		},
		{
			name: "Test 11 - Very large pack size",
			args: args{
				items:     0,
				packSizes: []int{},
			},
			want: map[int]int{},
		},
		{
			name: "Test 12 - Very large pack size",
			args: args{
				items:     1,
				packSizes: []int{},
			},
			want: map[int]int{},
		},
		{
			name: "Test 13 - Very large pack size",
			args: args{
				items:     -1,
				packSizes: []int{},
			},
			want: map[int]int{},
		},
		{
			name: "Test 14 - Very large pack size",
			args: args{
				items:     -1,
				packSizes: []int{1},
			},
			want: map[int]int{},
		},
		{
			name: "Test 15 - Very large pack size",
			args: args{
				items:     -1,
				packSizes: []int{0},
			},
			want: map[int]int{},
		},
		{
			name: "Test 16 - Very large pack size",
			args: args{
				items:     0,
				packSizes: []int{0},
			},
			want: map[int]int{},
		},
		{
			name: "Test 17 - Very large pack size",
			args: args{
				items:     1,
				packSizes: []int{0},
			},
			want: map[int]int{},
		},
		{
			name: "Test 18 - Very large pack size",
			args: args{
				items:     -1,
				packSizes: []int{1},
			},
			want: map[int]int{},
		},
		{
			name: "Test 19 - Very large pack size",
			args: args{
				items:     0,
				packSizes: []int{1},
			},
			want: map[int]int{},
		},
		{
			name: "Test 20 - Very large pack size",
			args: args{
				items:     1,
				packSizes: []int{1},
			},
			want: map[int]int{
				1: 1,
			},
		},
		{
			name: "Test 21 - Very large pack size",
			args: args{
				items:     2,
				packSizes: []int{1},
			},
			want: map[int]int{
				1: 2,
			},
		},
		{
			name: "Test 22 - Very large pack size",
			args: args{
				items:     2,
				packSizes: []int{1, 2},
			},
			want: map[int]int{
				2: 1,
			},
		},
		{
			name: "Test 22 - Very large pack size",
			args: args{
				items:     2,
				packSizes: []int{1, 5},
			},
			want: map[int]int{
				2: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePacks(tt.args.items, tt.args.packSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nitems: %v\npacks: %v\n%v\n%v", tt.args.items, tt.args.packSizes, got, tt.want)
			}
		})
	}
}
