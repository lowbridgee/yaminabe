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

var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish a task with id",
	Long:  `Finish a task with id.`,
	Run: func(cmd *cobra.Command, args []string) {
		id := 0
		fmt.Print("What is ID?: ")
		fmt.Scanln(&id)
		database.Finish(id)
		fmt.Println("--Finish changed--")
	},
}

func init() {
	RootCmd.AddCommand(finishCmd)
}
