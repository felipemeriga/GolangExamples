package example

import (
	"fmt"
)

func passSliceByReference(myRefenceSlice *[]int) {
	*myRefenceSlice = append(*myRefenceSlice, 10)
}

func dontPassReference(myPassedSlice []int) {
	myPassedSlice = append(myPassedSlice, 5)
}

// SliceNotPassByReference - When you send a slice do another function, it doesn't pass by reference, Golang don't
// pass by reference, so if you change the copy of the slice inside the called function, the original slice will remain
// untouched.
func SliceNotPassByReference() {
	mySlice := []int{1, 2, 3, 4}
	dontPassReference(mySlice)
	fmt.Println(mySlice)
}

// SlicePassByReferenceAddress - Here we are going to send the reference address of the slice, and change its content,
// and as we can see, the original slice will be changed.
func SlicePassByReferenceAddress() {
	mySlice := []int{1, 2, 3}
	passSliceByReference(&mySlice)

	fmt.Println(mySlice)
}

// ChangeArrayIndexByReference - Slices/Arrays are contiguous memory addresses, which means you can get the reference(address)
// of on of the indexes, and change its content.
func ChangeArrayIndexByReference() {
	var mySlice = []int{1, 2, 3, 4}
	firstNumber := &mySlice[0]
	*firstNumber = 10
	fmt.Println(mySlice)
}

func changeMap(myMap map[string]string) {
	myMap["1"] = "1"

}

// MapExample - Different from another data structures, maps and channels in Golang are automatically created as pointers
// under the hood, so that is why when you send a map through another function, if you change it, the original one will be also
// changed. This means that Golang can pass by reference? Not really, because on compile time, under the hood a pointer is created,
// which means that changeMap function is actually receiving the reference address of the map, you just don't need to define it as a pointer,
// to make the syntax easier and more user-friendly.
func MapExample() {
	myMap := make(map[string]string)
	changeMap(myMap)
	fmt.Println(myMap)
}
