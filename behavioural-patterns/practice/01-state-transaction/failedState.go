package main

import "fmt"

type FailedState struct {
	w *Withdraw
}

func (s *FailedState) submit() error {
	return fmt.Errorf("item has been failed")
}

func (s *FailedState) prepared() error {
	return fmt.Errorf("item has been failed")

}

func (s *FailedState) failed() error {
	return fmt.Errorf("item has been failed")
}

func (s *FailedState) success() error {
	return fmt.Errorf("item has been failed")
}
