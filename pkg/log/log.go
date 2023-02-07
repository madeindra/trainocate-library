package log

import "fmt"

func Print(msg string, data interface{}) {
	fmt.Printf("%v: %v", msg, data)
}
