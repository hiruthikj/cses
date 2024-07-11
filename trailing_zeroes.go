package main

import (
	"fmt"

)

func trailing_zeroes() {
	var n int
	fmt.Scanf("%d\n", &n)

	factor_5_count := 0
	divisor := 5

	for divisor <= n {
		factor_5_count += n/divisor
		divisor *= 5
	}

	fmt.Println(factor_5_count)
}
