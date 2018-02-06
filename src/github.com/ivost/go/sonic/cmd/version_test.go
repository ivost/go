package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Just start of the package.json - we neet the version
type NpmPackage struct {
	Name    string `json:"name"`
	Desc    string `json:"description"`
	Version string `json:"version"`
}

// TestMain is main test
func TestVersion(t *testing.T) {
	ok := true
	dir, err := os.Getwd()
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Fatal(err)
	}
	// need -v to see output
	t.Logf("Current dir %q\n", dir)
	s := filepath.Join(dir, "../internal/testdata/package.json")
	t.Logf("path %q\n", s)
	ok, s, err = getFile(s)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatalf("File %v not found\n", s)
	}
	t.Logf("file %q\n", s)
}

// func NewTestCase(t *testing.T, dir, name string) {
// 	rootPath := filepath.FromSlash(filepath.Join(dir, name))
// 	fmt.Print(rootPath)
// n := &TestCase{
// 	t:           t,
// 	name:        name,
// 	rootPath:    rootPath,
// 	initialPath: filepath.Join(rootPath, "initial"),
// 	finalPath:   filepath.Join(rootPath, "final"),
// }
// j, err := ioutil.ReadFile(filepath.Join(rootPath, "testcase.json"))
// if err != nil {
// 	t.Fatal(err)
// }
// err = json.Unmarshal(j, n)
// if err != nil {
// 	t.Fatal(err)
// }
// return n
//}

// // InitialPath represents the initial set of files in a project.
// func (tc *TestCase) InitialPath() string {
// 	return tc.initialPath
// }

// func (tc *TestCase) CompareFile(goldenPath, working string) {
// 	golden := filepath.Join(tc.finalPath, goldenPath)

// 	gotExists, got, err := getFile(working)
// 	if err != nil {
// 		tc.t.Fatalf("Error reading project file %q: %s", goldenPath, err)
// 	}
// 	wantExists, want, err := getFile(golden)
// 	if err != nil {
// 		tc.t.Fatalf("Error reading testcase file %q: %s", goldenPath, err)
// 	}

// 	if wantExists && gotExists {
// 		if want != got {
// 			tc.t.Errorf("%s was not as expected\n(WNT):\n%s\n(GOT):\n%s", filepath.Base(goldenPath), want, got)
// 		}
// 	} else if !wantExists && gotExists {
// 		tc.t.Errorf("%q created where none was expected", goldenPath)
// 	} else if wantExists && !gotExists {
// 		tc.t.Errorf("%q not created where one was expected", goldenPath)
// 	}
// }

func getFile(path string) (bool, string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, "", nil
	}
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return true, "", err
	}
	return true, string(f), nil
}
