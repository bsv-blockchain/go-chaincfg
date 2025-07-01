package template_test

import (
	"fmt"

	"github.com/bsv-blockchain/go-chaincfg"
)

// ExampleGreet demonstrates the usage of the Greet function.
func ExampleGreet() {
	msg := template.Greet("Alice")
	fmt.Println(msg)
	// Output: Hello Alice
}
