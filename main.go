package testmod

import (
	"fmt"
)

// Hi return a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}
