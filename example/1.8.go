package example

import (
	"fmt"
	"github.com/fatih/structs"
	"sort"
)

// -------------------------------  Sorting in Go  ----------------------------------

type Person struct {
	Name string
	Age  int
}

func Sorting() {
	family := []Person{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})

	sort.Slice(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
}

// -------------------------------------------------------------------------------

// ------------------------------- Transforming Structs to map[string]interface{} using "github.com/fatih/structs"  ----------------------------------

func TransformStructs() {
	myself := &Person{
		Name: "Felipe",
		Age:  30,
	}

	myselfMap := structs.Map(myself)
	fmt.Println(myselfMap["Name"])
}

// -------------------------------------------------------------------------------

// ------------------------------- Generics  ----------------------------------

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func GenericsMap[K comparable, V any](m map[K]V) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func GenericArrayPrint[T any](list []T) {
	for _, v := range list {
		fmt.Println(v)
	}
}

type FlyingBird struct {
	TypeOfWing string
}

type TerrestrialBird struct {
	GroundType string
}

type Bird[T any] struct {
	Name   string
	Family *T
}

func GenericStruct() {
	sparrow := Bird[FlyingBird]{
		Name: "Sparrow",
		Family: &FlyingBird{
			TypeOfWing: "curved",
		},
	}
	penguin := Bird[TerrestrialBird]{
		Name:   "penguin",
		Family: &TerrestrialBird{},
	}
	fmt.Println(sparrow, penguin)
}

// -------------------------------------------------------------------------------
