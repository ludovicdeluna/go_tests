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
	testCase := 7

	switch testCase {
	case 0 : result = HelloWorld()
	case 1 : result = HelloWorldReverse()
	case 2 : result = Slices()
	case 3 : result = SliceOnArray()
	case 4 : result = Sum(5, 5, 10)
	case 5 : result = Choose()
	case 6 : result = Pointers()
	case 7 : result = Label()
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
		default: result += msg("Neg", value)
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

func Pointers() string{
	// Thanks to GC, Go has no Arithmetic Pointer
	var variable int
	var pointer = new(int) // Our pointer
	var results = make([]string, 9) // Slice of 9 empty string

	variable = 255
	*pointer = 10 // *pointer to act on its content and not on pointer itself

	results[0] = fmt.Sprint( "a/ Pointer - Address : ", pointer )
	results[1] = fmt.Sprint( "a/ Pointer - Content : ", *pointer )
	results[2] = fmt.Sprint( "a/ -- Call by Ref (*) -- : ", PassByReference(pointer) )
	results[3] = fmt.Sprint( "a/ Pointer - After call : ", *pointer )

	// Pass by copy is the default parameter mechanism for functions in Go
	results[4] = fmt.Sprint( "b/ Variable - Content : ", variable )
	results[5] = fmt.Sprint( "b/ -- Call by Copy -- : ", PassByCopy(variable) )
	results[6] = fmt.Sprint( "b/ Variable - After call : ", variable )

	results[7] = fmt.Sprint( "c/ -- Call by Ref (&)-- : ", PassByReference(&variable) )
	results[8] = fmt.Sprint( "c/ Variable - After call : ", variable )

	return strings.Join(results, "\n")
}

// Sample used by Pointers - Signature of this function require a Pointer T int
func PassByReference(value *int) int{
	*value = *value + 18
	return *value
}

// Sample used by Pointers - Signature of this function require an Integer
func PassByCopy(value int) int{
	value = value + 18
	return value
}

func Label() (result string){
	num, err := strconv.Atoi( strings.Join(os.Args[1:], "") )
	if err != nil {goto ShowError} // goto or (in loop): continue / break

	if num > 10 {
		result = "Un chiffre au dessus de 10"
	}else{
		result = "Un chiffre en dessous de 10 ou négatif"
	}
	return

	ShowError:
	return "Saisir un chiffre"
}


func Show(result string) {
	if( len(result) == 0 ){
		result = "No result"
	}
	fmt.Printf( "%s\nEnd.\n", result )
}
