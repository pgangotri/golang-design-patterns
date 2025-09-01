package creational

import "fmt"

/**
The Factory Pattern is a creational design pattern.
Its job: to create objects without exposing the creation logic to the client.
Instead of using new directly everywhere in your code, you delegate the creation of objects to a factory.
Think of it as an object generator that decides which specific object (class/instance) to give you, based on input or conditions.

Why use Factory Pattern?
1. Encapsulation of object creation logic → Clients don’t need to know how objects are created.
2. Flexibility → Easier to swap out implementations (e.g., switch from MySQL to MongoDB connector).
3. Loose coupling → Client depends on abstraction, not concrete classes.

Types of Factory Patterns

1. Simple Factory
	A single method decides which class to instantiate.
	Example: A ShapeFactory returns Circle or Square depending on input.
2. Factory Method
	Defines an interface for creating an object, but lets subclasses decide the exact type.
	Example: DocumentFactory → WordFactory, PDFfactory each know how to make their own type.
3. Abstract Factory
	Factory of factories → used when you need families of related objects.
	Example: UIFactory → WindowsUIFactory and MacUIFactory create consistent sets of UI elements (button, textbox, etc.).
*/

type Shape interface {
	Area() float64
	Draw() string
}

// Circle is a concrete implementation of Shape.
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle.
func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Draw returns a string representation of the circle.
func (c *Circle) Draw() string {
	return "Drawing a Circle with radius: " + fmt.Sprintf("%.2f", c.Radius)
}

// Square is a concrete implementation of Shape.
type Square struct {
	Side float64
}

// Area calculates the area of the square.
func (s *Square) Area() float64 {
	return s.Side * s.Side
}

// Draw returns a string representation of the square.
func (s *Square) Draw() string {
	return "Drawing a Square with side: " + fmt.Sprintf("%.2f", s.Side)
}

// ShapeFactory is a simple factory for creating shapes.
type ShapeFactory struct{}

// CreateShape creates a shape based on the type provided.
func (f *ShapeFactory) CreateShape(shapeType string, dimension float64) Shape {
	switch shapeType {
	case "circle":
		return &Circle{Radius: dimension}
	case "square":
		return &Square{Side: dimension}
	default:
		return nil
	}
}

// Example usage of the Factory Pattern.
func ExampleFactory() {	
	factory := &ShapeFactory{}

	circle := factory.CreateShape("circle", 5)
	fmt.Println(circle.Draw())
	fmt.Println("Area of Circle:", circle.Area())

	square := factory.CreateShape("square", 4)
	fmt.Println(square.Draw())
	fmt.Println("Area of Square:", square.Area())
}
