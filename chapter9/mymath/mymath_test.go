package mymath

import "testing"

type tcase_t struct {
	vals []float64
	exp float64
}

var avr_set = []tcase_t {
	{[]float64{1,2,3,4,5}, 3},
	{[]float64{1,1,1,1,1}, 1},
	{[]float64{100}, 100},
}

var min_set = []tcase_t {
	{[]float64{1,5,55,-11, 0.333}, -11},
	{[]float64{1}, 1},
	{[]float64{1.5, 5.5, 11, 0.333}, 0.333},
}

var max_set = []tcase_t {
	{[]float64{1,5, 55, -11, 0.333}, 55},
	{[]float64{1}, 1},
	{[]float64{1.5, 5.5, 11, 0.333}, 11},
}

func TestAverage(t * testing.T) {
	for _, tcase := range avr_set  {
		v := Average(tcase.vals)
		if (v != tcase.exp) {
			t.Error("For", tcase.vals,
					"expected", tcase.exp,
					"got", v)
		}
	}
}

func TestMin(t * testing.T) {
	for _, tcase := range min_set {
		v := Min(tcase.vals)
		if (v != tcase.exp) {
			t.Error("For", tcase.vals,
					"expected", tcase.exp,
					"got", v)
		}
	}
}

func TestMax(t * testing.T) {
	for _, tcase := range max_set {
		v := Max(tcase.vals)
		if v != tcase.exp {
			t.Error("For", tcase.vals,
					"expected", tcase.exp,
					"got", v)
		}
	}
}