package stock

import (
	"fmt"
	"time"

	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg-charts/colors"
	"github.com/iostrovok/svg-charts/plast"
	"github.com/iostrovok/svg/style"
)

type OneVolume struct {
	T     time.Time
	Text  string
	Y     int
	Width int
	St    *style.STYLE
	Debug bool
}

func Volume(g *plast.Plast, vol OneVolume) error {

	converter := g.Converter()

	x1, err := converter.GetTimeX(vol.T)
	if err != nil {
		return err
	}

	volumeV, err := converter.GetY(float64(vol.Y))
	if err != nil {
		return err
	}

	width := float64(vol.Width)
	if vol.Width == 0 {
		width = 5
	}

	text := vol.Text
	if vol.Text == "" {
		text = fmt.Sprintf("%s\nvolume: %d", vol.T.Format("2006-01-02 15:04:05"), vol.Y)
	}

	var st style.STYLE
	if vol.St == nil {
		st = style.Style().StrokeWidth(0.5).Stroke("black").Fill(colors.GREEN)
	} else {
		st = *vol.St
	}

	title := svg.Title(text)
	resc := svg.Rect(x1, g.GetPoint(volumeV), width, float64(volumeV), st).Append(title)
	g.G.Append(resc)

	return nil
}
