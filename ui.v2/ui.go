package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChooseSample(samples []string) (stop bool, sampleNo int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Numero d'exemple, (q)uiter, (c)lear_screen ou (l)ister les exemlpes : ")
	response, _ := reader.ReadString('\n')

	response = strings.Replace(response, "\n", "", -1)
	if response == "" {
		return
	}

	var cmd string
	cmd = strings.ToLower(response[0:1])

	switch cmd {
	case "q":
		fmt.Println("Bye !")
		stop = true
	case "c":
		ClearScreen()
	default:
		if i, err := strconv.Atoi(response); err == nil {
			sampleNo = i
			fmt.Println("Sample choosen : ", sampleNo)
		} else {
			fmt.Println("Please, choose a number.")
		}
	}

	return
}
