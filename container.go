package svgcharts

import (
	"fmt"
	// "strconv"
	// "strings"
	"time"

	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"

	"github.com/iostrovok/svg-charts/colors"
	"github.com/iostrovok/svg-charts/grid"
	"github.com/iostrovok/svg-charts/window"
	// "github.com/iostrovok/svg-charts/plast"
	// "github.com/iostrovok/svg-charts/plast"
	stock "github.com/iostrovok/svg-charts/ext/stock"
)

type container struct {
	G *svg.GROUP

	GridListX []grid.PointTime

	width, hight float64
	winGap       float64

	candleWidth int

	globalTimeLeft, globalrealTimeRight       time.Time
	setGlobalTimeLeft, setGlobalrealTimeRight bool

	leftField, rightField float64
	topField, bottomField float64

	windows    map[string]*window.Window
	windows_id []string
}

func New(width, hight, candleWidth, winGap int) *container {

	c := &container{
		G:       svg.Group(),
		windows: map[string]*window.Window{},

		leftField:   50,
		rightField:  50,
		topField:    5,
		bottomField: 20,

		candleWidth: candleWidth,
		width:       float64(width),
		hight:       float64(hight),
		winGap:      float64(winGap),
	}

	st := style.Style().Fill(colors.GRAY)
	c.G.Style(st)

	c.Init()

	return c
}

func (c *container) Window(name string, space int, realYBottom, realYTop int, percentHight float64, realTimeLeftIn, realTimeRightIn time.Time) {

	c.windows_id = append(c.windows_id, name)

	d := realTimeRightIn.Sub(realTimeLeftIn)
	deltaD := time.Duration(int(float64(d) / 100.0))

	realTimeLeft := realTimeLeftIn
	if space&1 == 1 {
		realTimeLeft = realTimeLeft.Add(-1 * deltaD)
	}

	realTimeRight := realTimeRightIn
	if space&2 == 2 {
		realTimeRight = realTimeRight.Add(deltaD)
	}

	realYBottomAdd := float64(realYBottom)
	fmt.Printf("----> space&4: %d\n", space&4)
	if space&4 == 4 {
		realYBottomAdd -= float64(realYTop-realYBottom) / 10.0
	}

	fmt.Printf("----> realYBottom: %d\n", realYBottom)
	fmt.Printf("----> realYBottomAdd: %f\n", realYBottomAdd)

	realYTopAdd := float64(realYTop)
	if space&8 == 8 {
		realYTopAdd += float64(realYTop-realYBottom) / 10.0
	}

	width := c.width - c.leftField - c.rightField
	hight := percentHight * (c.hight - c.topField - c.bottomField) / 100.0
	c.windows[name] = window.New(realYBottomAdd, realYTopAdd, width, hight, percentHight, realTimeLeft, realTimeRight)

	c.GetGlobalPoints()
}

// func (c *container) Candle(name string, x time.Time, open, clos, high, low int) {
// 	c.windows[name].Plast.RealCandle(x, c.candleWidth, open, clos, high, low)
// }

// func (c *container) Volume(name string, x time.Time, volume int) {
// 	c.windows[name].Plast.RealVolume(x, c.candleWidth, volume)
// }

func (c *container) Complete() {
	for _, w := range c.windows {
		c.G.Append(w.Plast.G)
	}
}

func (c *container) Init() {
	c.Border()
	c.finishWindows()
	c.DrawScaleInterval()
}

func (c *container) Move(x, y float64) {
	tr := transform.Transform().Translate(x, y)
	c.G.Transform(tr)
}

func (c *container) W() float64 {
	return c.width
}

func (c *container) H() float64 {
	return c.hight
}

func (c *container) Border() {
	st2 := style.Style().StrokeWidth(2).Stroke("black")
	rec2 := svg.Rect(1, 1, c.width, c.hight, st2)
	c.G.Append(rec2)
}

// func (c *container) DrawSmoothLine(name string, list []Point) {
// 	c.windows[name].Plast.SmoothLine(c.candleWidth, list)
// }

func (c *container) DrawScaleInterval() {
	for name := range c.windows {
		c._drawScaleInterval(name)
	}
}

func (c *container) _drawScaleInterval(name string) {

	baseFont := "font-family:Verdana;font-size:10px;"
	fBottom := style.Font(baseFont + "text-anchor:middle;dominant-baseline:text-before-edge")
	fYLeft := style.Font(baseFont + "text-anchor:right;dominant-baseline:middle")
	fYRight := style.Font(baseFont + "text-anchor:left;dominant-baseline:middle")
	stTextBottom := style.Style().Fill("black").Font(fBottom).StrokeWidth(1)
	stTextyLeft := style.Style().Fill("black").Font(fYLeft).StrokeWidth(1)
	stTextyRight := style.Style().Fill("black").Font(fYRight).StrokeWidth(1)

	w := c.windows[name]

	for _, v := range w.GridListY() {
		textL := svg.Text(stTextyLeft).XY(float64(c.width-c.rightField+10), v.PosDraw+w.StartTop()).String(fmt.Sprintf("%d", int(v.PosReal)))
		textR := svg.Text(stTextyRight).XY(float64(c.leftField-30.0), v.PosDraw+w.StartTop()).String(fmt.Sprintf("%d", int(v.PosReal)))
		c.G.Append(textL)
		c.G.Append(textR)
	}

	for _, v := range w.GridListX() {
		text := svg.Text(stTextBottom).XY(c.leftField+v.PosDraw, c.hight-c.bottomField+3).String(v.PosTime.Format("15:04"))
		c.G.Append(text)
	}
}

func (c *container) finishWindows() {
	for _, name := range c.windows_id {
		c.windows[name].Finish(float64(c.leftField))
	}
}

func (c *container) GetGlobalPoints() {
	for _, name := range c.windows_id {
		w := c.windows[name]

		lt := w.LeftTime()
		rt := w.RightTime()

		if !c.setGlobalTimeLeft {
			c.globalTimeLeft = lt
		} else if c.globalTimeLeft.After(lt) {
			c.globalTimeLeft = lt
		}

		if !c.setGlobalrealTimeRight {
			c.globalrealTimeRight = rt
		} else if c.globalrealTimeRight.Before(rt) {
			c.globalrealTimeRight = rt
		}

		c.setGlobalTimeLeft = true
		c.setGlobalrealTimeRight = true
	}

	h := c.hight - c.topField - c.bottomField - float64(len(c.windows_id)-1)*c.winGap
	startTop := float64(c.topField)
	for i, name := range c.windows_id {

		w := c.windows[name]

		w.Hight(w.PercentHight() * h / 100.0)
		w.RightTime(c.globalrealTimeRight)
		w.LeftTime(c.globalTimeLeft)
		w.StartTop(startTop)

		startTop += float64(w.Hight()) + c.winGap

		fmt.Printf("%d [%s]: GetGlobalPoints -> w.startTop: %f\n", i, name, w.StartTop())
	}
}

//
func (c *container) StockCandle(name string, t time.Time, cWidth, open, clos, high, low int) error {
	w, ok := c.windows[name]
	if !ok {
		return fmt.Errorf("window %s not found", name)
	}
	stock.Candle(w.Plast, t, cWidth, open, clos, high, low)
	return nil
}
