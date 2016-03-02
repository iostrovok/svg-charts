package main

import (
	"fmt"
	"log"
	// "math"
	"os"
	"time"

	"github.com/iostrovok/svg"
	style "github.com/iostrovok/svg/style"

	chart "github.com/iostrovok/svg-charts"
	colors "github.com/iostrovok/svg-charts/colors"
	stock "github.com/iostrovok/svg-charts/ext/stock"
	points "github.com/iostrovok/svg-charts/points"
)

type Point struct {
	x string
	y int
}

type Candle struct {
	x          string
	open, clos int
	high, low  int
}

const (
	WIDTH = 1300
	HIGH  = 650

	candleWidth     = 50
	betweenCandle   = 8
	fieldSpaceHigh  = 30
	fieldSpaceWidth = 30

	TimePayShift   = 5000 * time.Millisecond // millisecond
	GoodStatusFile = 1
	DEBUG          = false
	// TIME_FRAME = 1 * time.Minute

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
		{x: "2015-02-01 15:05:00", y: 120},
		{x: "2015-02-01 15:10:00", y: 67},
		{x: "2015-02-01 15:15:00", y: 200},
		{x: "2015-02-01 15:20:00", y: 100},
		{x: "2015-02-01 15:25:00", y: 87},
		{x: "2015-02-01 15:30:00", y: 90},
		{x: "2015-02-01 15:35:00", y: 100},
		{x: "2015-02-01 15:40:00", y: 56},
		{x: "2015-02-01 15:45:00", y: 45},
		{x: "2015-02-01 15:50:00", y: 13},
		{x: "2015-02-01 15:55:00", y: 12},
	}
}

func testCandles() []Candle {
	return []Candle{
		{x: "2015-02-01 15:00:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:05:00", open: 500, clos: 300, high: 650, low: 50},
		{x: "2015-02-01 15:10:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:15:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:20:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:25:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:30:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:35:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:40:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:45:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:50:00", open: 100, clos: 200, high: 250, low: 50},
		{x: "2015-02-01 15:55:00", open: 100, clos: 200, high: 250, low: 50},
	}
}

func main() {

	g := chart.New(WIDTH-100, HIGH-50, 5)

	realYBottom := 100
	realYTop := 1000
	realTopVolume := 200
	realTimeLeft, realTimeRight := _time_test()

	g.Window("candles", 1+2+4+8, realYBottom, realYTop, 50.0, realTimeLeft, realTimeRight)
	g.Window("candles2", 1+2+4+8, realYBottom, realYTop, 25.0, realTimeLeft, realTimeRight)
	g.Window("volume", 1+2+8, 0, realTopVolume, 25.0, realTimeLeft, realTimeRight)
	g.Init()
	g.Move(10, 10)

	listPoints := []points.PointTime{}

	listCandle := testCandles()
	for _, c := range listCandle {

		t, err := time.Parse("2006-01-02 15:04:05", c.x)
		if err != nil {
			log.Fatalln(err)
		}

		cnd := stock.OneCandle{
			T:        t,
			Open:     c.open,
			Close:    c.clos,
			High:     c.high,
			Low:      c.low,
			Width:    candleWidth,
			StBorder: style.Style().StrokeWidth(2).Stroke("black").Fill("none").Ref(),
			StDown:   style.Style().StrokeWidth(0.5).Stroke("black").Fill(colors.LIGHT_GRAY).Ref(),
			StUp:     style.Style().StrokeWidth(0.5).Stroke("black").Fill(colors.GRAY).Ref(),
			Debug:    false,
		}

		g.StockCandle("candles", cnd)
		g.StockCandle("candles2", cnd)

		listPoints = append(listPoints, points.PointTime{
			Y: float64(c.high),
			X: t,

			DisX: float64(candleWidth) / 2,
		})
	}

	volumeCandle := testHistogram()
	for _, c := range volumeCandle {

		t, err := time.Parse("2006-01-02 15:04:05", c.x)
		if err != nil {
			log.Fatalln(err)
		}

		vol := stock.OneVolume{
			T:     t,
			Y:     c.y,
			Width: 50,
			St:    style.Style().StrokeWidth(0.5).Stroke("black").Fill(colors.GREEN).Ref(),
			Debug: true,
		}
		g.Volume("volume", vol)
	}

	g.SmoothByTime("candles", listPoints)
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
