package main

import "fmt"
import "unsafe"

func DataTypes() {
	// Int types set to max integer value of 9,223,372,036,854,775,807
	t_int := 9223372036854775807 //defaults to either 32 or 64 bits
	t_int8 := int8(1)
	t_int16 := int16(922)
	t_int32 := int32(922)
	t_int64 := int64(922337)
	
	fmt.Println("Integer types")
	fmt.Println(t_int, t_int8, t_int16, t_int32, t_int64)
	fmt.Printf("%T %T %T %T %T\n", t_int, t_int8, t_int16, t_int32, t_int64)


	// Unsigned integer types set to max negativ value -9,223,372,036,854,775,806
	u_int := uint(1)
	//u_int8 := uint8(-9223372036854775806)
	//u_int16 := uint16(-9223372036854775806)
	//u_int32 := uint32(-9223372036854775806)
	//u_int64 := uint64(-9223372036854775806)
	//u_intptr := uintptr(-9223372036854775806)
	// These do not compule due to overflow 

        u_int8 := uint8(9)
        u_int16 := uint16(9)
        u_int32 := uint32(9)
        u_int64 := uint64(9)
        u_intptr := uintptr(9) 

	fmt.Println("Unsignes integer types")
	fmt.Println(u_int, u_int8, u_int16, u_int32, u_int64, u_intptr)
	fmt.Printf("%T %T %T %T %T %T\n", u_int, u_int8, u_int16, u_int32, u_int64, u_intptr)

	x := 5

	fmt.Println("\n Pointers \n")
	x_add := uintptr(unsafe.Pointer(&x))
	var x_pointer *int
	x_pointer = &x 

	fmt.Printf("Variable x has value: %v and type: %T\n", x, x)
        fmt.Printf("Variable x_add has value: %v and type: %T\n", x_add, x_add)
	*x_pointer = *x_pointer + 1
	fmt.Printf("Variable x has value: %v and type: %T\n", x, x)
        fmt.Printf("Variable x_pointer has value: %v and type: %T\n", x_pointer, x_pointer)

}
