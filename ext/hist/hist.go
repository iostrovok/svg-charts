package hist

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
	BaseY int
	St    *style.STYLE
	Debug bool
}

func Base(g *plast.Plast, vol OneVolume) (err error) {

	defer func() {
		if err != nil && vol.Debug {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	if vol.Debug {
		fmt.Printf("vol.T: %s, vol.BaseY: %d, vol.Y: %d, vol.Width: %d\n", vol.T.Format("2006-01-02 15:04:05"), vol.BaseY, vol.Y, vol.Width)
	}

	converter := g.Converter()

	outX1, outY1, err := converter.GetByTime(vol.T, float64(vol.BaseY+vol.Y))
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

	if vol.Debug {
		fmt.Printf("x1: %d, g.GetPoint(volumeV): %f, volumeV: %f, width: %f\n", x1, g.GetPoint(volumeV), volumeV, width)
	}

	title := svg.Title(text)
	resc := svg.Rect(x1, g.GetPoint(volumeV), width, float64(volumeV), st).Append(title)
	g.G.Append(resc)

	return nil
}
