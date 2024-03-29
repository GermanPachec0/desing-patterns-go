package main

import "fmt"

type PreparedState struct {
	w *Withdraw
}

func (s *PreparedState) submit() error {
	s.w.setState(s.w.Submit)
	fmt.Println("The withdraw will be process again")
	return nil
}

func (s *PreparedState) prepared() error {
	if s.w.Amount < 0 {
		s.w.setState(s.w.Failed)
		return fmt.Errorf("withdraw cannot be below 0")
	}
	if s.w.User.Balance <= s.w.Amount {
		s.w.setState(s.w.Failed)
		return fmt.Errorf("User Insufficent balance")

	}
	fmt.Println("Item Prepared Correctly passed to success")
	s.w.setState(s.w.Success)
	return nil
}

func (s *PreparedState) failed() error {
	s.w.setState(s.w.Failed)
	return fmt.Errorf("item failed succefuly from prepared")
}

func (s *PreparedState) success() error {
	return fmt.Errorf("item cannot be success yet")
}
