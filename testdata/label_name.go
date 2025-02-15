package testdata

func _() {
téstLabel: // want `identifier "téstLabel" contain non-ASCII character: U\+00E9 'é'`
	return
	goto téstLabel
}
