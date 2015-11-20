// https://golang.org/doc/code.html
package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"github.com/ludovic/stringutil"
)

func main() {
	var result string
	testCase := 5

	switch testCase {
		case 0 : result = HelloWorld()
		case 1 : result = HelloWorldReverse()
		case 2 : result = Slices()
		case 3 : result = SliceOnArray()
		case 4 : result = Sum(5, 5, 10)
		case 5 : result = Choose()
	}

	Show(result)
}

func HelloWorld() (message string){
	message = strings.Join(os.Args[1:], " ")
	if len(message) == 0 { message = "Hello, world." }
	message = fmt.Sprintf("HelloWorld: %s", message)
	return
}

func HelloWorldReverse() (message string){
	message = strings.Join(os.Args[1:], " ")
	if len(message) == 0 { message = "Hello, world." }
	message = fmt.Sprintf( "HelloWorldReverse: %s", stringutil.Reverse(message) )
	return
}

func Slices() string{
	//message := make([]string, 5, 5)
	message := []string{"h","e","l","l","o"}
	return fmt.Sprintf("Slices [2:5] from %s : %s", message, message[2:5])
}

func SliceOnArray() string{
	myArray := [...]string{"h","e","l","l","o"} // ... -> No need to give lenght
	message := myArray[:4] // Make a slice pointing to myArray at point 0 len 4
	myArray[0] = "w"
	return fmt.Sprintf("SliceOnArray hello->[h e l l]->%s", message)
}

// Variadic Function
func Sum(numbers ...int) (result string){
	total := 0
	for _, num := range numbers {
		total += num
	}
	result = fmt.Sprintf("%d", total)
	return
}

func Choose() (result string){
	if len(os.Args) == 1 {
		return "Enter a digit (eg: 1 or 5 or 115 or -5)"
	}

	value, err := strconv.Atoi(os.Args[1])
	if err != nil {
		result = "'" + os.Args[1] + "' is not a digit !"
	} else {
		msg := ChooseMsg()
		switch {
			case value > 30 : result = msg(">", 30)
			case 21 <= value && value <= 30 : result = msg("[-]", 20, 30)
			case 11 <= value && value <= 20 : result = msg("[-]", 10, 20)
			case  0 <= value && value <= 10 : result = msg("<", 10)
			case value < -20 : result += msg("--", 20) ; fallthrough
			case value < -10 : result += msg("--", 10) ; fallthrough
		default:
			result += msg("Neg", value)
		}
	}
	return
}

// Higher-Order Function, used by Choose (and variadic with T: interface)
func ChooseMsg() func(code string, values ...interface{}) string{
	msg := map[string]string{
		"<": "Inférieur ou égal à %d",
		"[-]": "Compris entre %d et %d",
		">": "Supérieur à %d",
		"--": "Inférieur à %d et...",
		"Neg": "Un chiffre négatif. On s'en fou du quel (%d) !",
	}

	// Return a closure
	return func(code string, values ...interface{}) string{
		return fmt.Sprintf(msg[code], values...) // values -> value1, value2, value3
	}
}

func Show(result string) {
	if( len(result) == 0 ){
		result = "No result"
	}
	fmt.Printf( "%s\nEnd.\n", result )
}
