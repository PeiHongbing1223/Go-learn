package main

import (
	"fmt"

	"github.com/PeiHongbing1223/Go-learn/go_dev_doc/Tutorial_Getting_started/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
