package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
	"runtime"
	"strings"
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

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func DeferFunc() {
	// defer
	defer fmt.Println("world")
	fmt.Println("hello")

	// stacking defer
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

type Vertex struct {
	X int
	Y int
}

var (
	v3 = Vertex{1, 2}  // has type Vertex
	v4 = Vertex{X: 1}  // Y:0 is implicit
	v5 = Vertex{}      // X:0 and Y:0
	p2 = &Vertex{1, 2} // has type *Vertex
)

// making slices
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// range
var pow2 = []int{1, 2, 4, 8, 16, 32, 64, 128}

//slicing exercise
func Pic(dx, dy int) [][]uint8 {
	outer := make([][]uint8, dy)
	for i := range outer {
		outer[i] = make([]uint8, dx)
		for j := range outer[i] {
			outer[i][j] = uint8(j) // also try (x+y)/2, x*y, and x^y. !Wow! :-)
		}
	}
	return outer
}

//maps
type LocationCoordinate struct {
	Lat, Long float64
}

var map1 map[string]LocationCoordinate

var map2 = map[string]LocationCoordinate{
	"Bell Labs": LocationCoordinate{
		40.68433, -74.39967,
	},
	"Google": LocationCoordinate{
		37.42202, -122.08408,
	},
	"Apple": {37.42202, -122.08408},
}

//map exercise
func WordCount(s string) map[string]int {
	var map4 = make(map[string]int)

	for i, v := range strings.Fields(s) {
		if map4[v] != 0 {
			continue
		}
		fmt.Println("Looping", i)
		var count int = 0
		for _, word := range strings.Fields(s) {
			if word == v {
				count++
			}
		}
		map4[v] = count
	}
	return map4
}

//closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

var (
	previous, current int
)

func fibonacci() func() int {
	return func() int {
		sum := previous + current
		if sum == 0 {
			previous = 0
			current = 1
			return previous + current
		} else {
			previous = current
			current = sum
			return current
		}

	}
}

type FloatVertex struct {
	X, Y float64
}

func (v *FloatVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *FloatVertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//interface
type Abser interface {
	Abs() float64
}

func runInterface() {
	var abser Abser
	f2 := MyFloat(-math.Sqrt2)
	v9 := FloatVertex{3, 4}

	abser = f2  // a MyFloat implements Abser
	abser = &v9 // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// abser = v9

	fmt.Println(abser.Abs())
}

// implicit interface
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func runImplicitInterface() {
	fmt.Println("Implicit interface")
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")

	person := Person{"Arthur Dent", 42}
	anotherPerson := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(person, anotherPerson)
}

//stringer

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func runStringer() {
	fmt.Println("stringer---")
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func runErrors() {
	fmt.Println("errors")
	if err := run(); err != nil {
		fmt.Println(err)
	}
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

	v2 := 1.3 // change me!
	fmt.Printf("v is of type %T\n", v2)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// For is Go's "while"
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)

	//switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today, time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	DeferFunc()
	// end of flow control statements

	// pointers
	fmt.Println("Pointers ---")
	i, j := 42, 2701
	var p *int      // declaring a pointer to an int.
	p = &i          // point to i
	fmt.Println(*p) // read i through the pointer. This is known as "dereferencing" or "indirecting".
	*p = 21         // set i through the pointer. This is known as "dereferencing" or "indirecting".
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j.
	*p = *p / 37   // divide j through the pointer. This is known as "dereferencing" or "indirecting".
	fmt.Println(j) // see the new value of j
	fmt.Println("Pointers ---")

	// Structs
	fmt.Println(Vertex{1, 2})

	v1 := Vertex{1, 2}
	v1.X = 4
	fmt.Println(v1.X)
	p1 := &v1
	p1.X = 1e9
	fmt.Println(v1)

	fmt.Println(v3, p2, v4, v5)

	//arrays
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	// slicing
	fmt.Println("s[1:4] ==", s[1:4])

	// missing low index implies 0
	fmt.Println("s[:3] ==", s[:3])

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", s[4:])

	a1 := make([]int, 5)
	printSlice("a1", a1)
	b1 := make([]int, 0, 5)
	printSlice("b1", b1)
	c1 := b1[:2]
	printSlice("c1", c1)
	d1 := c1[2:5]
	printSlice("d1", d1)
	//Nil slices
	var zNil []int
	fmt.Println(zNil, len(zNil), cap(zNil))
	if zNil == nil {
		fmt.Println("nil!")
	}

	var a2 []int
	printSlice("a2", a2)

	// append works on nil slices.
	a2 = append(a2, 0)
	printSlice("a2", a2)

	// the slice grows as needed.
	a2 = append(a2, 1)
	printSlice("a2", a2)

	// we can add more than one element at a time.
	a2 = append(a2, 2, 3, 4)
	printSlice("a2", a2)

	// range
	for i, v := range pow2 {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow3 := make([]int, 10)
	for i := range pow3 {
		pow3[i] = 1 << uint(i)
	}
	for _, value := range pow3 {
		fmt.Printf("%d\n", value)
	}

	pic.Show(Pic)

	// Maps
	map1 = make(map[string]LocationCoordinate)
	map1["Bell Labs"] = LocationCoordinate{
		40.68433, -74.39967,
	}
	fmt.Println(map1["Bell Labs"])
	fmt.Println(map1)

	fmt.Println(map2)

	//Mutating Maps
	map3 := make(map[string]int)

	map3["Answer"] = 42
	fmt.Println("The value:", map3["Answer"])

	v6, ok1 := map3["Answer"]
	fmt.Println("The value:", v6, "Present?", ok1)

	map3["Answer"] = 48
	fmt.Println("The value:", map3["Answer"])

	delete(map3, "Answer")
	fmt.Println("The value:", map3["Answer"])

	v6, ok2 := map3["Answer"]
	fmt.Println("The value:", v6, "Present?", ok2)

	// map exercise
	wc.Test(WordCount)

	//functions arevalues too
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

	v7 := &FloatVertex{3, 4}
	fmt.Println("FloatVertex", v7.Abs())

	f1 := MyFloat(-math.Sqrt2)
	fmt.Println(f1.Abs())

	v8 := &FloatVertex{3, 4}
	v8.Scale(5)
	fmt.Println(v8, v8.Abs())

	runInterface()

	runImplicitInterface()
	runStringer()
	runErrors()
}
