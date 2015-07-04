package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //naked return
}

//1
var c1, python1, java1 bool

//2
var i2, j2 int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z4     complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
	fmt.Println(add1(42, 13))
	fmt.Println(add2(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	//1
	var i1 int
	fmt.Println(i1, c1, python1, java1)
	//2
	var c2, python2, java2 = true, false, "no!"
	fmt.Println(i2, j2, c2, python2, java2)

	//3
	var i3, j3 int = 1, 2
	k3 := 3
	c3, python3, java3 := true, false, "no!"

	fmt.Println(i3, j3, k3, c3, python3, java3)

	const f4 = "%T(%v)\n"
	fmt.Printf(f4, ToBe, ToBe)
	fmt.Printf(f4, MaxInt, MaxInt)
	fmt.Printf(f4, z4, z4)

	// data types
	// bool

	// string

	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr

	// byte // alias for uint8

	// rune // alias for int32
	//      // represents a Unicode code point

	// float32 float64

	// complex64 complex128

	// default value
	var i5 int
	var f5 float64
	var b5 bool
	var s5 string
	fmt.Printf("%v %v %v %q\n", i5, f5, b5, s5)

	var x6, y6 int = 3, 4
	var f6 float64 = math.Sqrt(float64(x6*x6 + y6*y6))
	var z6 int = int(f6)
	fmt.Println(x6, y6, z6)

	v := 1.3 // change me!
	fmt.Printf("v is of type %T\n", v)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
