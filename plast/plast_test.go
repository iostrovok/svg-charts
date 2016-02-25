package plast

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func TestPlast(t *testing.T) {
	TestingT(t)
}

type PlastTestsSuite struct{}

var _ = Suite(&PlastTestsSuite{})

func _plast_time_test(c *C) (time.Time, time.Time) {
	t0, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 15:00:00")
	c.Assert(err, IsNil)
	t1, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 16:00:00")
	c.Assert(err, IsNil)
	return t0, t1
}

func (s PlastTestsSuite) Test_new(c *C) {
	//c.Skip("Not now")

	timeLeft, timeRight := _plast_time_test(c)
	realYBottom := 1000.0
	realYTop := 900.0
	width := 800.0
	hight := 400.0
	c1 := New(realYBottom, realYTop, width, hight, timeLeft, timeRight)
	c.Assert(c1, NotNil)
}
