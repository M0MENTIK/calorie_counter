package main

import (
	"fmt"
	"os"
)

var fileName = "base.txt"

func main() {

	//showListProducts()
	listProducts := []Product{}
	for {
		fmt.Println("")
		var p Product
		var productNumber int
		fmt.Print("Enter the number of the product you want to find: ")
		fmt.Fscan(os.Stdin, &productNumber)
		err := findProductByNumber(&p, productNumber)
		if err == nil {
			listProducts = append(listProducts, p)
			fmt.Print("\nYou input:")
			GetProduct(&p)
		} else {
			fmt.Println(err)
		}

		var repeat string
		for {

			fmt.Print("\nAdding another product?(yes/no): ")
			fmt.Fscan(os.Stdin, &repeat)
			if repeat == "yes" || repeat == "no" {
				break
			} else {
				fmt.Println("Input: yes or no")
			}
		}
		if repeat == "yes" {
			continue
		} else if repeat == "no" {
			break
		}

	}

	var Total Product
	for _, v := range listProducts {
		Total.calories += v.calories
		Total.proteins += v.proteins
		Total.fats += v.fats
		Total.carbs += v.carbs
	}

	fmt.Println()
	fmt.Print("----------Food intake----------")
	GetProduct(&Total)
	fmt.Println("-------------------------------")
	// for _, v := range listProducts {
	// 	GetProduct(&v)
	// }

}
