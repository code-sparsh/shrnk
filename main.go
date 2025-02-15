package main

import (
	"fmt"
	"github.com/code-sparsh/shrnk/utils"
)


func main () {
	
	fmt.Println("Enter the URL to shorten:")
	var url string 
	fmt.Scanln(&url)

	// store := utils.URLStore{
	// 	urls: make(map[string]string),
	// }

	store := utils.NewURLStore()

	shortCode, err := store.StoreURL(url)

	if err != nil {
		fmt.Printf("Failed to store URL: %v\n", err)
		return
	}
	fmt.Printf("Shortened URL: http://short.url/%s\n", shortCode)

}



