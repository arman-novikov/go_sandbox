package mymath

import "fmt"

// returns avr of vals
func Average(vals []float64) float64 {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	vals_count := len(vals)
	total := float64(0);
	for i := 0; i < vals_count; i++ {
		total += vals[i]
	}

	return total / float64(vals_count)
}

// returns min of vals
func Min(vals []float64) float64 {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	vals_count := len(vals)
	res := vals[0]
	for i := 1; i < vals_count; i++ {
		if vals[i] < res {
			res = vals[i]
		}
	}

	return res
}

// returns max of vals
func Max(vals []float64) float64 {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	vals_count := len(vals)
	res := vals[0]
	for i := 1; i < vals_count; i++ {
		if vals[i] > res {
			res = vals[i]
		}
	}

	return res
}