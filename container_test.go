package svgcharts

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestContainer(t *testing.T) {
	TestingT(t)
}

type TestsSuite struct{}

var _ = Suite(&TestsSuite{})

func (s TestsSuite) Test_text(c *C) {
	//c.Skip("Not now")

	c1 := New(100, 5, 10)
	c.Assert(c1, NotNil)
}
