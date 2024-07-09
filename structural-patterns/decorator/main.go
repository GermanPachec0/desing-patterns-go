package main

import "fmt"

func main() {
	pizza := &VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//add tomatto topping

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and ccheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())

}
