package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)


const shortCodeLength = 7

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	sb := strings.Builder{}
	sb.Grow(shortCodeLength)
	for i := 0; i < shortCodeLength; i++ {
		sb.WriteByte(charSet[seededRand.Intn(len(charSet))])
	}
	return sb.String()
}

func main () {
	
	fmt.Println("Enter the URL to shorten:")
	var url string 
	fmt.Scanln(&url)

	shortCode := generateShortCode()
	fmt.Printf("Shortened URL: http://short.url/%s\n", shortCode)


}

