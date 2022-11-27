package main

func main() {

	// Declare a new item
	shirtItem := newItem("Adidas shirt")

	// Declare observers (users)
	observerFirst := &User{id: "aladar@gmail.com"}
	observerSec := &User{id: "bela@gmail.com"}
	observerThird := &User{id: "eva@gmail.com"}
	observerFourth := &User{id: "sofi@gmail.com"}

	// Register users for notifications
	shirtItem.register(observerFirst)
	shirtItem.register(observerSec)
	shirtItem.register(observerThird)
	shirtItem.register(observerFourth)

	// Update item availability status
	shirtItem.updateAvailability()
}
