package main

type Useable interface {
	Create()
}

type Item int8

// Item implements the Useable interface
func (i *Item) Create() {}

func main() {
	// Force checking that our iten actually implements Useable
	var _ Useable = new(Item)
}
