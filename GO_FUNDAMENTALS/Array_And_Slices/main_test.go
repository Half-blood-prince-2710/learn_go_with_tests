package main

import (
	"reflect"
	"slices"
	"testing"
)


func TestSum(t *testing.T){
	Testcase:= []struct{
		name string
		array []int
		output int
	}{
		// {"sum of all number",[]int{1,2,3,4,5},15},
		// {"sum of empty slice",[]int{},0},
		{"sum of 1 number",[]int{2},2},
		// {"sum of negative number",[]int{-1,-2,-1,-3},-7},
	}

	for _,tc := range Testcase {
		t.Run(tc.name,func(t *testing.T) {
			got:= Sum(tc.array)
			want:= tc.output
			assertError(t,got,want, tc.array)
		})
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

func assertError(t testing.TB, got int, want int ,arr []int){
	t.Helper()
	if got != want{
		t.Errorf("got %d want %d given, %v", got, want, arr)
	}

}