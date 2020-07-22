package testing

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey" // import进来的package的所有方法在当前的名字空间，直接使用方法即可
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
