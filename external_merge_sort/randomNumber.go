package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(amount int64) []byte {
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, amount)
	rand.Read(token)
	fmt.Println(token)
	return token
}
