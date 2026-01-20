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
		Total.calories += float64(v.gram) / 100 * v.calories
		Total.proteins += float64(v.gram) / 100 * v.proteins
		Total.fats += float64(v.gram) / 100 * v.fats
		Total.carbs += float64(v.gram) / 100 * v.carbs
	}

	fmt.Println()
	fmt.Println("----------Food intake----------")
	GetTotal(&Total)
	fmt.Println("-------------------------------")
	// for _, v := range listProducts {
	// 	GetProduct(&v)
	// }

}
