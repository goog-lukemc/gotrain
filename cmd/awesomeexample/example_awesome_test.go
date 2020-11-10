package awesomeexample

import (
	"fmt"
)

func ExampleHello() {
	name := Hello("Alex")
	fmt.Println(name)

	// Output: Hello, 👋 Alex
}

func ExampleHelloEverybody() {
	greeted := HelloEverybody([]string{"Alex", "Luke"})
	for _, s := range greeted {
		fmt.Println(s)
	}

	// Unordered output:
	// Hello, 👋 Luke
	// Hello, 👋 Alex writes buggy code
}

// Exercise 5 minutes:
// Run the tests above with `go test` and fix the broken example
// Run `go get -v  golang.org/x/tools/cmd/godoc` and then `godoc -http=:8080`
// Navigate to localhost:8080 and go to the third party section to check out our docs for awesomeexample

// BONUS:

// Testable examples in go
// https://blog.golang.org/examples

// godoc tips and tricks
// https://medium.com/@elliotchance/godoc-tips-tricks-cda6571549b
