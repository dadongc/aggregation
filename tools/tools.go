package tools

import (
	"math"
)

func Cosine(a []float64, b []float64) (cosine float64) {
	count := 0
	length_a := len(a)
	length_b := len(b)
	if length_a > length_b {
		count = length_a
	} else {
		count = length_b
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= length_a {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= length_b {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2))
}

func Euclidean(a []float64, b []float64) float64 {
	count := 0
	length_a := len(a)
	length_b := len(b)
	if length_a > length_b {
		count = length_a
	} else {
		count = length_b
	}
	sumA := 0.0
	for k := 0; k < count; k++ {
		if k >= length_a {
			sumA += math.Pow(b[k], 2)
			continue
		}
		if k >= length_b {
			sumA += math.Pow(a[k], 2)
			continue
		}
		sumA += math.Pow(a[k]-b[k], 2)
	}
	sumA = math.Sqrt(sumA)
	return sumA
}
