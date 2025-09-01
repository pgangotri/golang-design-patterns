package creational

import (
	"sync"
	"fmt"
)
// Singleton is a thread-safe singleton implementation in Go.
type Singleton struct {
	instance *Singleton
	once     sync.Once
}

// GetInstance returns the singleton instance of Singleton.
func (s *Singleton) GetInstance() *Singleton {
	s.once.Do(func() {
		s.instance = &Singleton{}
	})
	return s.instance
}

// Example usage of the Singleton pattern.
func ExampleSingleton() {
	singleton := &Singleton{}
	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()

	if instance1 == instance2 {
		fmt.Println("Both instances are the same.")
	} else {
		fmt.Println("Instances are different.")
	}
}