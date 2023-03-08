package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (m1 *Animal) Eat() {
	fmt.Println("The food it eats is", m1.food)
}

func (m1 *Animal) Move() {
	fmt.Println("It ", m1.locomotion)
}

func (m1 *Animal) Speak() {
	fmt.Printf("It %s's\n", m1.noise)
}

func query() {

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	var animalMap map[string]Animal
	animalMap = make(map[string]Animal)

	animalMap["cow"] = cow
	animalMap["bird"] = bird
	animalMap["snake"] = snake

	for {
		var animal, query_type string
		fmt.Println("Enter the animal and the query type")
		fmt.Scanln(&animal, &query_type)
		input := animalMap[animal]
		switch query_type {
		case "eat":
			input.Eat()
		case "move":
			input.Move()
		case "speak":
			input.Speak()
		default:
			fmt.Println("Wrong Query, Please try again")
		}
	}

}
