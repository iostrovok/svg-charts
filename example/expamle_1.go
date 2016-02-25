package main

import (
	"fmt"
	"log"
	// "math"
	"os"
	"time"

	"github.com/iostrovok/svg"
	"github.com/iostrovok/svg/style"

	chart "github.com/iostrovok/svg-charts"
)

type Point struct {
	x string
	y float64
}

type Candle struct {
	x          string
	open, clos float64
	high, low  float64
}

const (
	WIDTH = 6200
	HIGH  = 650

	candleWidth     = 6
	betweenCandle   = 8
	fieldSpaceHigh  = 30
	fieldSpaceWidth = 30

	TimePayShift   = 5000 * time.Millisecond // millisecond
	GoodStatusFile = 1
	DEBUG          = false
	// TIME_FRAME = 1 * time.Minute

	CANDEL_INFO     = "%s\nOpen: %d\nClose: %d\nLow: %d\nHigh: %d\nDelta: %d\nMiddel: %d\n"
	CANDEL_INFO_ADD = "%s\n-------\nx1: %d\ny1: %d\ny2: %d\nh_y1: %d\nh_y2: %d\n"
	RED             = "rgb(220,60,60)"
	GREEN           = "rgb(60,220,60)"
	LIGHT_GRAY      = "rgb(240,240,240)"
	GRAY            = "rgb(196,196,196)"

	FILE_OUT = "./index.svg"
)

func _time_test() (time.Time, time.Time) {
	t0, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 15:00:00")
	if err != nil {
		log.Fatalln(err)
	}
	t1, err := time.Parse("2006-01-02 15:04:05", "2015-02-01 16:00:00")
	if err != nil {
		log.Fatalln(err)
	}
	return t0, t1
}

func testHistogram() []Point {

	return []Point{
		{x: "2015-02-01 15:00:00", y: 100},
		{x: "2015-02-01 15:05:00", y: 444},
		{x: "2015-02-01 15:10:00", y: 678},
		{x: "2015-02-01 15:20:00", y: 1000},
		{x: "2015-02-01 15:25:00", y: 787},
		{x: "2015-02-01 15:30:00", y: 909},
		{x: "2015-02-01 15:35:00", y: 1000},
		{x: "2015-02-01 15:40:00", y: 567},
		{x: "2015-02-01 15:45:00", y: 456},
		{x: "2015-02-01 15:50:00", y: 123},
		{x: "2015-02-01 15:55:00", y: 12},
	}
}

func testCandles() []Candle {
	return []Candle{
		{x: "2015-02-01 15:00:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:05:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:10:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:20:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:25:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:30:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:35:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:40:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:45:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:50:00", open: 100, clos: 100, high: 200, low: 100},
		{x: "2015-02-01 15:55:00", open: 100, clos: 100, high: 200, low: 100},
	}
}

func main() {

	g := chart.New(WIDTH-100, HIGH-50, candleWidth, 5)

	realYBottom := 100
	realYTop := 1000
	realTopVolume := 200
	realTimeLeft, realTimeRight := _time_test()

	g.Window("candles", 1+2+4+8, realYBottom, realYTop, 50.0, realTimeLeft, realTimeRight)
	g.Window("candles2", 1+2+4+8, realYBottom, realYTop, 25.0, realTimeLeft, realTimeRight)
	g.Window("volume", 1+2+4+8, 0, realTopVolume, 25.0, realTimeLeft, realTimeRight)
	g.Init()
	g.Move(10, 10)

	listCandle := testCandles()
	for i, c := range listCandle {

		t1, err := time.Parse("2006-01-02 15:04:05", c.x)
		if err != nil {
			log.Fatalln(err)
		}

		// fmt.Printf("%s. Open: %d, Close: %d, High: %d, Low: %d\n", c.TimeOpen.Format("2006-01-02 15:04:05"), c.Open(), c.Close(), c.High(), c.Low())
		g.Candle("candles", t, c.open, c.clos, c.high, c.low)
		// g.Volume("volume", c.TimeOpen, c.Volume())
		g.Candle("candles2", t, c.open, c.clos, c.high, c.low)
	}

	// g.DrawSmoothLine("candles", listPoints)
	// g.DrawSmoothLine("candles2", listPoints)

	canvas := svg.New(WIDTH, HIGH)
	Border(canvas)
	canvas.Append(g.G)
	g.Complete()

	file, err := os.Create(FILE_OUT)
	if err != nil {
		log.Fatalln(err)
	}
	canvas.Save(file)

	fmt.Println(FILE_OUT)
	fmt.Println("Done")
}

func Border(canvas *svg.SVG) {
	st2 := style.Style().StrokeWidth(2).Stroke("black").Fill("none")
	canvas.Rect(1, 1, WIDTH-2, HIGH-2, st2)
}
