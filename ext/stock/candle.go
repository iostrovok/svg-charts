package stock

import (
	"fmt"
	"math"
	"time"

	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg/style"

	"github.com/iostrovok/svg-charts/colors"
	"github.com/iostrovok/svg-charts/converter"
	"github.com/iostrovok/svg-charts/plast"
)

type cand struct {
	x          float64
	text       string
	open, clos float64
	high, low  float64
	color      string
}

func Candle(g *plast.Plast, t time.Time, cWidth, open, clos, high, low int) error {

	converter := g.Converter()

	c, err := candlePrepare(t, open, clos, high, low, converter)
	if err != nil {
		return err
	}

	title := svg.Title(c.text)

	st := style.Style().StrokeWidth(0.5).Stroke("black")
	x2 := c.x + float64(cWidth)/2
	highLine := svg.Line(x2, g.GetPoint(c.high), x2, g.GetPoint(c.low), st).Append(title)
	g.G.Append(highLine)

	st2 := style.Style().StrokeWidth(0.5).Stroke("black").Fill(c.color)
	resc := svg.Rect(c.x, g.GetPoint(c.clos), float64(cWidth), math.Abs(c.clos-c.open), st2).Append(title)
	g.G.Append(resc)

	return nil
}

func candlePrepare(t time.Time, open, clos, high, low int, converter *converter.Converter) (*cand, error) {

	c := &cand{}
	var err error

	c.x, err = converter.GetTimeX(t)
	if err != nil {
		return nil, err
	}

	c.clos, err = converter.GetY(float64(clos))
	if err != nil {
		return nil, err
	}
	c.open, err = converter.GetY(float64(open))
	if err != nil {
		return nil, err
	}

	c.high, err = converter.GetY(float64(high))
	if err != nil {
		return nil, err
	}

	c.low, err = converter.GetY(float64(low))
	if err != nil {
		return nil, err
	}

	c.text = fmt.Sprintf("%s\nOpen: %d\nClose: %d\nHigh: %d\nLow: %d", t.Format("2006-01-02 15:04:05"), open, clos, high, low)

	c.color = colors.GREEN
	if open > clos {
		c.color = colors.RED
		c.clos, c.open = c.open, c.clos
	}

	return c, nil
}
