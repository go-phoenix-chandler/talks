package main

import (
	"fmt"
	"sync"
)

type contact interface {
	GetName() string
	SetName(string)
}

type BasicContactFactory struct{}

func (bcf *BasicContactFactory) MakeContact() contact {
	return &basicContact{}
}

// START OMIT
type basicContact struct { //implements sync.Locker, contact
	sync.Mutex
	name string
}

func (c *basicContact) GetName() string {
	return c.name
}

func (c *basicContact) SetName(name string) {
	c.Lock()
	c.name = name
	c.Unlock()
}

func main() {
	factory := BasicContactFactory{}
	contact := factory.MakeContact()
	contact.SetName("bob")
	fmt.Println(contact.GetName())
}

//END OMIT
