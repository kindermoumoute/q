package q

import "fmt"

// Q for the Quine
func Q() {
	fmt.Print(`package main;import."github.com/kindermoumoute/q";func main(){Q()}`)
}
