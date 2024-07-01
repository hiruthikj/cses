package main

import (
	"fmt"
)

func number_spiral() {
	var n int
	fmt.Scanf("%d\n", &n)
	
	var x,y int
	for i:=0; i<n; i++ {
		fmt.Scanf("%d %d\n", &y, &x)
		var start int
		if y > x {
			start = y%2 + (y - (y%2)) * (y - (y%2))   
			if y%2 == 0 {
				start -= x-1
			} else {
				start += x-1
			}
		} else {
			start = (x-1+(x%2))*(x-1+(x%2)) + 1 - (x%2)
			if x%2 == 0 {
				start += y-1
			} else {
				start -= y-1
			}
		}
		fmt.Println(start)
	}
}
