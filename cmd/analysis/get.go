/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package analysis

import (
	"os"

	"github.com/armory-io/kayentactl/internal/options"

	"github.com/armory-io/kayentactl/internal/report"

	"github.com/armory-io/kayentactl/pkg/kayenta"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var outFormat string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [execution-id]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		globals, _ := options.Globals(cmd.Root())

		kc := kayenta.NewDefaultClient(kayenta.ClientBaseURL(globals.KayentaURL))
		executionID := args[0]
		if executionID == "" {
			log.Fatal("execution id is required")
		}
		result, err := kc.GetStandaloneCanaryAnalysis(executionID)
		if err != nil {
			log.Fatalf("failed to fetch results of analysis: %s", err.Error())
		}

		if err := report.Report(result, outFormat, os.Stdout); err != nil {
			if err == report.ErrNotComplete {
				log.Errorf("cannot generate report for running analysis %s", executionID)
			} else {
				log.Errorf("failed to generate result report: %s", err.Error())
			}
			os.Exit(1)
		}
	},
}

func init() {
	analysisCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().StringVarP(&outFormat, "output", "o", "pretty", "output format: json|pretty")
}
