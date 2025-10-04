// init() is Go’s way of running setup code automatically before main(), without you explicitly calling it.
// No Arguments, No Return Values
// Automatically Executed
package main

import "fmt"

var msg string
func init() {
	// initialization code (automatically executes before main)

	msg="Hello world"

}
func main() {

	fmt.Println(msg) //Hello world

}