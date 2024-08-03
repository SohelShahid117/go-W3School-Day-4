// go_conditions,switch Statement,Go For Loops
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
	fmt.Printf("%v\n", x)

	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there.")
	} else {
		fmt.Println("It is cold out there.")
	}

	time2 := 22
	if time2 < 10 {
		fmt.Println("Good morning.")
	} else if time2 < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	a3 := 10
	b3 := 10
	if a3 < b3 {
		fmt.Println("a is less than b.")
	} else if a3 > b3 {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}

	x5 := 21
	if x5 >= 10 {
		fmt.Println("x is larger than or equal to 10.")
	} else if x5 > 20 {
		fmt.Println("x is larger than 20.")
	} else {
		fmt.Println("x is less than 10.")
	}

	num := 12
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		} else {
			fmt.Println("Num is also less than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

	/*
		The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement.
	*/

	// day := 4
	day := 8

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}

	//The Multi-case switch Statement

	day_ := 7

	switch day_ {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	/*
		Go For Loops
			Syntax:
				for statement1; statement2; statement3 {
				   // code to be executed for each iteration
				}

			statement1: Initializes the loop counter value.

			statement2: Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

			statement3: Increases the loop counter value.
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}

	/*
		The continue Statement

		The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
	*/

	for i := 0; i < 10; i++ {
		if i == 5 || i == 6 || i == 7 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	var adj = [2]string{"big", "tasty"}
	var fruits = [3]string{"apple", "orange", "banana"}

	for i := 0; i < len(adj); i++ {

		println("nested loop", i, "len", len(adj))

		for j := 0; j < len(fruits); j++ {
			fmt.Println(j, adj[i], fruits[j], len(fruits))
		}
	}

	/*
		The Range Keyword

		The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

		The range keyword is used like this:
	*/

	fruits2 := [3]string{"apple", "orange", "banana"}

	for idx, val := range fruits2 {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for i := 0; i < len(fruits2); i++ {
		fmt.Println(fruits2[i])
	}

	//Tip: To only show the value or the index, you can omit the other output using an underscore (_).

	fruits3 := [3]string{"apple", "orange", "banana"}
	for indx := range fruits3 {
		fmt.Printf("%v\n", indx)
	}

	for _, val := range fruits3 {
		fmt.Printf("%v\n", val)
	}

}
