package grid

import (
	"math"
	"time"

	"github.com/iostrovok/svg-charts/points"
)

var roundInt []int = []int{
	10, 20, 25, 50,
}

var roundTime []time.Duration = []time.Duration{
	0 * time.Minute,
	1 * time.Minute,
	2 * time.Minute,
	5 * time.Minute,
	10 * time.Minute,
	15 * time.Minute,
	30 * time.Minute,
	1 * time.Hour,
	2 * time.Hour,
	3 * time.Hour,
	4 * time.Hour,
	5 * time.Hour,
	6 * time.Hour,
	12 * time.Hour,
	1 * 24 * time.Hour,
	2 * 24 * time.Hour,
	3 * 24 * time.Hour,
	4 * 24 * time.Hour,
	7 * 24 * time.Hour,
	10 * 24 * time.Hour,
	2 * 7 * 24 * time.Hour,
	3 * 7 * 24 * time.Hour,
	4 * 7 * 24 * time.Hour,
	365 * 24 * time.Hour,
}

func GridByTime(
	startTimeX, finishTimeX time.Time,
	startY, finishY,
	countPerX, countPerY,
	width, hight int,
) ([]points.GridPointTime, []points.GridPointNum) {
	listY := GetLineGridPoints(startY, finishY, hight, countPerY)
	listX := GetTimeGridPoints(startTimeX, finishTimeX, width, countPerX)
	return listX, listY
}

func findBestGridIntPoint(d int) int {
	add := 1
	for {
		for _, i := range roundInt {
			t := i * add / 10
			if d < t {
				return t
			}
		}
		add *= 10
	}
	return 1
}

func GetLineGridPoints(from, to, lenSide, gridCount int) []points.GridPointNum {

	fullD := to - from
	d := fullD / gridCount
	step := findBestGridIntPoint(d)

	startPoint := step * int(math.Floor(float64(from)/float64(step)))
	out := []points.GridPointNum{}
	for startPoint < to {

		x := int(float64(lenSide) * float64(startPoint-from) / float64(fullD))

		if 0 < x && x < lenSide {
			out = append(out, points.GridPointNum{
				PosDraw: float64(lenSide - x),
				PosReal: float64(startPoint),
			})
		}
		startPoint += step
	}
	return out
}

func findBestGridTimePoint(d time.Duration) time.Duration {
	for i := 0; i < len(roundTime)-1; i++ {
		if roundTime[i] < d && d <= roundTime[i+1] {
			return roundTime[i+1]
		}
	}
	return roundTime[1]
}

func GetTimeGridPoints(t0, t1 time.Time, lenSide, gridCount int) []points.GridPointTime {

	fullD := t1.Sub(t0)
	d := findBestGridTimePoint(t1.Sub(t0) / time.Duration(gridCount))

	startPoint := t0.Add(time.Duration(float64(d) / 1.8)).Round(d)
	out := []points.GridPointTime{}
	for startPoint.Before(t1) {

		l := startPoint.Sub(t0)

		x := float64(lenSide) * float64(l) / float64(fullD)
		out = append(out, points.GridPointTime{
			PosDraw: x,
			PosTime: t0.Add(l),
		})
		startPoint = startPoint.Add(d)
	}
	return out
}
