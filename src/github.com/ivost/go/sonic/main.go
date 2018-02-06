package main

import "github.com/ivost/go/sonic/cmd"

func main() {
	cmd.Execute()
}

/*
 old sonic
    -v, --own_version      display own version
	-V, --package_version  display version from package.json

    -1, --inc_major        increments major version
    -2, --inc_minor        increments minor version
	-3, --inc_patch        increments patch

    -p, --publish          yarn publish using modified version
	-h, --help             output usage information
*/
