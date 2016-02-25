package grid

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func TestGrid(t *testing.T) {
	TestingT(t)
}

type GridTestsSuite struct{}

var _ = Suite(&GridTestsSuite{})

func (s GridTestsSuite) Test_ints(c *C) {
	// c.Skip("Not now")

	testInts := map[int]int{
		1:    2,
		2:    5,
		7:    10,
		55:   100,
		112:  200,
		554:  1000,
		1223: 2000,
	}

	for k, v := range testInts {
		p1 := findBestGridIntPoint(k)
		c.Assert(p1, Equals, v)
	}
}

func (s GridTestsSuite) Test_times(c *C) {
	// c.Skip("Not now")

	testTimes := map[time.Duration]time.Duration{
		1 * time.Second:    1 * time.Minute,
		15 * time.Second:   1 * time.Minute,
		75 * time.Second:   2 * time.Minute,
		1 * time.Minute:    1 * time.Minute,
		3 * time.Minute:    5 * time.Minute,
		7 * time.Minute:    10 * time.Minute,
		11 * time.Minute:   15 * time.Minute,
		17 * time.Minute:   30 * time.Minute,
		39 * time.Minute:   1 * time.Hour,
		7 * time.Hour:      12 * time.Hour,
		1 * 24 * time.Hour: 24 * time.Hour,
		36 * time.Hour:     48 * time.Hour,
		5 * 24 * time.Hour: 7 * 24 * time.Hour,

		// 10 * 24 * time.Hour:    0 * time.Hour,
		// 2 * 7 * 24 * time.Hour: 0 * time.Hour,
		// 3 * 7 * 24 * time.Hour: 0 * time.Hour,
		// 4 * 7 * 24 * time.Hour: 0 * time.Hour,
		// 365 * 24 * time.Hour:   0 * time.Hour,
	}

	for k, v := range testTimes {
		p := findBestGridTimePoint(k)
		c.Assert(p, Equals, v)
	}
}

func (s GridTestsSuite) Test_lines_int(c *C) {
	// c.Skip("Not now")

	testLine := GetLineGridPoints(-10, 90, 300, 5)

	c.Assert(len(testLine), Equals, 4)
	c.Assert(testLine[0].PosDraw, Equals, 270.0)
	c.Assert(testLine[0].PosReal, Equals, 0.0)

	c.Assert(testLine[1].PosDraw, Equals, 195.0)
	c.Assert(testLine[1].PosReal, Equals, 25.0)

	c.Assert(testLine[2].PosDraw, Equals, 120.0)
	c.Assert(testLine[2].PosReal, Equals, 50.0)

	c.Assert(testLine[3].PosDraw, Equals, 45.0)
	c.Assert(testLine[3].PosReal, Equals, 75.0)

}

func (s GridTestsSuite) Test_lines_time(c *C) {
	// c.Skip("Not now")

	t0, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 15:00:00")
	c.Assert(err, IsNil)
	t1, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 17:30:00")
	c.Assert(err, IsNil)

	testLine := GetTimeGridPoints(t0, t1, 300, 5)

	c.Assert(len(testLine), Equals, 4)
	c.Assert(testLine[0].PosDraw, Equals, 60.0)
	c.Assert(testLine[0].PosTime.Format("2006-01-02 15:04:05"), Equals, "2015-02-01 15:30:00")

	c.Assert(testLine[1].PosDraw, Equals, 120.0)
	c.Assert(testLine[1].PosTime.Format("2006-01-02 15:04:05"), Equals, "2015-02-01 16:00:00")

	c.Assert(testLine[2].PosDraw, Equals, 180.0)
	c.Assert(testLine[2].PosTime.Format("2006-01-02 15:04:05"), Equals, "2015-02-01 16:30:00")

	c.Assert(testLine[3].PosDraw, Equals, 240.0)
	c.Assert(testLine[3].PosTime.Format("2006-01-02 15:04:05"), Equals, "2015-02-01 17:00:00")

}
