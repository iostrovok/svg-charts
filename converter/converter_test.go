package converter

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func TestConverter(t *testing.T) {
	TestingT(t)
}

type ConverterTestsSuite struct{}

var _ = Suite(&ConverterTestsSuite{})

func _converter_time_test(c *C) (time.Time, time.Time) {
	t0, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 15:00:00")
	c.Assert(err, IsNil)
	t1, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 16:00:00")
	c.Assert(err, IsNil)
	return t0, t1
}

func (s ConverterTestsSuite) Test_new(c *C) {
	//c.Skip("Not now")

	t0, t1 := _converter_time_test(c)
	c1 := New(t0, t1, 0.0, 1000.0, 400, 300)
	c.Assert(c1, NotNil)
}

func (s ConverterTestsSuite) Test_GetTimeX_1(c *C) {
	//c.Skip("Not now")

	t0, t1 := _converter_time_test(c)
	c1 := New(t0, t1, 0.0, 1000.0, 400, 300)

	tx1, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 12:30:00")
	c.Assert(err, IsNil)
	x1, err := c1.GetTimeX(tx1)
	c.Assert(err, NotNil)
	c.Assert(x1, Equals, 0.0)

	tx2, err := time.Parse("2006-01-02 15:04:05", "2016-02-01 12:30:00")
	c.Assert(err, IsNil)
	x2, err := c1.GetTimeX(tx2)
	c.Assert(err, NotNil)
	c.Assert(x2, Equals, 0.0)
}

func (s ConverterTestsSuite) Test_GetTimeX_2(c *C) {
	//c.Skip("Not now")

	t0, t1 := _converter_time_test(c)
	c1 := New(t0, t1, 0.0, 1000.0, 400, 300)

	tx, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 15:30:00")
	c.Assert(err, IsNil)
	x, err := c1.GetTimeX(tx)
	c.Assert(err, IsNil)
	c.Assert(x, Equals, 200.0)
}

func (s ConverterTestsSuite) Test_GetY_1(c *C) {
	//c.Skip("Not now")

	t0, t1 := _converter_time_test(c)
	c1 := New(t0, t1, 0.0, 1000.0, 400, 300)

	y1, err := c1.GetY(-1500)
	c.Assert(err, NotNil)
	c.Assert(y1, Equals, 0.0)

	y2, err := c1.GetY(1500)
	c.Assert(err, NotNil)
	c.Assert(y2, Equals, 0.0)
}

func (s ConverterTestsSuite) Test_GetY_2(c *C) {
	//c.Skip("Not now")

	t0, t1 := _converter_time_test(c)
	c1 := New(t0, t1, 0.0, 1000.0, 400, 300)

	y, err := c1.GetY(500)
	c.Assert(err, IsNil)
	c.Assert(y, Equals, 150.0)
}

func (s ConverterTestsSuite) Test_GetByTime_1(c *C) {
	//c.Skip("Not now")

	t0, t1 := _converter_time_test(c)
	c1 := New(t0, t1, 0.0, 1000.0, 400, 300)

	tx, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 15:30:00")
	c.Assert(err, IsNil)

	x, y, err := c1.GetByTime(tx, 500.0)
	c.Assert(err, IsNil)
	c.Assert(y, Equals, 150.0)
	c.Assert(x, Equals, 200.0)
}
