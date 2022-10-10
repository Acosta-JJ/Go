package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	var amount int64

	fmt.Printf("Please introduce the total number of intergers that will sort and created\n")
	fmt.Scan(&amount)
	result := generateRandomNumbers(amount)
	f, err := os.Create("numbers.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.Write(result)

	if err != nil {
		log.Fatal(err)
	}
}
