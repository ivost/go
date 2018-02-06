package cmd

import (
	"fmt"
	"io/ioutil"

	// "github.com/golang/dep/gps"
	"github.com/Masterminds/vcs"
)

func runme() {
	//remote := "https://github.com/Masterminds/vcs"

	// remote := "https://github.com/Masterminds/vcs"

	// access denied - needs pw
	//remote := "https://Ivostoysd@bitbucket.org/smartdrive_systems/gotest2.git"

	remote := "git@bitbucket.org:smartdrive_systems/gotest2.git"

	local, _ := ioutil.TempDir("/tmp", "gotest")

	repo, err := vcs.NewRepo(remote, local)
	if err != nil {
		fmt.Printf("ERROR %v\n", err)
		return
	}

	ok := repo.Ping()

	fmt.Printf("REPO %v, PING: %v\n\n", local, ok)

	// ok = repo.CheckLocal()
	// fmt.Printf("ok %v\n", ok)

	latest, _ := repo.Date()
	fmt.Printf("latest commit %v\n", latest)
}
