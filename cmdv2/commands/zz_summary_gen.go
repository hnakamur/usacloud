// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/spf13/cobra"
)

// summaryCmd represents the command to manage SAKURA Cloud Summary
func summaryCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "summary",
		Short: "Show summary of resource usage",
		Long:  `Show summary of resource usage`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultCmd(cmd, args, "show")
		},
	}
}

func summaryShowCmd() *cobra.Command {
	summaryShowParam := params.NewShowSummaryParam()
	cmd := &cobra.Command{
		Use: "show",

		Short:        "Show Summary (default)",
		Long:         `Show Summary (default)`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return summaryShowParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, summaryShowParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if summaryShowParam.GenerateSkeleton {
				return generateSkeleton(ctx, summaryShowParam)
			}

			return funcs.SummaryShow(ctx, summaryShowParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&summaryShowParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&summaryShowParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&summaryShowParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&summaryShowParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&summaryShowParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&summaryShowParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&summaryShowParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&summaryShowParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&summaryShowParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&summaryShowParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&summaryShowParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&summaryShowParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.BoolVarP(&summaryShowParam.PaidResourcesOnly, "paid-resources-only", "", false, "Show paid-resource only (aliases: paid)")
	fs.SetNormalizeFunc(summaryShowNormalizeFlagNames)
	buildFlagsUsage(cmd, summaryShowFlagOrder(cmd))
	return cmd
}
