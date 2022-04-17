package main

import "fmt"

func main() {
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte -alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	// go에서 변수 만들고 안쓰면 에러남
	// var name = "Brad"

	// shorthand
	name := "Brad" // 함수 안에서만 이런 방식 사용 가능
	size := 1.3    // float로 들어가는데 디폴트는 float 64임 // var size float32 = 2.3
	// email := "syh30@gmail.com"

	name, email := "Brad", "syh30@gmail.com"

	var age int32 = 37 // var age = 37
	var isCool = true  // var 대신 const 쓰면 에러남
	isCool = false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
}
