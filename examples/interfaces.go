package main

type Useable interface {
	Create()
}

type Item int8

// Item implements the Useable interface
func (i *Item) Create() {}

func main() {
	// Force checking that our item actually implements Useable
	var _ Useable = new(Item)
}

// EMPTYIFSTART OMIT
func ComplexGenericStuff(i interface{}) string {
	switch str := i.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}

	panic("Only strings or string like objects supported!")
}

type Stringer interface {
	String() string
}

// EMPTYIFSTOP OMIT
