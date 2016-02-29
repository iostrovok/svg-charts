package stock

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"

	"github.com/iostrovok/svg-charts/plast"
	"github.com/iostrovok/svg/style"
)

func TestStock(t *testing.T) {
	TestingT(t)
}

type StockTestsSuite struct{}

var _ = Suite(&StockTestsSuite{})

func _stock_time_test(c *C, str string) time.Time {
	t0, err := time.Parse("2006-01-02 15:04:05", str)
	c.Assert(err, IsNil)
	return t0
}

func (s StockTestsSuite) Test_Candle(c *C) {
	// c.Skip("Not now")

	timeLeft := _stock_time_test(c, "2015-02-01 15:00:00")
	timeRight := _stock_time_test(c, "2015-02-01 16:00:00")
	realYBottom := 10.0
	realYTop := 900.0
	width := 800.0
	hight := 400.0

	p := plast.New(realYBottom, realYTop, width, hight, timeLeft, timeRight)
	c.Assert(p, NotNil)

	t := _stock_time_test(c, "2015-02-01 15:30:00")
	cnd := OneCandle{
		T:     t,
		Open:  300,
		Close: 400,
		High:  100,
		Low:   300,
		Width: 5,
	}

	err := Candle(p, cnd)
	c.Assert(err, IsNil)
}

func (s StockTestsSuite) Test_Candle_2(c *C) {
	// c.Skip("Not now")

	timeLeft := _stock_time_test(c, "2015-02-01 15:00:00")
	timeRight := _stock_time_test(c, "2015-02-01 16:00:00")
	realYBottom := 10.0
	realYTop := 900.0
	width := 800.0
	hight := 400.0

	p := plast.New(realYBottom, realYTop, width, hight, timeLeft, timeRight)
	c.Assert(p, NotNil)

	st := style.Style().StrokeWidth(0.5).Stroke("black").Fill("RED")
	t := _stock_time_test(c, "2015-02-01 15:30:00")
	cnd := OneCandle{
		T:        t,
		Text:     "My TEST",
		Open:     300,
		Close:    400,
		High:     100,
		Low:      300,
		StDown:   &st,
		StUp:     &st,
		StBorder: &st,
	}

	err := Candle(p, cnd)
	c.Assert(err, IsNil)
}
