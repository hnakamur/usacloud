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

// billCmd represents the command to manage SAKURA Cloud Bill
func billCmd() *cobra.Command {
	return &cobra.Command{
		Use: "bill",

		Short: "A manage commands of Bill",
		Long:  `A manage commands of Bill`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultCmd(cmd, args, "list")
		},
	}
}

func billCsvCmd() *cobra.Command {
	billCsvParam := params.NewCsvBillParam()
	cmd := &cobra.Command{
		Use: "csv",

		Short:        "Csv Bill",
		Long:         `Csv Bill`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return billCsvParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, billCsvParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if billCsvParam.GenerateSkeleton {
				return generateSkeleton(ctx, billCsvParam)
			}

			return funcs.BillCsv(ctx, billCsvParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&billCsvParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&billCsvParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&billCsvParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&billCsvParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&billCsvParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.BoolVarP(&billCsvParam.NoHeader, "no-header", "", false, "set output header flag")
	fs.StringVarP(&billCsvParam.BillOutput, "bill-output", "", "", "set bill-detail output path (aliases: file)")
	fs.VarP(newIDValue(0, &billCsvParam.BillId), "bill-id", "", "set bill ID")
	fs.SetNormalizeFunc(billCsvNormalizeFlagNames)
	buildFlagsUsage(cmd, billCsvFlagOrder(cmd))
	return cmd
}

func billListCmd() *cobra.Command {
	billListParam := params.NewListBillParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find"},
		Short:        "List Bill (default)",
		Long:         `List Bill (default)`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return billListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, billListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if billListParam.GenerateSkeleton {
				return generateSkeleton(ctx, billListParam)
			}

			return funcs.BillList(ctx, billListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&billListParam.Year, "year", "", 0, "set year")
	fs.IntVarP(&billListParam.Month, "month", "", 0, "set month")
	fs.StringVarP(&billListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&billListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&billListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&billListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&billListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&billListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&billListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&billListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&billListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&billListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&billListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&billListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(billListNormalizeFlagNames)
	buildFlagsUsage(cmd, billListFlagOrder(cmd))
	return cmd
}
