package functions

import (
	"fmt"
	"time"
)

// MiddlerwareLog
func MiddlerwareLog(f func(string)) func(string) {
	return func(name string) {
		fmt.Println(time.Now())
		f(name)
	}
}
