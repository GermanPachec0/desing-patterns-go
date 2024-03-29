package main

import "fmt"

type SuccessState struct {
	w *Withdraw
}

func (s *SuccessState) submit() error {

	return fmt.Errorf("item has not pass to submit again")
}

func (s *SuccessState) prepared() error {
	return fmt.Errorf("item has not pass to prepared again")
}

func (s *SuccessState) failed() error {
	return fmt.Errorf("item cannot be pass to fail")
}

func (s *SuccessState) success() error {
	s.w.User.Balance = s.w.User.Balance - s.w.Amount
	fmt.Printf("Success withdraw new user balance: %f \n", s.w.User.Balance)
	return nil
}
