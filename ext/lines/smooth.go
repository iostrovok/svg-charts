package lines

import (
	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg/style"

	"github.com/iostrovok/svg-charts/colors"
	"github.com/iostrovok/svg-charts/plast"
)

func Smooth(g *plast.Plast, width int, list []plast.Point) error {

	st := style.Style().StrokeWidth(0.5).Stroke(colors.RED).Fill("none")

	converter := g.Converter()

	line := svg.Path(st)
	for i, p := range list {

		outX, outY, err := converter.GetByTime(p.X, float64(p.Y))
		if err != nil {
			return err
		}

		outX += float64(width) / 2
		outY = g.GetPoint(outY)

		if i == 0 {
			line = line.M(outX, outY)
		} else {
			line = line.L(outX, outY)
		}
	}

	g.G.Append(line)

	return nil
}
