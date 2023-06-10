package Decorator

import "fmt"

func MainDecorator() {
	fmt.Println("Hello Decorator")

	pizza := &VeggieMania{}

	//Add Cheese
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of a Veggie Mania with Cheese and Tomato is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
