// https://golang.org/doc/code.html
package main

import (
	"fmt"
	"github.com/ludovic/stringutil"
	"os"
	"strconv"
	"strings"
	"reflect"
)

func main() {
	var result string
	testCase := 16

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
	case 13:
		result = IndirectAccess()
	case 14:
		BlockScope()
	case 15:
		result = EmbeddedType()
	case 16:
		result = ErrorGenerator()
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
	var results = make([]string, 9) // Slice of 9 empty string, return it's reference
	// reference is like a pointer, but the address is hidden

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

// As any parameter, methods parameters are copy of original object.
// It's possible to create setter methdos by using Pointer:
func (dog *Dog) setAge(age int) int {
	dog.age = age
	return dog.age
}

// Interface that ensure methods definition. Used by Interfaces function.
type Animal interface {
	madeSound() string
	setAge(newage int) int // Because setAge use pointer, any object of Animal MUST be a pointer.
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

func (cat *Cat) setAge(age int) int {
	cat.age = age
	return cat.age
}

// Interface sample
func Interfaces() (result string) {
	// Because we rely on animal and one of its methods use pointer, any
	// object must be a reference. Animal, as interface, is more
	// focused on methods, this is no tricky at this point of view.
	animals := []Animal{
		&Dog{name: "SnoopyTheDog", age: 7},
		&Cat{name: "SweetyTheCat", age: 2, color: "white"},
		&Dog{name: "DingoCartoonDog", age: 75},
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
func Setter() (result string) {
	dog := Dog{name: "SnoopyTheDog", age: 7}
	if age := dog.setAge(55); age > 10 { // Test with block var into it
		result = fmt.Sprintf("Age supérieur à 10 : %d", age)
	} else {
		result = fmt.Sprintf("Age inchangé (echec) : %d", age)
	}
	return
}

// Duck typing sample
func DuckTyping() string {
	cat := Cat{name: "SweetyTheCat", age: 2, color: "white"}
	dog := Dog{name: "Snoopy", age: 5}
	number := cat.age
	str := "Un Chat"
	results := make([]string, 6)

	results[0] = WhatIAm(cat)
	results[1] = WhatIAm(number)
	results[2] = WhatIAm(str)

	fmt.Printf("Age of dog is : %d\n", dog.age)
	ChangeMe(&dog)
	fmt.Printf("Age of dog now is : %d\n", dog.age)

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
func WhatIAm(something interface{}) (result string) {
	if animal, is_animal := something.(Animal); is_animal {
		return animal.madeSound()
	}

	switch identified_object := something.(type) {
	case Cat:
		result = fmt.Sprintf("Ma couleur : %s", identified_object.color)
	case string:
		result = fmt.Sprintf("Un chaine : %s", identified_object)
	case int:
		result = fmt.Sprintf("Un chiffre : %d", identified_object)
	}
	return
}

// Use empty interface to apply Duck Typing techniques on Pointer
func ChangeMe(something interface{}) {

	switch identified_object := something.(type) {
	case *Cat: // Test Nammed Type give access fields and methods
		identified_object.color = "Black"
	case Animal: // Test interface give only access to methods
		fmt.Printf("Test Animal from Pointer: %s\n", identified_object.madeSound())
		identified_object.setAge(55)
	case *string:
		*identified_object = "Un truc bizarre !"
	case *int:
		*identified_object = *identified_object * 2
	}
}

// Sample - Pointer can throw invalid indirect access. This is why.
func IndirectAccess() (result string) {
	// Predeclared type: Boolean, numeric, string.
	// Create pointer, require * to access content because it's directly pointed by it.
	var p_direct = new(int)
	*p_direct = 8 //Point directly to content with *, direct access

	// Named type follow the same rule with predeclared type
	type Age int
	var p_age = new(Age)
	*p_age = 10 //Point directly to content with *, direct access

	// Composite types : Maps, Slice, Struct, Interface, Channel: Content access is indirect.
	// Create pointer to a composite type, acting itself like a pointer.
	// Attempt to access content with * (eg: *p_struct) throw an invalid indirect access:
	// Composite type have no data, but only pointer(s) to other type (named type,
	// predeclared type or itself composite type).
	// Always use indirect access capabilities to access
	// members and methods, even when used in Pointer like bellow :
	var p_struct = new(struct {
		name string
		age  int
	})
	p_struct.name = "Ludo" // Here, age is initialized to 0. No need *

	var p_array = new([3]string)
	p_array[0] = "Hello"
	p_array[1] = "The" // Index 2 is initialized to empty string (""). No need *

	// Named type follow the same rule with composite type
	type Name struct {
		firstname string
		lastname  string
	}
	var p_name = new(Name)
	p_name.firstname = "Ludo" //No need *

	// ... And so forth for slice and maps. Keep this in mind.

	result = fmt.Sprintf(
		" Array (Join) -> %s \n Struct -> [%s - %d] \n int -> %d\n",
		strings.Join(p_array[0:], " "),
		p_struct.name, p_struct.age,
		*p_direct,
	)

	// Here, we prove we have pointer :
	ChangeStruct(p_struct)
	ChangeArray(p_array)
	ChangeInt(p_direct)

	result = fmt.Sprintln(result, "---\n", fmt.Sprintf(
		"Array (Join) -> %s \n Struct -> [%s - %d] \n int -> %d",
		strings.Join(p_array[0:], " "),
		p_struct.name, p_struct.age,
		*p_direct,
	),
	)
	return
}

// Used by sample Indirect Access. No need * to access content.
func ChangeStruct(p *struct {
	name string
	age  int
}) {
	p.age = 18
}

// Used by sample Indirect Access. No need * to access content.
func ChangeArray(p *[3]string) {
	p[2] = "World !"
}

// Used by sample Indirect Access. Here, point directly to content with *
func ChangeInt(p *int) {
	*p = 8996
}

// Sample: Block and scopes
func BlockScope() {

	result := 0 // Here, result MUST BE an INTEGER !!! (See shadowing method)
	var should_be int
	should_be = should_be + 0

	fmt.Printf("My number is: %d\n", result)
	{ // Context created with {} or block (like Java or C) :
		// Variable created inside block exist only inside it.
		// You can use any variable already created before your block (result here).
		value, _ := strconv.Atoi("185") // value is created here, scopped into this block.
		result = value                  // result existe before, so it's linked to it.
		fmt.Printf("Local scopped value with simple block: %d (forget outside scope)\n", value)
	}
	fmt.Printf("Change number 1: %d\n", result)

	// Same apply to if {}, for {}, func {}, etc... Here, sample with if {}
	if value, err := strconv.Atoi("208"); err == nil {
		fmt.Printf("Local scopped value with if block: %d (forget outside scope)\n", value)
		result = result + value // We change the result again,
	} else { // but value and err will be lost outside this scope.
		fmt.Println("This is not a value")
	}
	fmt.Printf("Change number 2: %d\n", result)

	// Like C (and unlike local variable in Java), you have shadowing in Go.
	// Use of shadow is natural, but be aware of this.
	message := "Change number 3"
	{
		// We shadow external variables message and result. To keep access, use Pointer
		p_int := &result                                      // Local pointer to external variable 'result' (int)
		value := 118                                          // Local
		sum := *p_int                                         // Local (copy content of pointer p_int)
		message := fmt.Sprint(message, " (shadow method) : ") // Shadowed.
		result := ""                                          // Shadowed.

		// In this reusable code (sick !), we have already this code.
		// We do not follow the previous types for the same names, they are shadowed.
		sum = sum + value
		result = fmt.Sprint(message, sum)
		fmt.Println(result)

		// Now, we update external value 'result' through Pointer p_int.
		// Without that, somme will be lost. And result will never been updated.
		*p_int = sum
	}
	// Outside this scope, we use our variable like we do before:
	value := 810                                                      // Here value is not already declared, thanks to scopes.
	result = result + value                                           // result (int)
	message = fmt.Sprintf("%s (outside scope) : %d", message, result) // Message (not modified before)
	result = 10
	value = 0
	fmt.Println(message)

	// Why shadow can be TRICKY ! Don't forget, any {} is scopped, even if{}, for{} and so forth.
	// Look at this code. You have an error into, but nor the compiler nor you notice it.
	// We have reseted before result to 10 and value to 0. Think we have complex calculs :)
	if 0 <= value && value <= 150 {
		should_be := ( (5 * 2 * 2) - 10 + 1 * ( value ) )
		if should_be >= 10 {
			// Now, it should be 11.
			should_be++
		} else {
			// Never enter in this test. Normal.
			should_be--
		}
	}
	// Show the bad result ! Wich variable did you think we use in your overall if statment ?
	fmt.Println(
		"You result in this tricky exemple is wrong : Expected 11, but result is",
		should_be,
		// Why the result is 0 ??? Error is in the calcul ?
		// This is the bad question. Question is : Why this is not a compil error.
	)
	// -> It's look like a beginner bug. Yes, it is. Variable should_be was declared
	// a the start of this sample. Never used. And a newcomer in Go will write
	// this code. Without an already declared should_be variable, this code
	// produce a complil error on the Println (should_be not declared).
	// But this is not the case, and our developper
	// dive into calcul because he think error come from it.
	// The reality, he use ":=" operator (line 2) that initate a new variable. Because
	// we are in a scope (if {}), we shadow initial "should_be" variable with
	// a new one and use it into the overall if statement. But this variable
	// is lost outside the "if". And we fall in this case in our "just" initialized
	// "should_be" variable as we have at the start: To 0. Funny, yah ?


	return
}

// Sample : Embedded Type into other Type
func EmbeddedType() (result string){
	type Human struct{
		age int
		color string
	}

	type Worker struct{Human; name string} // Embbed Human at the root
	georges := Worker{Human: Human{age: 18, color: "white"}, name: "Georges"}

	result = fmt.Sprintf(
		"As a Worker: I'm %s, a'm %d years old and my color is %s\n",
		georges.name,
		georges.age,
		georges.color,
	)


	// Shadow color. No ambiguities when initialized
	type Cyborg struct{Human; name string; color int} // Embbed Human at the root
	machine := Cyborg{Human: Human{age: 3}, name: "Machine", color: 4485}

	// But when used, you get accessor of Cyborg (int), not Human (string)
	result += fmt.Sprintf(
		"As a Cyborg: I'm %s, a'm %d years old and my color reference is %d\n",
		machine.name,
		machine.age,
		machine.color,
	)


	type Driver struct{race Human; name string} // Embbed Human into race accessor
	mickael := Driver{race: Human{age: 32, color: "black"}, name: "Mickael"}

	result += fmt.Sprintf(
		"As a Driver: I'm %s, a'm %d years old and my color is %s\n",
		mickael.name,
		mickael.race.age,
		mickael.race.color,
	)

	return
}

// Use by ErrorGenerator: Custom Type to handle Error.
type cnxError struct {
	error, // Embed type: Predeclared interface in the universe package.
	message string
	code int
}

// Use by ErrorGenerator: Error method required by the interface "error".
func(e cnxError) Error() string {
	// When you want a printable version, Error() is called for any object
	// of error interface or wich embed it.
	return fmt.Sprintf("You got an error (%d): %s", e.code, e.message)
}

// Use by ErrorGenerator: Factory to build cnxError objects. Not required.
func CnxError(message string, code int) *cnxError {
	return &cnxError{message: message, code: code} // return reference, not a copie
}

// Sample ErrorGenerator
func ErrorGenerator() (result string) {
	// Not required. To show type returned
	show_me := func(object interface{}) string {
		return fmt.Sprintf("Type retourné : %s\n", reflect.TypeOf(object))
	}

	// By using custom type wich implement interface 'error'
	err := CnxError("Pas de connexion", 500)
	result += fmt.Sprintln(err) // -> You got an error (500): Pas de connexion
	result += fmt.Sprintf("Code de l'erreur : %d\n", err.code)
	result += show_me(err)

	// Same, but without encapsulation / not DRY
	err_with_fmt := fmt.Errorf("You got an error (%d): %s", 404, "Access denied")
	result += fmt.Sprintln(err_with_fmt)
	result += fmt.Sprintln("Code de l'erreur : Il faut parser le texte !!!")
	result += show_me(err_with_fmt)

	return
}

func Show(result string) {
	if len(result) == 0 {
		result = "No result"
	}
	fmt.Printf("%s\nEnd.\n", result)
}
