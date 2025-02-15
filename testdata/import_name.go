package testdata

import téstImport "fmt" // want `identifier "téstImport" contain non-ASCII character: U\+00E9 'é'`

var _ téstImport.Stringer
