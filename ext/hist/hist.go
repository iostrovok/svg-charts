package hist

import (
	"fmt"
	"math"
	"time"

	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg-charts/colors"
	"github.com/iostrovok/svg-charts/plast"
	"github.com/iostrovok/svg/style"
)

type OneVolume struct {
	T     time.Time
	Text  string
	Value int
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
		fmt.Printf("vol.T: %s, vol.BaseY: %d, vol.volue: %d, vol.Width: %d\n", vol.T.Format("2006-01-02 15:04:05"), vol.BaseY, vol.Value, vol.Width)
	}

	converter := g.Converter()

	y := vol.BaseY
	if vol.Value > 0 {
		y = vol.BaseY + vol.Value
	}

	outX1, outY1, err := converter.GetByTime(vol.T, float64(y))
	if err != nil {
		return err
	}

	width := float64(vol.Width)
	if vol.Width == 0 {
		width = 5
	}

	text := vol.Text
	if vol.Text == "" {
		text = fmt.Sprintf("%s\nvolume: %d", vol.T.Format("2006-01-02 15:04:05"), vol.Value)
	}

	var st style.STYLE
	if vol.St == nil {
		st = style.Style().StrokeWidth(0.5).Stroke("black").Fill(colors.GREEN)
	} else {
		st = *vol.St
	}

	if vol.Debug {
		fmt.Printf("outX1: %f, g.GetPoint(outY1): %f, outY1: %f, width: %f\n", outX1, g.GetPoint(outY1), outY1, width)
	}

	value, err := converter.GetSizeY(math.Abs(float64(vol.Value)))
	if err != nil {
		return err
	}

	title := svg.Title(text)
	resc := svg.Rect(outX1, g.GetPoint(outY1), width, float64(value), st).Append(title)
	g.G.Append(resc)

	return nil
}
