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

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
)


func main()  {
	// fmt.Println(add2(42,13))
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	// var i, j = 1, 2
	// k := 3
	// c, python, java := true, false, "no!"
	// fmt.Println(i, j, k, c, python, java)

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Println(math.MaxInt32)
	// fmt.Println(math.MaxInt64)
	// fmt.Println(math.MaxInt)
	// fmt.Println(math.MaxFloat32)
	// fmt.Println(math.MaxFloat64)

	var x, y = 3, 4
	f := float64(x)
	n := string(y)
	fmt.Println(f)
	fmt.Println(n)
}
