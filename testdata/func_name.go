package testdata

func TéstFunc() {} // want `identifier "TéstFunc" contain non-ASCII character: U\+00E9 'é'`

type _ interface {
	TéstFunc() // want `identifier "TéstFunc" contain non-ASCII character: U\+00E9 'é'`
}
