package main

type WithdrawState interface {
	//Validates if the withdraw have all the field setss
	submit() error
	// Validates if the user have the sufficient founds
	prepared() error
	// If the user have not founds
	failed() error
	// Success withdraw
	success() error
}
