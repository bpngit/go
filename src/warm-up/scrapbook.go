package main

import (
	"fmt"
	"math"
	"runtime"
)

var scala, java, python bool

func main() {
	//const s:="Bipin" does not work
	s := "Bipin"
	//const  s string = "constant"  error as s is already declared
	const constant string = "constant"

	fmt.Println("value of s => ", s)
	fmt.Println("value of constant => ", constant)
	fmt.Println("value of Pi=>", 22/7)
	fmt.Println("sin value 30 =>", math.Sin(30))
	//wrong
	/*
	   for i:=0; i<= 3 {
	       fmt.Println(i)

	   }
	*/

	for i := 0; i <= 3; i++ {
		fmt.Println(i)

	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	k := 0
	for k < 3 { // for is now act as while.
		fmt.Println("value of k =>", k)
		k++ // loop does not terminate if the value of k is not incremented.
	}

	for {
		fmt.Println("test break.. print only once")
		break
	}
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum <1000 in for loop =>", sum)

	if 8%4 == 0 {
		fmt.Println("int%int is mod operation,8%4 is 0")
	}
	//const n int =10
	if n := 10; n > 11 { // variable declared in if is avilable in all branches.
		fmt.Println(" if block if n>11 where n =10")
	} else {
		fmt.Println("else block if n>11 where n =10")
	}

	i := 10
	for j := 1; j <= 3; j++ { // j=0 will bring integer divide by 0 error.
		k = i % j
		fmt.Println("value of switch =>", k)
		switch k {
		case 0:
			fmt.Println(" print 0") // expoerted methods starts with capital letter name.
		case 1:
			fmt.Println("print 1")
		case 2:
			fmt.Println("print2")
		case 3:
			fmt.Println("print 3")
		default:
			fmt.Println("default case")
		}
	}

	fmt.Println(g) // local function name starts with small letter .
	fmt.Println("Add 5 and 6 =>", add(5, 6))
	a, b := swap("Hello", "world")
	fmt.Println(a, b)
	fmt.Println(nakedFuntionSplit(8))
	//fmt.Println(i, scala, python, java)
	fmt.Println("accessing global var", 0, scala)
	h, j := 1, 2
	l := 3.514
	fmt.Println(h, j)
	fmt.Println("Check type of var by using %T")
	fmt.Printf("h is of type %T\n", h)
	fmt.Printf("k is of type %T\n", l)

	//switch statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s,", os)
	}
	// pointers

	aa, bb := 42, 2701
	fmt.Println(bb)
	p := &aa        // pointer of i
	fmt.Println(*p) // pointer of P which has pointer of i
	*p = 21         // p's pointer is 21 now
	fmt.Println(aa)
	p = &bb      // point to bb
	*p = *p / 37 // divide bb through the pointer
	fmt.Println(bb)

}

func g() string {
	return "I am go laguage"
}
func add(a, b int) int { // return type and parameter type comes after varuable. Refer Martin Oderskey comment along with stackexchange.com
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func nakedFuntionSplit(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	fmt.Print("I am naked thus not returning any variable, but mapped to arguments")
	return
}
