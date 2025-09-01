package creational

import (
	"fmt"
	"sync"
)

// Prototype is a thread-safe prototype implementation in Go.
type Prototype struct {
	instance *Prototype
	once     sync.Once
}

// GetInstance returns the prototype instance of Prototype.
func (p *Prototype) GetInstance() *Prototype {
	p.once.Do(func() {
		p.instance = &Prototype{}
	})
	return p.instance
}

// Clone creates a copy of the prototype instance.
func (p *Prototype) Clone() *Prototype {
	return &Prototype{
		instance: p.instance,
	}
}

// Example usage of the Prototype pattern.
func ExamplePrototype() {
	prototype := &Prototype{}
	instance1 := prototype.GetInstance()
	instance2 := prototype.Clone()

	if instance1 == instance2 {
		fmt.Println("Both instances are the same.")
	} else {
		fmt.Println("Instances are different.")
	}
}
