package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type products struct {
}

var fileName = "base.txt"

// Show list with products
func showListProducts() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error to open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println()
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		// Delete "Space" in the elements
		for i, v := range parts {
			parts[i] = strings.TrimSpace(v)
		}

		// Checking, it the date i correct
		if len(parts) != 3 {
			fmt.Printf("String №%s isn't correct. Please cheak and corrert it\n", parts[0])
			continue
		}
		fmt.Printf("№%s. Product: %s. Macros: %s\n", parts[0], parts[1], parts[2])

	}
	fmt.Println()

	// Check on errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//

func findProductByNumber() {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error to open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Print("Enter the number of the product you want to find: ")
	//	var productNumber = "1" // temp
	//fmt.Fscan(os.Stdin, &productNumber)

	//

	//

	//

	//

	//

	//

	// Check maybe errors, if they'll be
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
