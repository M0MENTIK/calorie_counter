package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Show list with products
func showListProducts() {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error to open file:", err)

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
		if len(line) == 0 {
			fmt.Printf("String isn't empty. Please cheak and corrert it\n")
			continue
		}
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

//

//

func findProductByNumber(p *Product, productNumber int) error {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error to open file:", err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()
	defer func() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		for i, v := range parts {
			parts[i] = strings.TrimSpace(v)
		}
		if len(parts) != 3 {
			return errors.New("Incorrect len of data")
		}

		// Comparison input value with value of database
		number, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("Error to conver string to int:", err)
		}
		//Если найден нужный продукт, выписываю КБЖУ, привожу к корректному виду и добавляю их в структуру
		if number == productNumber {
			macros := strings.Split(parts[2], "/")
			for i, v := range macros {
				macros[i] = strings.TrimSpace(v)
			}

			var macrosFloat []float64
			for _, v := range macros {
				number, err := strconv.ParseFloat(v, 64)
				if err != nil {
					log.Fatal("Error to conver string to float64:", err)
				} else {
					macrosFloat = append(macrosFloat, number)
				}

			}
			if len(macrosFloat) == 4 {
				p.addProduct(parts[1], macrosFloat[0], macrosFloat[1], macrosFloat[2], macrosFloat[3])
			} else {
				fmt.Println("Incorrect data. Don't point all the value of Macros")
			}
			return err
		}
	}

	return errors.New("Product isn't find")

}

func mealPlanning() {

}
