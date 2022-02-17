package functions

import (
	"fmt"
	"time"
)

//MyFunction
type MyFunction func(string)

// MiddlerwareLog
func MiddlerwareLog(f MyFunction) MyFunction {
	return func(name string) {
		fmt.Println(time.Now().Format("2006-01-02"))
		f(name)
	}
}
