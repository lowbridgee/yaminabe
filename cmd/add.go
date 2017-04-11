// Copyright © 2017 Lowbridgee
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

	"github.com/lowbridgee/yaminabe/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a memo to yaminabe",
	Long:  `Add a memo to yaminabe to list up tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		memo, finished, due_at, priority, tag := "", 0, "2000-01-01", "", ""
		fmt.Println("--Add memo to database--")
		fmt.Print("Memo: ")
		fmt.Scanln(&memo)

		fmt.Print("Finished?: (true = 1 or false = 0) (default 0) ")
		fmt.Scanln(&finished)

		fmt.Print("Due_at: (YYYY-MM-DD) (default 2000-01-01) ")
		fmt.Scanln(&due_at)

		fmt.Print("Priority: (Red or Yellow or Green) ")
		fmt.Scanln(&priority)

		fmt.Print("Tag: (カンマ区切り ")
		fmt.Scanln(&tag)

		data := database.Data{Memo: memo, Finished: finished, Due_at: due_at, Priority: priority, Tag: tag}
		database.AddDB(data)
		fmt.Println("--Added complete--")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
