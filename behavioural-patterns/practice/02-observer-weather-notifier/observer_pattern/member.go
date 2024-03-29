package observer_pattern

import "fmt"

type Member struct {
	Id string
}

func (c *Member) update(weather string) {
	fmt.Printf("Notifying to Member %s for item %s \n", c.Id, weather)
}

func (c *Member) getID() string {
	return c.Id
}
