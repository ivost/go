package cmd

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 	go test -v to see output

func TestVersion(t *testing.T) {
	expected := "1.1.11"
	// Only pass t into top-level Convey calls
	Convey("Given some npm package json", t, func() {
		Convey("Obtain version", func() {
			var pkg NpmPackage
			dir, _ := os.Getwd()
			path := filepath.Join(dir, "../internal/testdata/package.json")
			//_ = json.Unmarshal([]byte(readNpmPackage(t)), &pkg)
			ok := readNpmPackage(path, &pkg)
			Convey("The version should be "+expected, func() {
				So(ok, ShouldBeTrue)
				So(pkg.Version, ShouldEqual, expected)
			})
		})
	})
}

// func readNpmPackage(t *testing.T) string {
// 	ok := true
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// this assumes particular level (cmd in our case)
// 	s := filepath.Join(dir, "../internal/testdata/package.json")
// 	ok, s, err = getFile(s)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !ok {
// 		t.Fatalf("File %v not found\n", s)
// 	}
// 	return s
// }

// recursively search for given target - e.g. testdata
// go up to the parent until limited number of levels
// func findTestdata(target string) (bool, string) {
// 	_, err := os.Stat(path)
// 	if err != nil {
// 		return false, "", nil
// 	}
// 	f, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return true, "", err
// 	}
// 	return true, string(f), nil
// }

// func getFile(path string) (bool, string, error) {
// 	_, err := os.Stat(path)
// 	if err != nil {
// 		return false, "", nil
// 	}
// 	f, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return true, "", err
// 	}
// 	return true, string(f), nil
// }
