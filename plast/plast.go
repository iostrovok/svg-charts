package plast

import (
	"time"

	"github.com/iostrovok/svg-charts/colors"
	"github.com/iostrovok/svg-charts/converter"
	"github.com/iostrovok/svg-charts/grid"

	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

type Point struct {
	Y int
	X time.Time
}

type Plast struct {
	G *svg.GROUP

	width, hight float64

	realYBottom, realYTop float64

	realTimeLeft, realTimeRight time.Time

	_gridListX []grid.PointTime
	_gridListY []grid.PointNum

	converter *converter.Converter
}

func New(realYBottom, realYTop, width, hight float64, realTimeLeft, realTimeRight time.Time) *Plast {

	g := &Plast{
		G: svg.Group(),

		realYBottom:   realYBottom,
		realYTop:      realYTop,
		realTimeLeft:  realTimeLeft,
		realTimeRight: realTimeRight,
		width:         width,
		hight:         hight,
	}

	g.converter = converter.New(
		realTimeLeft, realTimeRight,
		realYBottom, realYTop,
		width, hight,
	)

	st := style.Style().Fill(colors.LIGHT_GRAY)
	g.G.Style(st)

	return g
}

func (g *Plast) Converter() *converter.Converter {
	return g.converter
}

func (g *Plast) Move(x, y float64) {
	tr := transform.Transform().Translate(x, y)
	g.G.Transform(tr)
}

func (g *Plast) Init() {
	g.Border()
}

func (g *Plast) W() float64 {
	return g.width
}

func (g *Plast) H() float64 {
	return g.hight
}

func (g *Plast) LeftTime(t time.Time) {
	g.realTimeLeft = t
}

func (g *Plast) RightTime(t time.Time) {
	g.realTimeRight = t
}

func (g *Plast) Border() {
	st2 := style.Style().StrokeWidth(2).Stroke(colors.BLACK)
	rec2 := svg.Rect(1, 1, float64(g.width), float64(g.hight), st2)
	g.G.Append(rec2)
}

func (g *Plast) GridListX() []grid.PointTime {
	return g._gridListX
}
func (g *Plast) GridListY() []grid.PointNum {
	return g._gridListY
}

func (g *Plast) Grid(countPerX, countPerY int) {

	g._gridListX = grid.GetTimeGridPoints(g.realTimeLeft, g.realTimeRight, int(g.width), countPerX)
	g._gridListY = grid.GetLineGridPoints(int(g.realYBottom), int(g.realYTop), int(g.hight), countPerY)

	st := style.Style().StrokeWidth(1).Stroke(colors.GRAY)
	for _, v := range g._gridListY {
		line := svg.Line(2, v.PosDraw, g.width, v.PosDraw, st)
		g.G.Append(line)
	}

	for _, v := range g._gridListX {
		line := svg.Line(v.PosDraw, 2, v.PosDraw, g.hight, st)
		g.G.Append(line)
	}
}

func (g *Plast) GetPoint(y float64) float64 {
	return g.hight - y
}
