package entities

import (
	"math"
)

type Triangle struct {
	FirstSide  float64 `json:"first_side"`
	SecondSide float64 `json:"second_side"`
	ThirdSide  float64 `json:"third_side"`
}

func (t Triangle) Area() float64 {
	s := (t.Perimeter()) / 2
	return math.Sqrt(s * (s - t.FirstSide) * (s - t.SecondSide) * (s - t.ThirdSide))
}

func (t Triangle) Perimeter() float64 {
	return t.FirstSide + t.SecondSide + t.ThirdSide
}
