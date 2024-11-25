package testdata

func _[TеstGeneric any]() { // want `identifier "TеstGeneric" contain non-ASCII character: U\+0435 'е'`
}

type _[TеstGeneric any] struct { // want `identifier "TеstGeneric" contain non-ASCII character: U\+0435 'е'`
}
