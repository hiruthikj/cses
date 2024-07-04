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
			if y%2 == 0 {
				start = (y*y) - (x-1)
			} else {
				start = (y-1)*(y-1) + x
			}
		} else {
			if x%2 == 0 {
				start = (x-1)*(x-1) + y
				} else {
				start = x*x - y + 1
			}
		}
		fmt.Println(start)
	}
}
