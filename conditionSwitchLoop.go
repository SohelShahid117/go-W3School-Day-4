// go_conditions
package main

import "fmt"

func main() {
	x := 5
	y := 7
	z := 2
	fmt.Println(x > y)
	fmt.Println(x != y)

	fmt.Println((x > y) && (y > z))

	a := false
	fmt.Println((x <= y) || a)

	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	// x1 := 20
	// y1 := 18
	var x1 int = 20
	var y1 int = 18
	if x1 > y1 {
		fmt.Println("x1=%v is greater than y1=%v", x1, y1)
		fmt.Printf("x1=%v is greater than y1=%v\n", x1, y1)
	}
	fmt.Println("%v", x)
	fmt.Printf("%v", x)
}
