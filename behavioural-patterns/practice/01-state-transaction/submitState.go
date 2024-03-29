package main

import "fmt"

type SubmitState struct {
	w *Withdraw
}

func (s *SubmitState) submit() error {
	if s.w == nil {
		s.w.setState(s.w.Failed)
		return fmt.Errorf("Withdraw not found")
	}
	fmt.Println("withdraw correctly created passed to prepared")
	s.w.setState(s.w.Prepared)
	return nil
}

func (s *SubmitState) prepared() error {
	return fmt.Errorf("item has not be submitted yet")
}

func (s *SubmitState) failed() error {
	return fmt.Errorf("item cannot be failed yet ")
}

func (s *SubmitState) success() error {
	return fmt.Errorf("item Cannot be success yet")
}
