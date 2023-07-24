package dto

type ImageLabel struct {
	Image []byte
	Label byte
}

type ImageVector struct {
	Image []float64
	Label int
}

type CosineLabel struct {
	Cosine float64
	Label  int
}
