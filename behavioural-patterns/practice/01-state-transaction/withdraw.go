package main

type Withdraw struct {
	User     *User
	Amount   float64
	TID      string
	State    WithdrawState
	Submit   WithdrawState
	Prepared WithdrawState
	Failed   WithdrawState
	Success  WithdrawState
	SendTo   string
}

type User struct {
	ID      int
	Name    string
	Balance float64
}

func newWithdraw(user *User, amount float64, sendTo string) *Withdraw {
	withdraw := &Withdraw{
		User:   user,
		Amount: amount,
		SendTo: sendTo,
		TID:    "TID001",
	}
	submitState := &SubmitState{
		w: withdraw,
	}

	preparedState := &PreparedState{
		w: withdraw,
	}

	failedState := &FailedState{
		w: withdraw,
	}

	successState := &SuccessState{
		w: withdraw,
	}

	withdraw.setState(submitState)
	withdraw.Prepared = preparedState
	withdraw.Success = successState
	withdraw.Failed = failedState
	return withdraw
}

func (w *Withdraw) setState(s WithdrawState) {
	w.State = s
}

func (w *Withdraw) pending() error {
	return w.State.submit()
}

func (w *Withdraw) prepared() error {
	return w.State.prepared()
}

func (w *Withdraw) fail() error {
	return w.State.failed()
}

func (w *Withdraw) success() error {
	return w.State.success()
}
