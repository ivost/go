package store

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)


func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given tsv file parse and return array of POI", t, func() {

		Convey("positive Parse test", func() {
			filename := "../data/offers_poiA.tsv"
			if recs, err := Parse(filename); err == nil {
				//log.Printf("%v", recs)
				So(len(recs), ShouldBeGreaterThan, 0)
				for _, rec := range recs {
					log.Printf("%v\n", rec)
					So(len(rec.Name) > 0, ShouldBeTrue)
					So(rec.Lat > 30, ShouldBeTrue)
					So(rec.Lng < -90, ShouldBeTrue)
					So(rec.Radius > 0, ShouldBeTrue)
				}
			} else {
				t.Fail()
			}
		})

		Convey("not found file test", func() {
			filename := "../data/notfound.tsv"
			recs, err := Parse(filename)
			log.Printf("%v\n", err)
			So(len(recs) == 0, ShouldBeTrue)
			So(err, ShouldNotBeNil)
		})

	})

	Convey("Given tsv file process it", t, func() {

		Convey("positive Process test", func() {
			filename := "../data/offers_poiA.tsv"
			err := Process(filename)
			So(err, ShouldBeNil)
		})
	})


}

