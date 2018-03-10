package database

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"os"
	"testing"
	"server/model"
)


func init() {
	log.SetOutput(os.Stdout)
}

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("db test", t, func() {

		Convey("insert test", func() {
			poi := model.POI{"name", "", "", "92126" ,
				"", "619-1234-123", 32.12, -132.34, 4.9}
			AddPoi(&poi)
		})

	})
}

