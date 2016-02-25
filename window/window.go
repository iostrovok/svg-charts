package window

import (
	"fmt"
	"time"

	"github.com/iostrovok/svg-charts/grid"
	"github.com/iostrovok/svg-charts/plast"
)

type Window struct {
	Plast *plast.Plast

	realYBottom float64
	realYTop    float64
	width       float64
	hight       float64

	startTop     float64
	percentHight float64

	realTimeLeft  time.Time
	realTimeRight time.Time
}

func New(realYBottom, realYTop, width, hight, percentHight float64, realTimeLeft, realTimeRight time.Time) *Window {
	return &Window{
		realYBottom:   realYBottom,
		realYTop:      realYTop,
		Plast:         nil,
		realTimeLeft:  realTimeLeft,
		realTimeRight: realTimeRight,
		width:         width,
		hight:         hight,
		percentHight:  percentHight,
	}

}

func (w *Window) PercentHight(t ...float64) float64 {

	if len(t) > 0 {
		w.percentHight = t[0]
	}

	return w.percentHight
}

func (w *Window) Hight(t ...float64) float64 {

	if len(t) > 0 {
		w.hight = t[0]
	}

	return w.hight
}

func (w *Window) Width(t ...float64) float64 {

	if len(t) > 0 {
		w.width = t[0]
	}

	return w.width
}

func (w *Window) StartTop(t ...float64) float64 {

	if len(t) > 0 {
		w.startTop = t[0]
	}

	return w.startTop
}

func (w *Window) GridListX() []grid.PointTime {
	return w.Plast.GridListX()
}

func (w *Window) GridListY() []grid.PointNum {
	return w.Plast.GridListY()
}

func (w *Window) LeftTime(t ...time.Time) time.Time {

	if len(t) > 0 {
		w.realTimeLeft = t[0]
	}

	return w.realTimeLeft
}

func (w *Window) RightTime(t ...time.Time) time.Time {

	if len(t) > 0 {
		w.realTimeRight = t[0]
	}

	return w.realTimeRight
}

func (w *Window) Finish(leftField float64) {

	w.Plast = plast.New(w.realYBottom, w.realYTop,
		w.width, w.hight,
		w.realTimeLeft, w.realTimeRight)

	w.Plast.Init()

	// w.Plast.Grid(60, int((w.realYTop-w.realYBottom)/2000))
	w.Plast.Grid(60, int(35.0*w.percentHight/100))

	fmt.Printf("Move -> float64(c.leftField): %f, w.startTop: %f\n", leftField, w.startTop)
	w.Plast.Move(leftField, w.startTop)
}
