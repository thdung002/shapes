package entities

type Rectangle struct {
	Length float64 `json:"length"`
	Width  float64 `json:"with"`
}

func (s Rectangle) Area() float64 {
	return s.Length * s.Width
}

func (s Rectangle) Perimeter() float64 {
	return 2 * (s.Length + s.Width)
}
