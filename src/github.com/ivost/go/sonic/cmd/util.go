package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	// "github.com/golang/dep/gps"
	"github.com/Masterminds/vcs"
)

type cmd struct {
	// ctx is provided by the caller; SIGINT is sent when it is cancelled.
	ctx context.Context
	Cmd *exec.Cmd
}

func isGitFound() bool {
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Printf("git binary not found\n")
		return false
	}
	return true
}

func runme() {

	ok := isGitFound()
	fmt.Printf("git found: %v\n", ok)

	//remote := "https://github.com/Masterminds/vcs"
	// remote := "https://github.com/Masterminds/vcs"
	// access denied - needs pw
	//remote := "https://Ivostoysd@bitbucket.org/smartdrive_systems/gotest2.git"

	remote := "git@bitbucket.org:smartdrive_systems/gotest2.git"
	local, _ := ioutil.TempDir("/tmp", "gotest")
	repo, err := vcs.NewRepo(remote, local)
	if err != nil {
		fmt.Printf("ERROR %q\n", err)
		return
	}
	ok = repo.Ping()
	fmt.Printf("REPO %v, PING: %v\n\n", local, ok)
	// ok = repo.CheckLocal()
	// fmt.Printf("ok %v\n", ok)

	fmt.Printf("running git version\n")
	cmd := exec.Command("git", "version")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sync result: %q\n", out.String())

	// latest, _ := repo.Date()
	// fmt.Printf("latest commit %v\n", latest)
}
