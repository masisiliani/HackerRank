package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	WritePhrase("Hello, World.")
	inputString, err := GetPhrase()
	if err != nil {
		fmt.Println("There is an error to get the phrase. Err:", err)
	}
	WritePhrase(inputString)
}

//WritePhrase Write a phrase on terminal
func WritePhrase(strPhrase string) {
	fmt.Println(strPhrase)
}

//GetPhrase scans a phrase by terminal
func GetPhrase() (string, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()
	err := scanner.Err()

	return inputString, err
}
