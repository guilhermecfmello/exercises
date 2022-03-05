package main

import "fmt"

func main() {
	fmt.Println("Hello World Animal!")

	cat := Cat{
		Mammal: Mammal{
			Name:  "Noir",
			Noise: "Meow",
		},
		ShowerTookAmount: 189,
	}

	dog := Dog{
		Mammal: Mammal{
			Name:  "Toby",
			Noise: "Auf",
		},
		LoveAmount: 999,
	}

	fmt.Println(cat)
	fmt.Println(dog)

	fmt.Println(cat.GetNoise())
	fmt.Println(dog.GetNoise())
}

func (mammal Mammal) GetNoise() string {
	FullNoise := "Noise comming..." + mammal.Noise
	return FullNoise
}

type Animal interface {
	GetNoise() string
}

type Mammal struct {
	Name  string
	Noise string
}

type Cat struct {
	Mammal
	ShowerTookAmount int
}

type Dog struct {
	Mammal
	LoveAmount int
}
