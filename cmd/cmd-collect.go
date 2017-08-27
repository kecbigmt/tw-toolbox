// Copyright © 2017 Kecy <keishiro.oym@gmail.com>
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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/kecbigmt/tw-toolbox/lib"
)

var exportFlag bool
var exportPath string
var showFlag bool
var language string

// collectCmd represents the collect command
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "collect tweets",
	Long: `Collect tweets match to the given keyword by the given number.

This command uses the search/tweets endpoint of Twitter REST API. More detailed info can be found at:
https://dev.twitter.com/rest/reference/get/search/tweets`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			s string
			c string
		)
		switch args_n := len(args); args_n{
		case 2:
			s = args[0]
			c = args[1]
		default:
			fmt.Printf("2 arguments required, but %d given.\n", args_n)
			return
		}
		switch {
		case exportFlag:
			tweets := lib.Collect(s, c, language)
			path := lib.ExportCsv(tweets.Header, tweets.Rows, exportPath)
			fmt.Println("exported!: ", path)
		case showFlag:
			fmt.Println(lib.Collect(s,c, language))
		default:
			fmt.Println(lib.Collect(s, "5", language))
			fmt.Println("by default, only 5 tweets are shown without flags.\nuse flag '--show' or '-s' to display all tweets.\nuse flag '--export' or '-e' to export all tweets to CSV file.")
		}
		return
	},
}

func init() {
	collectCmd.Flags().BoolVarP(&exportFlag, "export", "e", false, "export collected tweets to CSV")
	collectCmd.Flags().StringVarP(&exportPath, "path", "p", "tweets.csv", "change filepath to export CSV")
	collectCmd.Flags().BoolVarP(&showFlag, "show", "s", false, "show all collected tweets briefly in command line")
	collectCmd.Flags().StringVarP(&language, "lang", "l", "", "restrict tweets to the given language, given by an ISO 639-1 code.")
	RootCmd.AddCommand(collectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// collectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
