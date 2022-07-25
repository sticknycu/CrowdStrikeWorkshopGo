package main

import (
	// face formatarea
	"fmt"
	"math"
	"reflect"
	"time"
	//pkg "example.com/demo/examplepkg"
)

// go mod tidy ca sa-l descarce

import log "github.com/sirupsen/logrus"

func main() {
	println("Hello world!")

	fmt.Println("Hello World!", 1)

	// go build ./...
	// go run ./...

	// go test ./...
	// go test -race

	// go fmt examples.go -- pentru issues de exemplu paranteze lipsa si etc -- se ocupa de formatare spatii si diverse "infrumusetare a codului"/identare
	// go vet examples.go --- probleme pentru ampersant la scanf -- issues / typos mai cu sens looking

	fmt.Printf("%f\n", math.Sqrt(2))

	log.Info("Hello Info")

	// functions

	// ()
	// func add(x int, y int) int
	//---- func add(x, y int) int

}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func getNextPos(x, y int) (nextX, nextY int) {
	nextX = (y + 2) / 3
	nextY = x + 1
	return // daca sunt deja definite sus, nu mai zic ce returnez :D (lul)
}

/// Variables & Constants
// definire la nivel de packet
var (
	p = 2
	s = "abc"
)

const (
	pi = 1
	a  = iota // lul , iota? ce-i asta mor
	b
	c
)

func variables() {
	// keyword var
	//var i int
	//var i, j, k int
	var i, j, k = 1, 2, 3
	//k := 3
	v1, v2, v3 := 42, false, "no!"

	println(i, j, k, v1, v2, v3)

	// keyword const
	const Pi = 3.14
	// high-precision values
	// multiple constants are usually in a group like imports
	// no short form

	// la nivel de functie
	var (
		r = 2
		l = "abc"
	)

	println(r, l)
}

// Basic types

// bool , string, int int8 int16 int32 int64, uint uint8 uint16 uint32 uint64 uintptr
// byte == aluas for uint8 -- char?
// rune == alias for int32 -- represents a Unicode code point

// float32 float64
// complex64 complex128

// int, uint
// zero values
// 0 for numeric types
// false for boolean type, and "" for strings
// nil for pointers -- strings are NEVER nil (bruh nil meme lmaooo)
// i := 42 // defaulted int
// f := float64(i)

var (
	i = 333
	f = float64(34.5436) // cast to float64
)

func g(i int64) int64 {
	return i * i
}

// %v -- tip de formatare care ne face efectiv valoarea -- din orice afiseaza string
// %q -- string quote ""

// putem face _ = f -- ca sa-l ignore ca sa nu mai tipe compilatorul ca nu e folosit

func typees() {
	f := 3.14
	println(reflect.TypeOf(f))
}

// Loops
// generic keyword for
// keywords break and continue
func loops(stop bool) {
	var s = 0
	for i := 0; i < 10 && !stop; i++ {
		s += i
		if i > 10 {
			stop = true
		}
	}

	for !stop {
		s += i
		if s > 10 {
			stop = true
			break
		}
		i++
	}

	e1 := 0
	for j := 0; j < 2; j++ {
		//Loop1:
		for !stop {
			e1 += i
			if e1 > 10 {
				continue
				//				stop = true
				//				break Loop1
			}
			i++
		}
	}
}

// if & else

func ifff() {
	/*if v := executeSmth(); v < lim {
		return v
	}*/
}

// switch
func swwwt() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tommorow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// defer -- urmareste pe dos call stack-ul
func f2() {
	defer fmt.Println("abc")
	fmt.Println("def")
}

func f3() {
	f2()
}

// Pointers and arrays

func ppp() {
	//var p *int

	// Operator & for generating pointers
	// Operator * for retrieving the value of a pointer
	// Zero value nil
	// NO pointer arithmertic
	// if p != nil {
	//	i := p
	// }
	// we get an explicit error for bad deferences: panic
	// demo!!!
	/*var pp *int
	var ppp int
	pp & ppp

	fmt.Println(pp, *pp)*/
}

/// Arrays & slices
func arrr() {
	//var ar [10] int
	// array length is part of the type (cannot be resized)
	// accesing elements: i := a[2]
	// dynamically sized views: slices

	//var b []int
	// slices are formed by specifying limits: a[low : high]
	// we can omit one or both of these limits
	// slices are like references to arrays -> changes to a slice modify the array
	// length and capacity len(s), cap(s)
	// zero value nil
	// demo

	a := [3]int{1, 2, 3}
	fmt.Println(a)

	aa := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", aa)

	s := a[0:2]
	// sau avem a[:3] , a[:], a[0:]
	fmt.Println(s) // [low, high)
	// slices - pointeri , nu copii
	fmt.Println("%v %v\n", len(s), cap(s))

	var s1 []string
	fmt.Println("%v", s1) // nil

	if s1 == nil {
		println("nil")
	}
}

// Constructing slices dynamically : make & append
func dinslices() {
	// a := make([]int, 5) --- len(a) = 5
	// b := make([]int, 0, 5) --- len(b) = 0, cap(b) = 5
	// we have a "constructor" built in function to declare slices
	// no need to import anything
	// the second slice here is filled with zero values for its type
	// adding values to a slice is done with another built in function
	// a = append(a, 1)
	// b = append(b, 2, 3, 4)
	// b = append(b, a...) // appends the whole list
	var s2 = make([]string, 10)
	for i := 0; i < 10; i++ {
		s2 = append(s2, fmt.Sprintf("%d", i))

	}

	fmt.Println(s2)

	var s3 = make([]string, 5)
	for i := 0; i < len(s3); i++ {
		s3[i] = fmt.Sprintf("%d ", i)

	}

	fmt.Println(s3, cap(s3)) // cap(s3) = 5 -- nr elemente, nu cata memorie ocupa (avem "0 1 2 3 4 ")

	var s4 = make([]string, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, fmt.Sprintf("%d", i))
	}

	fmt.Println(s3, cap(s3)) // cap(s3) = 20 -- face realocare dinamica automat

}

/// Iterating over arrays made easy range

// the range keyword returns the index and a 'copy' of the value
// we can omit either one of them
// for i := range pow
// for index, val := range pow

func itt() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Println("@**%s = %d\n", i, v)
	}

	for index, value := range pow {
		fmt.Println(index, value)
	}

	str := "abcdef"
	for _, v := range str {
		// uint32, uint8, se face el frumos cumva
		fmt.Println(int(v), reflect.TypeOf(v))
	}

	// sau casting to byte and after we iterate
	for _, v := range []byte(str) {
		// uint32, uint8, se face el frumos cumva
		fmt.Println(int(v), reflect.TypeOf(v))
	}

	// defining types
	type abc int32
	// si putem crea tipuri si crea operatii
	var v1 abc
	fmt.Println(reflect.TypeOf(v1))
	/*for _, v := range v1 {
		str1 += v.(abc)
	}*/
}
