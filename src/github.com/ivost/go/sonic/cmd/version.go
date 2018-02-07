// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display package version",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command`,
	Run: func(cmd *cobra.Command, args []string) {
		VERSION := "1.0"
		fmt.Println("version called")
		if viper.GetBool("own") {
			fmt.Printf("sonic version %v\n", VERSION)
			return
		}
		var pkg NpmPackage
		path := ""
		_ = readNpmPackage(path, &pkg)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly
	versionCmd.Flags().BoolP("own", "t", false, "sonic version")
	viper.BindPFlag("own", versionCmd.Flags().Lookup("own"))
}

// Just start of the package.json - we neet the version
type NpmPackage struct {
	Name    string `json:"name"`
	Desc    string `json:"description"`
	Version string `json:"version"`
}

func readNpmPackage(path string, pkg *NpmPackage) bool {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	err = json.Unmarshal([]byte(string(f)), &pkg)
	return err == nil
}
