package haversine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	d := DistanceEarth(Pos{
		Lat: 10,
		Lon: 15,
	}, Pos{
		Lat: 20,
		Lon: 25,
	})
	assert.InDelta(t, 1544.76, d, 0.001)
}

func TestMinLatDistance(t *testing.T) {
	d := MinLatDistance(Pos{
		Lat: 10,
		Lon: 15,
	}, 25)
	assert.InDelta(t, 10.151081711048132, d, 0.001)

	d = MinLatDistance(Pos{
		Lat: 30,
		Lon: 8,
	}, 25)
	assert.InDelta(t, 31.120657588976233, d, 0.001)

	origin := Pos{
		Lat: -50,
		Lon: 8,
	}
	refLon := 25.0
	d = MinLatDistance(origin, refLon)

	const expectedLat = -51.255223591747246
	assert.InDelta(t, expectedLat, d, 0.001)

	minLat := -85.0
	minDistance := math.MaxFloat64

	for lat := -85.0; lat <= 85; lat += 0.0001 {
		d := DistanceEarth(origin, Pos{
			Lat: lat,
			Lon: refLon,
		})
		if d < minDistance {
			minDistance = d
			minLat = lat
		}
	}
	fmt.Println(minDistance, minLat)

	d = DistanceEarth(origin, Pos{
		Lat: expectedLat,
		Lon: refLon,
	})
	fmt.Println("Computed Min Distance", d)
}
