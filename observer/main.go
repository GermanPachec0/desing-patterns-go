package main

func main() {
	shirtItem := newItem("Nike Shirt")
	observerFisrt := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFisrt)
	shirtItem.register(observerSecond)

	mouseItem := newItem("Logitech Mouse")

	mouseItem.register(observerFisrt)
	mouseItem.register(observerSecond)
	mouseItem.deregister(observerFisrt)
	shirtItem.updateAvailability()
	mouseItem.updateAvailability()
}
