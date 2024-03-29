package main

import (
	"fmt"
	"log"
)

func getOne(id int) (*User, error) {
	users := FakeUserData()

	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, fmt.Errorf("User Not Found")
}
func FakeUserData() []*User {
	users := []*User{
		{ID: 1, Name: "Alice", Balance: 1000.0},
		{ID: 2, Name: "Bob", Balance: 500.0},
		{ID: 3, Name: "Charlie", Balance: 2000.0},
		{ID: 4, Name: "David", Balance: 150.0},
		// Add more fake users as needed
	}

	return users
}

func main() {
	user, err := getOne(2)
	if err != nil {
		fmt.Println("user not found")
	}
	withdraw := newWithdraw(user, 100.00, "Pepito")

	err = withdraw.pending()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = withdraw.prepared()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = withdraw.success()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
