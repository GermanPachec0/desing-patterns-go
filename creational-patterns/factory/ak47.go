package main

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "Ak47",
			power: 2,
		},
	}
}
