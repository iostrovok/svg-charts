package converter

import (
	"fmt"
	"time"
)

type Converter struct {
	koef            float64
	xDur            time.Duration
	startTimeX      time.Time
	finishTimeX     time.Time
	startY, finishY float64
	width, hight    float64
}

func New(
	startTimeX, finishTimeX time.Time,
	startY, finishY, width, hight float64) *Converter {

	c := &Converter{
		koef:        hight / (finishY - startY),
		xDur:        finishTimeX.Sub(startTimeX),
		startTimeX:  startTimeX,
		finishTimeX: finishTimeX,
		startY:      startY,
		finishY:     finishY,
		width:       width,
		hight:       hight,
	}

	return c
}

func (c *Converter) GetTimeX(x time.Time) (float64, error) {
	if x.Before(c.startTimeX) || x.After(c.finishTimeX) {
		return 0, fmt.Errorf("Bad time %s", x.Format("2006-01-02 15:04:05"))
	}
	outX := c.width * float64(x.Sub(c.startTimeX)) / float64(c.xDur)
	return outX, nil
}

func (c *Converter) GetY(y float64) (float64, error) {

	if y < c.startY || c.finishY < y {
		return 0, fmt.Errorf("Bad Y: %f < %f || %f < %f", y, c.startY, c.finishY, y)
	}

	outY := (y - c.startY) * c.koef
	return outY, nil
}

func (c *Converter) GetByTime(x time.Time, y float64) (float64, float64, error) {

	outX, err := c.GetTimeX(x)
	if err != nil {
		return 0, 0, err
	}

	outY, err := c.GetY(y)
	if err != nil {
		return 0, 0, err
	}

	return outX, outY, nil
}
