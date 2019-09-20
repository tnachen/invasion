package invasion

import "fmt"

const alienNameLength = 8

// Alien represents an alien
type Alien struct {
	Name  string
	City  string
	Turns int
}

func (a Alien) String() string {
	return fmt.Sprintf("%s city=%s turns=%d", a.Name, a.City, a.Turns)
}
