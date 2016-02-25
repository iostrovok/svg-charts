package window

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func TestWindow(t *testing.T) {
	TestingT(t)
}

type TestsWindowSuite struct{}

var _ = Suite(&TestsWindowSuite{})

func _window_time_test(c *C) (time.Time, time.Time) {
	t0, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 15:00:00")
	c.Assert(err, IsNil)
	t1, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 16:00:00")
	c.Assert(err, IsNil)
	return t0, t1
}

func (s TestsWindowSuite) Test_text(c *C) {
	//c.Skip("Not now")

	realTimeLeft, realTimeRight := _window_time_test(c)

	realYBottom := 100.0
	realYTop := 100.0
	width := 100.0
	hight := 100.0
	percentHight := 100.0

	c1 := New(realYBottom, realYTop, width, hight, percentHight, realTimeLeft, realTimeRight)
	c.Assert(c1, NotNil)
}
