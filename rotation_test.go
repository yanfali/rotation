package rotation_test

import (
	"math"
	"testing"
)

func Rotate(x, y int, radian float64) (int, int) {
	var xp, yp float64
	xp = float64(x)*math.Cos(radian) - float64(y)*math.Sin(radian)
	yp = float64(x)*math.Sin(radian) + float64(y)*math.Cos(radian)
	return int(Round(xp)), int(Round(yp))
}

func DegreeToRadian(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func Round(n float64) float64 {
	if n >= 0.5 {
		return math.Trunc(n + 0.5)
	}
	if n <= -0.5 {
		return math.Trunc(n - 0.5)
	}
	if math.IsNaN(n) {
		return math.NaN()
	}
	return 0
}

type testCase struct {
	X      int
	Y      int
	Radian float64
	Xp     int
	Yp     int
}

func TestRotation(t *testing.T) {
	radian := DegreeToRadian(-90)
	testCases := []testCase{
		{
			X:      1,
			Y:      1,
			Radian: radian,
			Xp:     1,
			Yp:     -1,
		},
		{
			X:      4,
			Y:      1,
			Radian: radian,
			Xp:     1,
			Yp:     -4,
		},
		{
			X:      1,
			Y:      4,
			Radian: radian,
			Xp:     4,
			Yp:     -1,
		},
		{
			X:      4,
			Y:      4,
			Radian: radian,
			Xp:     4,
			Yp:     -4,
		},
	}
	for _, tc := range testCases {
		xp, yp := Rotate(tc.X, tc.Y, tc.Radian)
		if xp != tc.Xp || yp != tc.Yp {
			t.Fatalf("for input %d, %d expected %d, %d but received %d, %d", tc.X, tc.Y, tc.Xp, tc.Yp, xp, yp)
		}
	}
}
