package lines

import (
	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg/style"

	"github.com/iostrovok/svg-charts/colors"
	"github.com/iostrovok/svg-charts/plast"
	"github.com/iostrovok/svg-charts/points"
)

func SmoothByTime(g *plast.Plast, list []points.PointTime) error {

	st := style.Style().StrokeWidth(0.5).Stroke(colors.RED).Fill("none")

	converter := g.Converter()

	line := svg.Path(st)
	for i, p := range list {

		outX, outY, err := converter.GetByTime(p.X, float64(p.Y))
		if err != nil {
			return err
		}

		outX += p.DisX
		outY = g.GetPoint(outY + p.DisY)

		if i == 0 {
			line = line.M(outX, outY)
		} else {
			line = line.L(outX, outY)
		}
	}

	g.G.Append(line)

	return nil
}
