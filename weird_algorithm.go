package main

import (
  "fmt"
)

func WeirdAlgorithm() {
  var n int
  fmt.Scanf("%d\n", &n)
  printConjecture(n)
  fmt.Print("\n")
}

func printConjecture(n int) {
  fmt.Printf("%d ", n)
  if n == 1 {
    return
  }
  if n%2 == 0 {
	  n /= 2
  } else {
	  n = n*3 + 1
  }
  printConjecture(n)
}
