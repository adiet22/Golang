package Task03

type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
}

type Hitung interface {
	Hitung2d
	Hitung3d
}
