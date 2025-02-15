package testdata

import "fmt"

const téstGlobalConst = 0 // want `identifier "téstGlobalConst" contain non-ASCII character: U\+00E9 'é'`

var téstGlobalVar int // want `identifier "téstGlobalVar" contain non-ASCII character: U\+00E9 'é'`

func _() {
	const téstConst = 0 // want `identifier "téstConst" contain non-ASCII character: U\+00E9 'é'`

	var téstVar int // want `identifier "téstVar" contain non-ASCII character: U\+00E9 'é'`
	téstVar = 0
	fmt.Println(téstVar)

	var ch chan int
	téstVar2, _ := <-ch // want `identifier "téstVar2" contain non-ASCII character: U\+00E9 'é'`
	fmt.Println(téstVar2)
}

func _(téstParam int)     {}         // want `identifier "téstParam" contain non-ASCII character: U\+00E9 'é'`
func _() (téstResult int) { return } // want `identifier "téstResult" contain non-ASCII character: U\+00E9 'é'`

type _ interface {
	m1(téstParam int)     // want `identifier "téstParam" contain non-ASCII character: U\+00E9 'é'`
	m2() (téstResult int) // want `identifier "téstResult" contain non-ASCII character: U\+00E9 'é'`
}

type Recv struct{}

func (téstRecv Recv) _() {} // want `identifier "téstRecv" contain non-ASCII character: U\+00E9 'é'`
