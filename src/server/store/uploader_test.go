package store

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)


func TestSpec(t *testing.T) {

	//// Only pass t into top-level Convey calls
	//Convey("Given some integer with a starting value", t, func() {
	//	x := 1
	//
	//	Convey("When the integer is incremented", func() {
	//		x++
	//
	//		Convey("The value should be greater by one", func() {
	//			So(x, ShouldEqual, 2)
	//		})
	//	})
	//})

	// Only pass t into top-level Convey calls
	Convey("Given tsv file parse into 2 dim array", t, func() {

		Convey("pos test", func() {
			filename := "../data/offers_poiA.tsv"
			if r, err := Parse(filename); err != nil {
				log.Printf("%v", r)
				So(len(r), ShouldBeGreaterThan, 0)
			}
		})

		Convey("neg test", func() {
			filename := "../data/notfound.tsv"
			_, err := Parse(filename)
			So(err, ShouldNotBeNil)
		})

	})

}

//func Test_parse(test * testing)  {
//	r := parse(filename)
//	log.Printf("%v", r)
//}
