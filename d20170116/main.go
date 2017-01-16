package main

import "fmt"

//CtCI 8.1
func main() {
	fmt.Println(dp_solve(5))
}

func solve(n int) int {
	if n < 0 {
		return 1
	}
	num := solve(n - 1)
	if n >= 2 {
		num += solve(n - 2)
	}
	if n >= 3 {
		num += solve(n - 3)
	}
	return num
}

func dp_solve(n int) int {
	a := make([]int, n+1)
	a[1] = 1
	a[2] = 2
	a[3] = 4
	for i := 4; i <= n; i++ {
		a[i] = a[i-3] + a[i-2] + a[i-1]
	}
	fmt.Println(a)
	return a[n]
}
