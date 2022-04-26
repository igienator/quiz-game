package input

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetInput(message string) (string, error) {

	fmt.Print(message)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Reading user input failed")
		return "", err
	}
	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	return userInput, nil
}
