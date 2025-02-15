package testdata

type _ struct {
	Tést     int // want `identifier "Tést" contain non-ASCII character: U\+00E9 'é'`
	_, Tést2 int // want `identifier "Tést2" contain non-ASCII character: U\+00E9 'é'`
}
