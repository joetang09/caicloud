package caicloud

import (
	"testing"
)

func Test_solution(t *testing.T) {

	genFunc := func(num, at int) (arr []int) {
		for i := 0; i < num; i++ {
			if i <= at {
				arr = append(arr, at-i)
			} else {
				arr = append(arr, at+i)
			}
		}
		return
	}

	var (
		testData = []struct {
			d []int
			a int
		}{
			{
				[]int{1, 2, 3, 4, 5},
				0,
			},
			{
				[]int{5, 4, 3, 2, 1},
				4,
			},
			{
				[]int{5, 4, 2, 1, 5},
				3,
			},
			{
				[]int{},
				-1,
			},
			{
				[]int{1},
				0,
			},
			{
				genFunc(1000000, 100),
				100,
			},
			{
				genFunc(10000000, 1000000),
				1000000,
			},
		}
	)

	for _, d := range testData {
		s := solution(d.d)
		if s != d.a {
			t.Errorf("array%v should %d, but %d", d.d, d.a, s)
		}
	}
}
