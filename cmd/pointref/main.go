package main

import (
	"fmt"
)

func main() {
	// Same as slice := []int{1, 2, 3, 4, 5}
	// Only for you understanding: A slice reference an array, always.
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[:]

	fmt.Println("Declare a slice in this way : slice := []int{1, 2, 3, 4, 5}\n")
	fmt.Printf(Msg("slice","As all composed type, reference an object. Here: array."), slice, slice)
	fmt.Printf(Msg("array","Is a fixed length array, referenced by slice."), array, &array)
	fmt.Printf(" We work on a part of the array (array[3:]) -> Address: %p\n", array[3:])

	// Pass the slice BY COPY: cp_slice and slice are two distincts objects,
	// but they reference the same array.
	ChangeByCopy(slice)
	fmt.Printf(Msg(" slice","Array is changed. Not the slice itself."), slice, slice)

	// Pass the slice BY REFERENCE create a pointer to passed object: p_slice is pointing slice.
	// ... And slice is referencing array.
	ChangeByRef(&slice)
	fmt.Printf(Msg(" slice","Slice and Array are changed."), slice, slice)
}

func Msg(obj, msg string) string {
	return fmt.Sprintln(obj, ": %#v (Address %p) -", msg)
}

func ChangeByCopy(cp_slice []int) {
	// Normal way pass parameters by copy: copy_of_slice is a copy of slice passed to.
	// copy_of_slice and the original slice are referencing the same object: array
	// So, you can change the array but not the slice itself.
	// Be aware: append or shrink array will produce a new one, not referenced by the original slice.
	cp_slice = cp_slice[3:] // Get a new slice to the partial array.
	cp_slice[0] = 25
	cp_slice[1] = 38
	fmt.Printf(Msg("\nChangeByCopy(slice)", "Slice passed by copy but reference same array."), cp_slice, cp_slice)
}

func ChangeByRef(p_slice *[]int) {
	// Because we use pointer, we can shrink or append data to
	// the array and send-back to pointer the result.
	slice := *p_slice // Dereferencing the pointer to get the pointed object (a convenient way)

	// Only for you understanding. Same as: *p_slice = (*p_slice)[3:]
	slice = slice[3:] // 1: Get a new slice to the partial array.
	*p_slice = slice  // 2: Store into pointed object the new slice : Original slice is changed.

	slice[0] = 185
	slice[1] = 175
	fmt.Printf(Msg("\nChangeByRef(&slice)", "Slice passed by reference can be changed, and the array too."), p_slice, *p_slice)
}
