package example

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func ExecuteFibo() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

type ResponseBody struct {
	CustomTime CustomTime `json:"time"`
}
type CustomTime time.Time

const ctLayout = "2006-01-02T15:04:05"

// UnmarshalJSON Parses the json string in the custom format
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = CustomTime(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *CustomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(ctLayout))
}

func RemoveRightZeros() {
	dateInEST := []byte("{ \"time\": \"2022-10-12T17:30:00-04:00\" }")
	customTime := ResponseBody{}

	err := json.Unmarshal(dateInEST, &customTime)
	if err != nil {
		log.Fatal(err)
	}

	//currentTime, _ := time.Now().UTC().MarshalText()
	//fmt.Println(string(currentTime))
	//customTime := CustomTime(time.Now())
	bytes, _ := json.Marshal(customTime)

	fmt.Println(bytes)
}
