package hist

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
	"time"

	"github.com/iostrovok/svg-charts/plast"
	"github.com/iostrovok/svg/style"
)

func TestHist(t *testing.T) {
	TestingT(t)
}

type HistTestsSuite struct{}

var _ = Suite(&HistTestsSuite{})

func _stock_time_test(c *C, str string) time.Time {
	t0, err := time.Parse("2006-01-02 15:04:05", str)
	c.Assert(err, IsNil)
	return t0
}

func (s HistTestsSuite) Test_Base(c *C) {
	// c.Skip("Not now")

	res := `<g style="fill:rgb(240,240,240)"><rect height="10" width="5" x="50" y="40"><title>2015-02-01 15:30:00
volume: 10</title></rect><rect height="10" width="5" x="58.33" y="50"><title>2015-02-01 15:35:00
volume: -10</title></rect></g>`

	timeLeft := _stock_time_test(c, "2015-02-01 15:00:00")
	timeRight := _stock_time_test(c, "2015-02-01 16:00:00")
	realYBottom := 0.0
	realYTop := 100.0
	width := 100.0
	hight := 100.0

	p := plast.New(realYBottom, realYTop, width, hight, timeLeft, timeRight)
	c.Assert(p, NotNil)

	t := _stock_time_test(c, "2015-02-01 15:30:00")
	cnd := OneVolume{
		T:     t,
		Value: 10,
		BaseY: 50,
		Debug: true,
		St:    style.Style().Ref(),
	}

	err := Base(p, cnd)
	c.Assert(err, IsNil)

	t1 := _stock_time_test(c, "2015-02-01 15:35:00")
	cnd = OneVolume{
		T:     t1,
		Value: -10,
		BaseY: 50,
		Debug: true,
		St:    style.Style().Ref(),
	}

	err = Base(p, cnd)
	c.Assert(err, IsNil)

	fmt.Printf("\n------\np.G.Source(): %s\n------\n", p.G.Source())

	c.Assert(p.G.Source(), Equals, res)
	// c.Assert(false, IsNil)

}
