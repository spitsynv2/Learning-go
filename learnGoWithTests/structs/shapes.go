package structs

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Heigth float64
}

func (r Rectangle) Area() float64 {
	return r.Heigth * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
