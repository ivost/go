package integration

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Just start of the package.json - we neet the version
type NpmPackage struct {
	Name    string `json:"name"`
	Desc    string `json:"description"`
	Version string `json:"version"`
}

// TestMain is main test
/*
It is sometimes necessary for a test program to do extra setup or teardown before or after testing. It is also sometimes necessary for a test to control which code runs on the main thread. To support these and other cases, if a test file contains a function:

func TestMain(m *testing.M)
then the generated test will call TestMain(m) instead of running the tests directly. TestMain runs in the main goroutine and can do whatever setup and teardown is necessary around a call to m.Run. It should then call os.Exit with the result of m.Run. When TestMain is called, flag.Parse has not been run. If TestMain depends on command-line flags, including those of the testing package, it should call flag.Parse explicitly.

A simple implementation of TestMain is:

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
*/
func TestMain(m *testing.M) {
	fmt.Printf("TestMain %v\n", m)
	flag.Parse()
	// if !testing.Short() {
	//     setupDatabase()
	// }
	// result := m.Run()
	// if !testing.Short() {
	//     teardownDatabase()
	// }
	result := 0
	os.Exit(result)
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
