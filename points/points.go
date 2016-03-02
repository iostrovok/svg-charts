package points

import (
	"time"
)

type GridPointTime struct {
	PosDraw float64
	PosTime time.Time
}

type GridPointNum struct {
	PosDraw, PosReal float64
}

type PointTime struct {
	X time.Time
	Y float64

	// displacement
	DisX, DisY float64
}

type PointNum struct {
	X, Y float64

	// displacement
	DisX, DisY float64
}
