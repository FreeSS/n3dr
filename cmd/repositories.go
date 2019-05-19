// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"n3dr/cli"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	names    bool
	count    bool
	download bool
)

// repositoriesCmd represents the repositories command
var repositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "Count the number of repositories or return their names",
	Long: `Count the number of repositories, count the total or
download artifacts from all repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		if names {
			cli.RepositoryNames()
		}
		if count {
			cli.CountRepositories()
		}
		if download {
			na := cli.Nexus3All{URL: n3drURL, User: n3drUser, Pass: n3drPass}
			err := na.Downloads()
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(repositoriesCmd)
	repositoriesCmd.Flags().BoolVarP(&names, "names", "a", false, "Print all repository names")
	repositoriesCmd.Flags().BoolVarP(&count, "count", "c", false, "Count the number of repositories")
	repositoriesCmd.Flags().BoolVarP(&download, "download", "d", false, "Download artifacts from all repositories")
	repositoriesCmd.Flags().StringVarP(&n3drURL, "n3drURL", "n", "http://localhost:8081", "The Nexus3 URL")
	repositoriesCmd.Flags().StringVarP(&n3drPass, "n3drPass", "p", "admin123", "The Nexus3 password")
	repositoriesCmd.Flags().StringVarP(&n3drUser, "n3drUser", "u", "admin", "The Nexus3 user")
}