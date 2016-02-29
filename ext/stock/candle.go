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

type OneCandle struct {
	T           time.Time
	Text        string
	Open, Close int
	High, Low   int
	Width       int
	StDown      *style.STYLE
	StUp        *style.STYLE
	StBorder    *style.STYLE
	Debug       bool
}

type cand struct {
	x          float64
	text       string
	open, clos float64
	high, low  float64
	color      string
	width      float64
	st         style.STYLE
	stBorder   style.STYLE
}

func Candle(g *plast.Plast, candle OneCandle) error {

	converter := g.Converter()

	c, err := candlePrepare(candle, converter)
	if err != nil {
		return err
	}

	title := svg.Title(c.text)

	x2 := c.x + c.width/2
	if candle.Debug {
		fmt.Printf("x2: %f, g.GetPoint(c.high): %f, x2: %f, g.GetPoint(c.low): %f\n", x2, g.GetPoint(c.high), x2, g.GetPoint(c.low))
	}
	highLine := svg.Line(x2, g.GetPoint(c.high), x2, g.GetPoint(c.low), c.stBorder)
	g.G.Append(highLine)

	if candle.Debug {
		fmt.Printf("c.x: %f, g.GetPoint(c.clos): %f, c.width: %f, math.Abs(c.clos-c.open): %f\n", c.x, g.GetPoint(c.clos), c.width, math.Abs(c.clos-c.open))
	}

	resc := svg.Rect(c.x, g.GetPoint(c.clos), c.width, math.Abs(c.clos-c.open), c.st).Append(title)
	g.G.Append(resc)

	return nil
}

func candlePrepare(candle OneCandle, converter *converter.Converter) (*cand, error) {

	c := &cand{}
	var err error

	c.x, err = converter.GetTimeX(candle.T)
	if err != nil {
		return nil, err
	}

	c.clos, err = converter.GetY(float64(candle.Close))
	if err != nil {
		return nil, err
	}
	c.open, err = converter.GetY(float64(candle.Open))
	if err != nil {
		return nil, err
	}

	c.high, err = converter.GetY(float64(candle.High))
	if err != nil {
		return nil, err
	}

	c.low, err = converter.GetY(float64(candle.Low))
	if err != nil {
		return nil, err
	}

	c.text = candle.Text
	if c.text == "" {
		c.text = fmt.Sprintf("%s\nOpen: %d\nClose: %d\nHigh: %d\nLow: %d", candle.T.Format("2006-01-02 15:04:05"), candle.Open, candle.Close, candle.High, candle.Low)
	}

	if candle.StBorder == nil {
		c.stBorder = style.Style().StrokeWidth(0.5).Stroke("black")
	} else {
		c.stBorder = *candle.StBorder
	}

	if candle.Width > 0 {
		c.width = float64(candle.Width)
	} else {
		c.width = 5.0
	}

	c.color = colors.GREEN
	if candle.Open > candle.Close {
		c.clos, c.open = c.open, c.clos
		if candle.StDown == nil {
			c.st = style.Style().StrokeWidth(0.5).Stroke("black").Fill(colors.RED)
		} else {
			c.st = *candle.StDown
		}
	} else {
		if candle.StUp == nil {
			c.st = style.Style().StrokeWidth(0.5).Stroke("black").Fill(colors.GREEN)
		} else {
			c.st = *candle.StUp
		}
	}

	return c, nil
}
