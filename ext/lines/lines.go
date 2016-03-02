package lines

import (
	"time"

	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg/style"

	"github.com/iostrovok/svg-charts/plast"
)

func LineTime(g *plast.Plast, x1 time.Time, y1 int, x2 time.Time, y2 int, sts ...style.STYLE) error {

	st := _style(sts...)

	converter := g.Converter()

	outX1, outY1, err := converter.GetByTime(x1, float64(y1))
	if err != nil {
		return err
	}

	outX2, outY2, err := converter.GetByTime(x2, float64(y2))
	if err != nil {
		return err
	}

	highLine := svg.Line(outX1, g.GetPoint(outY1), outX2, g.GetPoint(outY2), st) //.Append(title)
	g.G.Append(highLine)

	return nil
}

func Hor(g *plast.Plast, y1 int, sts ...style.STYLE) error {

	st := _style(sts...)

	outY, err := g.Converter().GetY(float64(y1))
	if err != nil {
		return err
	}

	line := svg.Path(st).M(0.0, g.GetPoint(outY)).H(g.W())
	g.G.Append(line)

	return nil
}

func VerByTime(g *plast.Plast, x time.Time, sts ...style.STYLE) error {

	st := _style(sts...)

	outX, err := g.Converter().GetTimeX(x)
	if err != nil {
		return err
	}

	line := svg.Path(st).M(outX, 0.0).V(g.H())
	g.G.Append(line)

	return nil
}
