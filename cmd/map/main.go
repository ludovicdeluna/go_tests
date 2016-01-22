package main

import "fmt"

type Cactus map[string]string

func (cactus Cactus) whoIAm() (answer string) {
	if name, ok := cactus["name"] ; ok {
		answer = fmt.Sprintf("Hello, cactus : %s", name)
	} else {
		answer = fmt.Sprint("Hello cactus inconnu.")
	}
	return
}

func main() {
	var cactus = make(Cactus)
	cactus["name"] = "The Big Cactus"
	cactus["country"] = "US"

	for idx := range cactus {
		fmt.Println(cactus.whoIAm())
		delete(cactus, idx)
	}

}
