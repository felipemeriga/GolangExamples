package example

import "fmt"

func PointersAndMemory() {
	stringer := "ae"
	fistPointer := &stringer
	secondPointer := fistPointer
	*secondPointer = "ae2"

	ae := make(map[string]interface{})

	ae3 := ae["test"]

	if ae["test"] == "ae" {

	}
	fmt.Println(*fistPointer)
	fmt.Println(ae3.(map[string]interface{})["ae"])

}

func AddOnArray(mySlice []int) {
	mySlice = append(mySlice, 1)

}

func TestArray() {
	var myFirstSlice []int
	AddOnArray(myFirstSlice)

	fmt.Println(myFirstSlice)
}
