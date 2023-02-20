package main

import (
	"errors"
	"fmt"
	"time"
)

// p7 - Functions
func func_name(a, b int, c, d float32, s string) (int, float64, string) {
	a = a + b
	res := float64(c) + float64(d)
	s = "This is s: " + s
	return a, res, s
}

// p9 - pointer
func p_sum(a, b *int) *int {
	*a = *a + *b
	return a
}

// p14 - interface
type Shape interface {
	area() float64
}
type Circ struct {
	raidus float64
}
type Rect struct {
	width, height float64
}

func (shape Circ) area() float64 {
	return shape.raidus * shape.raidus * 3.14
}
func (shape Rect) area() float64 {
	return shape.width * shape.height
}

// p15 - error handling
func checknum(ip *int) (int, error) {
	if ip == nil {
		return 0, errors.New("Nil Address!")
	}
	return *ip, nil
}

// p16 - concurrency
func routine1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "Good"
}

func routine2(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "Morning"
}

// p17 - defer
func checkDefer() {
	fmt.Println("From defer: init")
	defer fmt.Println("Hello 1")
	defer fmt.Println("Hello 2")
	defer fmt.Println("Hello 3")
	fmt.Println("From defer: close")
	return
}

func main() {
	// Simple hello world
	fmt.Print("Hello World\n")

	// p1 - basic syntax
	var a, b, c int
	a = 10
	b = 11
	c = 12
	a = b + c
	fmt.Printf("This is a: %d\n", a)

	// p2 - data types - e1 - from 0 to max
	var ui8 uint8 = 255
	// var ui16 uint = 65535
	// var ui32 uint32 = 4294967295
	// var ui64 uint64 = 18446744073709551615
	fmt.Println(ui8)

	// p2 - data types - e2
	// var i8 int8 = -128
	// var i16 int16 = -32768
	// var i32 int32 = -2147483648
	// var i64 int64 = -9223372036854775808

	// p2 - data types - e3
	// var f32 float32 = 0.1
	// var f64 float64 = 0.1
	var c64 complex64 = 0.1 + 0.1i // two f32
	// var c128 complex128 = 0.1 + 0.1i // two f64
	fmt.Println("Complex ", c64)

	// p2 - data types - e4
	// var ui8 byte = 255
	// var i32 rune = -2147483648
	// var ui uint // uint32 or uint64
	// var i int // int32 or int64
	// uintptr uip var // an unsigned integer to store the uninterpreted bits of a pointer value

	// p3 - variable definition or like above
	d := 255.1
	fmt.Printf("this is a num %.2f with type %T\n", d, d)
	var v1, v2 = 3, "foo"
	fmt.Println(v1, " ", v2)

	// p4 - Integer Literals
	// const binary int = 0b11111111;
	// const octal int = 0377;
	// const decimal int = 255;
	// const hex int = 0xEF;
	// const i_val = 30;
	// const ui_val = 30u;
	// const l_val = 30l;
	// const ul_val = 30ul;

	// p5 - Operators (Bitwise, rest is like C/C++)
	const p uint8 = 0b10100111
	const q uint8 = 0b11101111
	var r uint8 = (p | q) // & is and, | is or, ^ is xor
	r = r << 2
	fmt.Printf("This is p : %b\nThis is q : %b\nThis is r : %b\n", p, q, r)

	// p6 - Decision Making (if/else like c,)
	var x uint8 = 3
	// We can have a value assignmemt at the enterance
	// also no need for break in each case!
	switch x = 2; x {
	case 1:
		fmt.Println("is one")
	case 2:
		fmt.Println("is two")
	case 3:
		fmt.Println("is three")
	default:
		fmt.Println("more than 3")
	}
	// Also there is a select which I have to review in a seperate file!
	// select is like switch but is used for channel!

	// p7 - functions
	out1, _, out3 := func_name(1, 2, 0.1, 0.1, "Hello") // ignore out2
	fmt.Println(out1, " ", out3)

	// p8 - arrays
	var myArr1 = [5]int{1, 2, 3}
	var myArr2 [3][2][4]float32
	var myArr3 = []uint8{1, 3, 5}
	fmt.Println(myArr1)
	fmt.Println(myArr2)
	fmt.Println(myArr3[2])

	// p9 - pointers (use it like c)
	var ip *int = &a
	var nip **int = nil
	nip = &ip
	var arrp [10]*int
	arrp[0] = &a
	arrp[1] = &b
	_ = p_sum(arrp[0], arrp[1])
	fmt.Printf("this is the address of a : %x\n", ip)
	fmt.Printf("this is the address of nip : %x\n", nip)
	fmt.Printf("a after p_sum is : %d\n", a)

	// p10 - structure
	type book struct {
		number int
		name   string
		price  float64
	}
	var b1 book
	var ib1 *book
	ib1 = &b1
	b1.name = "Hello"
	b1.number = 1
	b1.price = 1.32
	fmt.Println(*ib1)

	// p11 - slices
	var s1 []int
	if s1 == nil {
		fmt.Println("Nil slice")
	}
	s1 = make([]int, 4)
	var s2 []int = make([]int, len(s1)*2)
	copy(s2, s1)
	s1 = append(s1, 12)
	fmt.Println("S2 is : ", s2)

	// p12 - range
	var range_array []int = make([]int, 10)
	for i := range range_array {
		fmt.Println("This is from range : ", i)
	}

	// p13 - map
	// mp1 := map[string]string{"ali": "krm"}
	mp1 := make(map[string]string)
	mp1["kali"] = "bidin"
	mp1["malo"] = "badin"
	for fn, dn := range mp1 {
		fmt.Println(fn, " ", dn)
	}

	// p14 - interface
	var myCirc Circ
	myCirc.raidus = 1.5
	var myRect Rect
	myRect.width = 2
	myRect.height = 3
	fmt.Println("Circ: ", myCirc.area(), " - Rect: ", myRect.area())

	// p15 - error handling
	val, err := checknum(nil)
	if err == nil {
		fmt.Println("Value:", val)
	} else {
		fmt.Println("err: ", err)
	}

	// p16 - concurrency (Goroutines + con)
	myChannel := make(chan string)
	go routine1(myChannel)
	go routine2(myChannel)
	res := <-myChannel
	fmt.Println("Channel res is random: ", res)
	// close(myChannel)
	cc1 := make(chan string)
	cc2 := make(chan string)
	go routine1(cc1)
	go routine2(cc2)
	// here we have the first available (which is random again :D)
	select {
	case op1 := <-cc1:
		fmt.Println("From cc1 : ", op1)
	case op2 := <-cc2:
		fmt.Println("From cc2 : ", op2)
	}

	// p17 - deffer
	checkDefer()

}
