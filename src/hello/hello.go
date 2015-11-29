// https://golang.org/doc/code.html
package main

import (
	"fmt"
	"github.com/ludovic/stringutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result string
	testCase := 12

	switch testCase {
	case 0:
		result = HelloWorld()
	case 1:
		result = HelloWorldReverse()
	case 2:
		result = Slices()
	case 3:
		result = SliceOnArray()
	case 4:
		result = Sum(5, 5, 10)
	case 5:
		result = Choose()
	case 6:
		result = Pointers()
	case 7:
		result = Label()
	case 8:
		result = BreakOnLabel()
	case 9:
		result = Structures()
	case 10:
		result = Interfaces()
	case 11:
		result = Setter()
	case 12:
		result = DuckTyping()
	}

	Show(result)
}

func HelloWorld() (message string) {
	message = strings.Join(os.Args[1:], " ")
	if len(message) == 0 {
		message = "Hello, world."
	}
	message = fmt.Sprintf("HelloWorld: %s", message)
	return
}

func HelloWorldReverse() (message string) {
	message = strings.Join(os.Args[1:], " ")
	if len(message) == 0 {
		message = "Hello, world."
	}
	message = fmt.Sprintf("HelloWorldReverse: %s", stringutil.Reverse(message))
	return
}

func Slices() string {
	message := []string{"h", "e", "l", "l", "o"}
	return fmt.Sprintf("Slices [2:5] from %s : %s", message, message[2:5])
}

func SliceOnArray() string {
	myArray := [...]string{"h", "e", "l", "l", "o"} // ... -> No need to give lenght
	message := myArray[:4]                          // Make a slice pointing to myArray at point 0 len 4
	myArray[0] = "w"
	return fmt.Sprintf("SliceOnArray hello->[h e l l]->%s", message)
}

// Variadic Function
func Sum(numbers ...int) (result string) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	result = fmt.Sprintf("%d", total)
	return
}

func Choose() (result string) {
	if len(os.Args) == 1 {
		return "Enter a digit (eg: 1 or 5 or 115 or -5)"
	}

	value, err := strconv.Atoi(os.Args[1])
	if err != nil {
		result = "'" + os.Args[1] + "' is not a digit !"
	} else {
		msg := ChooseMsg()
		switch {
		case value > 30:
			result = msg(">", 30)
		case 21 <= value && value <= 30:
			result = msg("[-]", 20, 30)
		case 11 <= value && value <= 20:
			result = msg("[-]", 10, 20)
		case 0 <= value && value <= 10:
			result = msg("<", 10)
		case value < -20:
			result += msg("--", 20)
			fallthrough
		case value < -10:
			result += msg("--", 10)
			fallthrough
		default:
			result += msg("Neg", value)
		}
	}
	return
}

// Higher-Order Function, used by Choose (and variadic with T: interface)
func ChooseMsg() func(code string, values ...interface{}) string {
	msg := map[string]string{
		"<":   "Inférieur ou égal à %d",
		"[-]": "Compris entre %d et %d",
		">":   "Supérieur à %d",
		"--":  "Inférieur à %d et...",
		"Neg": "Un chiffre négatif. On s'en fou du quel (%d) !",
	}

	// Return a closure
	return func(code string, values ...interface{}) string {
		return fmt.Sprintf(msg[code], values...) // values -> value1, value2, value3
	}
}

func Pointers() string {
	// Thanks to GC, Go has no Arithmetic Pointer
	var variable int
	var pointer = new(int)          // Our pointer
	var results = make([]string, 9) // Slice of 9 empty string

	variable = 255
	*pointer = 10 // *pointer to act on its content and not on pointer itself

	results[0] = fmt.Sprint("a/ Pointer - Address : ", pointer)
	results[1] = fmt.Sprint("a/ Pointer - Content : ", *pointer)
	results[2] = fmt.Sprint("a/ -- Call by Ref (*) -- : ", PassByReference(pointer))
	results[3] = fmt.Sprint("a/ Pointer - After call : ", *pointer)

	// Pass by copy is the default parameter mechanism for functions in Go
	results[4] = fmt.Sprint("b/ Variable - Content : ", variable)
	results[5] = fmt.Sprint("b/ -- Call by Copy -- : ", PassByCopy(variable))
	results[6] = fmt.Sprint("b/ Variable - After call : ", variable)

	results[7] = fmt.Sprint("c/ -- Call by Ref (&)-- : ", PassByReference(&variable))
	results[8] = fmt.Sprint("c/ Variable - After call : ", variable)

	return strings.Join(results, "\n")
}

// Sample used by Pointers - Signature of this function require a Pointer T int
func PassByReference(value *int) int {
	*value = *value + 18
	return *value
}

// Sample used by Pointers - Signature of this function require an Integer
func PassByCopy(value int) int {
	value = value + 18
	return value
}

func Label() (result string) {
	num, err := strconv.Atoi(strings.Join(os.Args[1:], ""))
	if err != nil {
		goto ShowError
	}

	if num > 10 {
		result = "Un chiffre au dessus de 10"
	} else {
		result = "Un chiffre en dessous de 10 ou négatif"
	}
	return

ShowError:
	return "Saisir un chiffre"
}

func BreakOnLabel() (result string) {
	fmt.Print("Count number from 0 to 10 and print them without 5, 6, 7 and 10\n\n")
	x := 0
LoopLabel: // A Break/Continue Label, MUST be declared ahead a loop statment
	for { // Infinite loop.
		switch x {
		case 1: // Break without label exit the current statement (here: switch)
			if true == true {
				break
			}
			x = 4
		case 5: // Exit For and re-run: Jump to value 8, do not print value 5 to 7
			x = 8
			continue LoopLabel
		case 10: // Exit For and Stop: Do not print value 10
			break LoopLabel
		}
		// Code bellow will not be executed when Exit For is applied
		fmt.Printf("Show number => %d\n", x)
		x += 1
	}
	result = fmt.Sprintf("Le résultat de x est : %d", x)
	return
}

func Structures() (result string) {
	// The simplest usage of structs, with initialize of values (keys are optionals)
	cats := []struct {
		name      string
		age       int
		madeSound string
	}{
		{"SweetyTheCat", 2, "SweetyTheCat make the sound 'Miaow'"},
		{"CopyCat", 155, "CopyCat make the sound 'Miaow'"},
	} // Houch: We'r not DRY !

	// Show all cats in the loop
	for _, cat := range cats {
		fmt.Printf(
			"Name : %s, age : %d -> %s\n",
			cat.name,
			cat.age,
			cat.madeSound,
		)
	}

	// Be DRY by using custom type wich embed a method function (see bellow).
	// -> Commonly used with struct and interface, but can work with all types
	dogs := []Dog{
		{name: "SnoopyTheDog", age: 7},
		{name: "DingoCartoonDog", age: 75},
	}

	// Place all dogs into result
	for _, dog := range dogs {
		result += fmt.Sprintf(
			"Name : %s, age : %d -> %s\n",
			dog.name,
			dog.age,
			dog.madeSound(),
		)
	}

	return
}

// Named Type, used by Structures function: the T Dog
type Dog struct {
	name string
	age  int
}

// Method of Named Type MUST be declared outside functions.
// IMPOSSIBLE on Pointers and interface
func (d Dog) madeSound() string {
	return fmt.Sprintf("%s make the sound '%s' !", d.name, "Waf")
}

// Interface that ensure methods definition. Used by Interfaces function.
type Animal interface {
	madeSound() string
}

// Named Type, used by Interfaces function: the T Cat
type Cat struct {
	name  string
	age   int
	color string
}

// Method of Named Type for Cat. Used by Interfaces function.
func (c Cat) madeSound() string {
	return fmt.Sprintf("%s make a %s '%s' !", c.name, c.color, "Roroon")
}

// Interface sample
func Interfaces() (result string) {
	animals := []Animal{
		Dog{name: "SnoopyTheDog", age: 7},
		Cat{name: "SweetyTheCat", age: 2, color: "white"},
		Dog{name: "DingoCartoonDog", age: 75},
	}

	// Also possible to use in parameter (animals[0]) or (animals[0], animals[1])
	// but not just animals (an array)
	return TellWhatYouSpeak(animals...)
}

// Used by Interface function. When interface used in parameter, alway rely on its methods.
// Interface are used to inform on What this object can do. Not on what type is contained into it.
func TellWhatYouSpeak(any_animals ...Animal) string {
	var array_of_sounds = make([]string, len(any_animals))
	for index := range any_animals {
		array_of_sounds[index] = any_animals[index].madeSound()
	}
	return strings.Join(array_of_sounds, "\n")
}

// Sample: Create setter by use pointer.
func Setter() (result string){
	dog := Dog{name: "SnoopyTheDog", age: 7}
	if age := dog.setAge(55) ; age > 10 { // Test with block var into it
		result = fmt.Sprintf("Age supérieur à 10 : %d", age)
	} else {
		result = fmt.Sprintf("Age inchangé (echec) : %d", age)
	}
	return
}

// As any parameter, methods parameters are copy of original object.
// It's possible to create setter methdos by using Pointer:
func (dog *Dog) setAge(age int) int {
	dog.age = age
	return dog.age
}

// Duck typing sample
func DuckTyping() string {
	cat := Cat{name: "SweetyTheCat", age: 2, color: "white"}
	number := cat.age
	str := "Un Chat"
	results := make([]string, 6)

	results[0] = WhatIAm(cat)
	results[1] = WhatIAm(number)
	results[2] = WhatIAm(str)

	ChangeMe(&cat)
	ChangeMe(&number)
	ChangeMe(&str)
	ChangeMe(&Dog{name: "Yark yark !"}) //Empty T Dog object: Dog{}

	results[3] = WhatIAm(cat)
	results[4] = WhatIAm(number)
	results[5] = WhatIAm(str)

	return strings.Join(results, "\n")
}

// Use empty interface to apply Duck Typing techniques
func WhatIAm(something interface{}) (result string){
	if animal, is_animal := something.(Animal) ; is_animal {
		return animal.madeSound()
	}

	switch identified_object := something.(type){
	case string:
		result = fmt.Sprintf("Un chaine : %s", identified_object)
	case int:
		result = fmt.Sprintf("Un chiffre : %d", identified_object)
	}
	return
}

// Use empty interface to apply Duck Typing techniques on Pointer
func ChangeMe(something interface{}) {

	switch identified_object := something.(type){
	case *Cat: // Test Nammed Type give access fields and methods
		identified_object.color = "Black"
	case Animal: // Test interface give only access to methods
		fmt.Printf("Test Animal from Pointer: %s\n", identified_object.madeSound())
	case *string:
		*identified_object = "Un truc bizarre !"
	case *int:
		*identified_object = *identified_object * 2
	}
}

// Next see: https://golang.org/doc/effective_go.html#type_switch
// See also: https://www.golang-book.com/books/intro/9 -> Embedded Types with P

func Show(result string) {
	if len(result) == 0 {
		result = "No result"
	}
	fmt.Printf("%s\nEnd.\n", result)
}
