package stock

import (
	"fmt"
	"time"

	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg-charts/colors"
	"github.com/iostrovok/svg-charts/plast"
	"github.com/iostrovok/svg/style"
)

func Volume(g *plast.Plast, x time.Time, cWidth, volume int) error {

	converter := g.Converter()

	x1, err := converter.GetTimeX(x)
	if err != nil {
		return err
	}

	volumeV, err := converter.GetY(float64(volume))
	if err != nil {
		return err
	}

	text := fmt.Sprintf("%s\nvolume: %d", x.Format("2006-01-02 15:04:05"), volume)

	color := colors.GREEN
	if volumeV > 5000.0 {
		color = colors.RED
	}

	title := svg.Title(text)
	st2 := style.Style().StrokeWidth(0.5).Stroke(colors.BLACK).Fill(color)
	resc := svg.Rect(x1, g.GetPoint(volumeV), float64(cWidth), float64(volume), st2).Append(title)
	g.G.Append(resc)

	return nil
}
