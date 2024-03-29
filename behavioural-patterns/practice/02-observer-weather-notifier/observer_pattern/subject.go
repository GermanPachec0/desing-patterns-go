package observer_pattern

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll(weather string)
}
