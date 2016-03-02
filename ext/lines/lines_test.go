package lines

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"

	"github.com/iostrovok/svg-charts/plast"
	"github.com/iostrovok/svg-charts/points"
)

func TestLines(t *testing.T) {
	TestingT(t)
}

type LinesTestsSuite struct{}

var _ = Suite(&LinesTestsSuite{})

func _stock_time_test(c *C, str string) time.Time {
	t0, err := time.Parse("2006-01-02 15:04:05", str)
	c.Assert(err, IsNil)
	return t0
}

func (s LinesTestsSuite) Test_LineTime(c *C) {
	// c.Skip("Not now")

	timeLeft := _stock_time_test(c, "2015-02-01 15:00:00")
	timeRight := _stock_time_test(c, "2015-02-01 16:00:00")
	realYBottom := 10.0
	realYTop := 900.0
	width := 800.0
	hight := 400.0

	p := plast.New(realYBottom, realYTop, width, hight, timeLeft, timeRight)
	c.Assert(p, NotNil)

	x1 := _stock_time_test(c, "2015-02-01 16:00:00")
	x2 := _stock_time_test(c, "2015-02-01 16:00:00")

	err := LineTime(p, x1, 100, x2, 200)
	c.Assert(err, IsNil)
}

func (s LinesTestsSuite) Test_Smooth(c *C) {
	// c.Skip("Not now")

	timeLeft := _stock_time_test(c, "2015-02-01 15:00:00")
	timeRight := _stock_time_test(c, "2015-02-01 16:00:00")
	realYBottom := 10.0
	realYTop := 900.0
	width := 800.0
	hight := 400.0

	p := plast.New(realYBottom, realYTop, width, hight, timeLeft, timeRight)
	c.Assert(p, NotNil)

	list := []points.PointTime{}
	err := SmoothByTime(p, list)
	c.Assert(err, IsNil)
}
