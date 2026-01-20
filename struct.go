package main

import "fmt"

type Product struct {
	name     string
	calories float64
	proteins float64
	fats     float64
	carbs    float64
	gram     int
}

func (p *Product) addProduct(name string, calories, proteins, fats, carbs float64, gram int) {
	if name != "" {
		p.name = name
	} else {
		fmt.Print("Incorrect name")
		return
	}
	if calories >= 0 && proteins >= 0 && fats >= 0 && carbs >= 0 && gram >= 0 {
		p.calories = calories
		p.proteins = proteins
		p.fats = fats
		p.carbs = carbs
		p.gram = gram
	} else {
		fmt.Print("Incorrect data. (Calories, proteins, fars or carbs >= 0)")
		return
	}
}

func GetProduct(p *Product) {
	fmt.Println()
	fmt.Println("Product:", p.name)
	fmt.Println("Kcal:", p.calories)
	fmt.Println("Proteins:", p.proteins)
	fmt.Println("Fats:", p.fats)
	fmt.Println("Carbs:", p.carbs)
	fmt.Println("Grams:", p.gram)
	fmt.Println()
}

func GetTotal(p *Product) {
	fmt.Printf("Kcal: %.2f\n", p.calories)
	fmt.Printf("Proteins: %.2f\n", p.proteins)
	fmt.Printf("Fats: %.2f\n", p.fats)
	fmt.Printf("Carbs: %.2f\n", p.carbs)
}
