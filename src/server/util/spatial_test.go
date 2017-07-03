package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
	"math"
)


func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("spatial tests", t, func() {

		Convey("ToRadians test", func() {
			r1 := ToRadians(180.)
			//log.Printf("r1 %v", r1)
			So(r1, ShouldAlmostEqual, math.Pi)

			r1 = ToRadians(0.)
			//log.Printf("r1 %v", r1)
			So(r1, ShouldAlmostEqual, 0)

			r1 = ToRadians(270.)
			//log.Printf("r1 %v", r1)
			So(r1, ShouldAlmostEqual, (3./2.) * math.Pi)
		})

		Convey("distance test1", func() {
			lat1 := 32.0
			lon1 := -117.0
			lat2 := 32.0
			lon2 := -117.0

			dist := Dist(lat1, lon1, lat2, lon2)
			//log.Printf("dist %v", dist)
			So(dist, ShouldAlmostEqual, 0)
		})

		Convey("distance test2", func() {
			// SD 32.7157° N, 117.1611° W
			lat1 := 32.7157
			lon1 := -117.1611
			// LA 34.0522° N, 118.2437° W
			lat2 := 34.0522
			lon2 := -118.2437

			dist := Dist(lat1, lon1, lat2, lon2)
			//log.Printf("dist %f km", dist/1000.)	// km
			//mi := MeterToMile(dist)
			log.Printf("dist %f mi", dist)	// mi
			So(int(dist), ShouldAlmostEqual, 111)
		})

	})



}

