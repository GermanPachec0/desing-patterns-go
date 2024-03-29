package observer_pattern

type Observer interface {
	update(string)
	getID() string
}
