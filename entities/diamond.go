package entities

type Diamond struct {
	Length float64 `json:"length"`
	Height float64 `json:"height"`
	Side   float64 `json:"side"`
}

func (d Diamond) Area() float64 {
	return (d.Length * d.Height) / 2
}

func (d Diamond) Perimeter() float64 {
	return 4 * d.Side
}
