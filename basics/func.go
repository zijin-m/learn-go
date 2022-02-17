package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func main()  {
	fmt.Println(add2(42,13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	
}
