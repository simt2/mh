// Copyright © 2018 Cisco Systems, Inc.
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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply apps",
	Long: `Apply one or more MultiHelm apps. If you do not specify one or more
apps, MultiHelm acts on all apps in your MultiHelm config.`,
	Run: func(cmd *cobra.Command, args []string) {
		lateInit("apply")

		apps := getApps(args)

		configFile := getConfigFile()
		appSources := getAppSources()
		printRendered := getPrintRendered()

		apps.Apply(configFile, appSources, printRendered)
	},
}

func init() {
	RootCmd.AddCommand(applyCmd)

	applyCmd.PersistentFlags().BoolP("printRendered", "p", false, "print rendered override values")
	viper.BindPFlags(applyCmd.PersistentFlags())
}
