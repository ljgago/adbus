// Copyright © 2019 Leonardo Javier Gago <ljgago@gmail.com>
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

package gen // import "github.com/ljgago/adbus/cmd/gen"

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ljgago/adbus/cmd"
	"github.com/ljgago/adbus/pkg/uuid"
)

// commandDefine represents the subcommand
var commandDefine = &cobra.Command{
	Use:   "gen [flags]",
	Short: "Generates a Device ID",
	Long: `
Generates a Device ID in
UUID format form a MAC address.`,
	DisableAutoGenTag: true,
	Run: func(command *cobra.Command, args []string) {
		runGen(args)
	},
}

func init() {
	cmd.Root.AddCommand(commandDefine)
}

func runGen(args []string) {
	id := uuid.GenerateDeviceID()
	fmt.Println(id)
}
