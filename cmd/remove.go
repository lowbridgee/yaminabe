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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove yaminabe tasks with id.",
	Long:  `Remove yaminabe tasks with id.`,
	Run: func(cmd *cobra.Command, args []string) {
		id := 0
		fmt.Print("What is ID?: ")
		fmt.Scanln(&id)
		database.Remove(id)
		fmt.Println("--Remove complete--")
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
