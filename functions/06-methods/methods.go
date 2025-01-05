package functions

type Person struct {
	Name string
}

func (p *Person) greet() string {
	return "Hello " + p.Name
}
