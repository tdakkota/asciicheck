package testdata

type (
	_          struct{}
	TéstStruct struct{}   // want `identifier "TéstStruct" contain non-ASCII character: U\+00E9 'é'`
	TéstAlias  = struct{} // want `identifier "TéstAlias" contain non-ASCII character: U\+00E9 'é'`

	_[_ any]           struct{}
	_[TéstGeneric any] struct{} // want `identifier "TéstGeneric" contain non-ASCII character: U\+00E9 'é'`
)

func _[TéstGeneric any]() {} // want `identifier "TéstGeneric" contain non-ASCII character: U\+00E9 'é'`
