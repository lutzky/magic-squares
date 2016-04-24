package main

import (
	"reflect"
	"testing"
)

func TestAllUnique(t *testing.T) {
	testCases := []struct {
		sq            Square
		wantAllUnique bool
		wantRowsMatch bool
		wantColsMatch bool
		wantDiagMatch bool
	}{
		{sq: Square{
			Size: 3,
			Data: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
			wantAllUnique: true,
		},
		{sq: Square{
			Size: 3,
			Data: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
			wantDiagMatch: true,
			wantRowsMatch: true,
			wantColsMatch: true,
		},
		{sq: Square{
			Size: 3,
			Data: [][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 3, 2},
			},
		},
			wantRowsMatch: true,
		},
		{sq: Square{
			Size: 3,
			Data: [][]int{
				{1, 1, 4},
				{2, 2, 0},
				{3, 3, 2},
			},
		},
			wantColsMatch: true,
		},
		{sq: Square{
			Size: 3,
			Data: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
			wantDiagMatch: true,
		},
	}
	for _, tc := range testCases {
		if got := tc.sq.allUnique(); got != tc.wantAllUnique {
			t.Errorf("allUnique(%v) == %t; want %t", tc.sq, got, tc.wantAllUnique)
		}
		if got := tc.sq.rowsMatch(); got != tc.wantRowsMatch {
			t.Errorf("rowsMatch(%v) == %t; want %t", tc.sq, got, tc.wantRowsMatch)
		}
		if got := tc.sq.colsMatch(); got != tc.wantColsMatch {
			t.Errorf("ColsMatch(%v) == %t; want %t", tc.sq, got, tc.wantColsMatch)
		}
		if got := tc.sq.diagMatch(); got != tc.wantDiagMatch {
			t.Errorf("DiagMatch(%v) == %t; want %t", tc.sq, got, tc.wantDiagMatch)
		}
	}
}

func TestIncrementModulo(t *testing.T) {
	testCases := []struct{ start, want Square }{
		{
			Square{Size: 3, Data: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
			Square{Size: 3, Data: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		},
		{
			Square{Size: 3, Data: [][]int{{9, 9, 0}, {0, 0, 0}, {0, 0, 0}}},
			Square{Size: 3, Data: [][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}},
		},
		{
			Square{Size: 3, Data: [][]int{{9, 9, 9}, {0, 0, 0}, {0, 0, 0}}},
			Square{Size: 3, Data: [][]int{{0, 0, 0}, {1, 0, 0}, {0, 0, 0}}},
		},
		{
			Square{Size: 3, Data: [][]int{{0, 0, 0}, {1, 0, 0}, {0, 0, 0}}},
			Square{Size: 3, Data: [][]int{{1, 0, 0}, {1, 0, 0}, {0, 0, 0}}},
		},
	}

	for _, tc := range testCases {
		got := tc.start.incrementModulo(10)
		if !reflect.DeepEqual(got.Data, tc.want.Data) {
			t.Errorf("incremented %v, got %v; want %v", tc.start, got, tc.want)
		}
	}
}
