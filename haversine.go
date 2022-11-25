package haversine

import (
	"math"
)

type Pos struct {
	Lat float64 // in degrees
	Lon float64 // in degrees
}

func degreeToRadian(d float64) float64 {
	return d * math.Pi / 180
}

func posDegreeToRadian(p Pos) Pos {
	return Pos{
		Lat: degreeToRadian(p.Lat),
		Lon: degreeToRadian(p.Lon),
	}
}

func radianToDegree(r float64) float64 {
	return r * 180 / math.Pi
}

// Distance computes the Haversine distance for a sphere with unit radius
func Distance(a, b Pos) float64 {
	a = posDegreeToRadian(a)
	b = posDegreeToRadian(b)

	firstSin := math.Sin((b.Lat - a.Lat) / 2)
	firstSin = firstSin * firstSin

	secondSin := math.Sin((b.Lon - a.Lon) / 2)
	secondSin = secondSin * secondSin

	sum := firstSin + secondSin*math.Cos(a.Lat)*math.Cos(b.Lat)

	return 2 * math.Asin(math.Sqrt(sum))
}

// DistanceRadius computes the Haversine distance for a sphere with the radius
func DistanceRadius(a, b Pos, radius float64) float64 {
	return Distance(a, b) * radius
}

const earthRadius = 6371.009

// DistanceEarth computes the Haversine distance for a sphere with the radius = 6371.009 km
func DistanceEarth(a, b Pos) float64 {
	return DistanceRadius(a, b, earthRadius)
}

// MinLatDistance ...
func MinLatDistance(origin Pos, lon float64) float64 {
	origin = posDegreeToRadian(origin)
	lon = degreeToRadian(lon)

	cosLon := math.Cos(lon - origin.Lon)
	latRadian := math.Atan(math.Tan(origin.Lat) / cosLon)
	return radianToDegree(latRadian)
}
