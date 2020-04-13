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
	"errors"
	"sync"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// isoImageCmd represents the command to manage SAKURA Cloud ISOImage
func isoImageCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "iso-image",
		Short: "A manage commands of ISOImage",
		Long:  `A manage commands of ISOImage`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func isoImageListCmd() *cobra.Command {
	isoImageListParam := params.NewListISOImageParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find", "selector"},
		Short:        "List ISOImage",
		Long:         `List ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageListParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageListParam)
			}

			return funcs.ISOImageList(ctx, isoImageListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&isoImageListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &isoImageListParam.Id), "id", "", "set filter by id(s)")
	fs.StringVarP(&isoImageListParam.Scope, "scope", "", "", "set filter by scope('user' or 'shared')")
	fs.StringSliceVarP(&isoImageListParam.Tags, "tags", "", []string{}, "set filter by tags(AND) (aliases: selector)")
	fs.IntVarP(&isoImageListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&isoImageListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&isoImageListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&isoImageListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&isoImageListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&isoImageListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&isoImageListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&isoImageListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&isoImageListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&isoImageListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&isoImageListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(isoImageListNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageListFlagOrder(cmd))
	return cmd
}

func isoImageCreateCmd() *cobra.Command {
	isoImageCreateParam := params.NewCreateISOImageParam()
	cmd := &cobra.Command{
		Use: "create",

		Short:        "Create ISOImage",
		Long:         `Create ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageCreateParam)
			}

			// confirm
			if !isoImageCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.ISOImageCreate(ctx, isoImageCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&isoImageCreateParam.Size, "size", "", 5, "set iso size(GB)")
	fs.StringVarP(&isoImageCreateParam.ISOFile, "iso-file", "", "", "set iso image file")
	fs.StringVarP(&isoImageCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&isoImageCreateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.StringSliceVarP(&isoImageCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &isoImageCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&isoImageCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&isoImageCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&isoImageCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&isoImageCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&isoImageCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&isoImageCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&isoImageCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&isoImageCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&isoImageCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(isoImageCreateNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageCreateFlagOrder(cmd))
	return cmd
}

func isoImageReadCmd() *cobra.Command {
	isoImageReadParam := params.NewReadISOImageParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read ISOImage",
		Long:         `Read ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findISOImageReadTargets(ctx, isoImageReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				isoImageReadParam.SetId(id)
				go func(p *params.ReadISOImageParam) {
					err := funcs.ISOImageRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(isoImageReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&isoImageReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&isoImageReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&isoImageReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&isoImageReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&isoImageReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&isoImageReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&isoImageReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&isoImageReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&isoImageReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &isoImageReadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(isoImageReadNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageReadFlagOrder(cmd))
	return cmd
}

func isoImageUpdateCmd() *cobra.Command {
	isoImageUpdateParam := params.NewUpdateISOImageParam()
	cmd := &cobra.Command{
		Use: "update",

		Short:        "Update ISOImage",
		Long:         `Update ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findISOImageUpdateTargets(ctx, isoImageUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !isoImageUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				isoImageUpdateParam.SetId(id)
				go func(p *params.UpdateISOImageParam) {
					err := funcs.ISOImageUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(isoImageUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&isoImageUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&isoImageUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&isoImageUpdateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.StringSliceVarP(&isoImageUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &isoImageUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&isoImageUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&isoImageUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&isoImageUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&isoImageUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&isoImageUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&isoImageUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&isoImageUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&isoImageUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&isoImageUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &isoImageUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(isoImageUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageUpdateFlagOrder(cmd))
	return cmd
}

func isoImageDeleteCmd() *cobra.Command {
	isoImageDeleteParam := params.NewDeleteISOImageParam()
	cmd := &cobra.Command{
		Use:          "delete",
		Aliases:      []string{"rm"},
		Short:        "Delete ISOImage",
		Long:         `Delete ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findISOImageDeleteTargets(ctx, isoImageDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !isoImageDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				isoImageDeleteParam.SetId(id)
				go func(p *params.DeleteISOImageParam) {
					err := funcs.ISOImageDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(isoImageDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&isoImageDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&isoImageDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&isoImageDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&isoImageDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&isoImageDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&isoImageDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&isoImageDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&isoImageDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&isoImageDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&isoImageDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &isoImageDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(isoImageDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageDeleteFlagOrder(cmd))
	return cmd
}

func isoImageUploadCmd() *cobra.Command {
	isoImageUploadParam := params.NewUploadISOImageParam()
	cmd := &cobra.Command{
		Use: "upload",

		Short:        "Upload ISOImage",
		Long:         `Upload ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageUploadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageUploadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageUploadParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageUploadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findISOImageUploadTargets(ctx, isoImageUploadParam)
			if err != nil {
				return err
			}

			// confirm
			if !isoImageUploadParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("upload", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				isoImageUploadParam.SetId(id)
				go func(p *params.UploadISOImageParam) {
					err := funcs.ISOImageUpload(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(isoImageUploadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&isoImageUploadParam.ISOFile, "iso-file", "", "", "set iso image file")
	fs.StringSliceVarP(&isoImageUploadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&isoImageUploadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&isoImageUploadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageUploadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageUploadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageUploadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageUploadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&isoImageUploadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&isoImageUploadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&isoImageUploadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&isoImageUploadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&isoImageUploadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&isoImageUploadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&isoImageUploadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &isoImageUploadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(isoImageUploadNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageUploadFlagOrder(cmd))
	return cmd
}

func isoImageDownloadCmd() *cobra.Command {
	isoImageDownloadParam := params.NewDownloadISOImageParam()
	cmd := &cobra.Command{
		Use: "download",

		Short:        "Download ISOImage",
		Long:         `Download ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageDownloadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageDownloadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageDownloadParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageDownloadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findISOImageDownloadTargets(ctx, isoImageDownloadParam)
			if err != nil {
				return err
			}

			// confirm
			if !isoImageDownloadParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("download", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				isoImageDownloadParam.SetId(id)
				go func(p *params.DownloadISOImageParam) {
					err := funcs.ISOImageDownload(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(isoImageDownloadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&isoImageDownloadParam.FileDestination, "file-destination", "", "", "set file destination path")
	fs.StringSliceVarP(&isoImageDownloadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&isoImageDownloadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&isoImageDownloadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageDownloadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageDownloadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageDownloadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageDownloadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &isoImageDownloadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(isoImageDownloadNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageDownloadFlagOrder(cmd))
	return cmd
}

func isoImageFTPOpenCmd() *cobra.Command {
	isoImageFTPOpenParam := params.NewFTPOpenISOImageParam()
	cmd := &cobra.Command{
		Use: "ftp-open",

		Short:        "FTPOpen ISOImage",
		Long:         `FTPOpen ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageFTPOpenParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageFTPOpenParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageFTPOpenParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageFTPOpenParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findISOImageFTPOpenTargets(ctx, isoImageFTPOpenParam)
			if err != nil {
				return err
			}

			// confirm
			if !isoImageFTPOpenParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ftp-open", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				isoImageFTPOpenParam.SetId(id)
				go func(p *params.FTPOpenISOImageParam) {
					err := funcs.ISOImageFTPOpen(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(isoImageFTPOpenParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&isoImageFTPOpenParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&isoImageFTPOpenParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&isoImageFTPOpenParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageFTPOpenParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageFTPOpenParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageFTPOpenParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageFTPOpenParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&isoImageFTPOpenParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&isoImageFTPOpenParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&isoImageFTPOpenParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&isoImageFTPOpenParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&isoImageFTPOpenParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&isoImageFTPOpenParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&isoImageFTPOpenParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &isoImageFTPOpenParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(isoImageFTPOpenNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageFTPOpenFlagOrder(cmd))
	return cmd
}

func isoImageFTPCloseCmd() *cobra.Command {
	isoImageFTPCloseParam := params.NewFTPCloseISOImageParam()
	cmd := &cobra.Command{
		Use: "ftp-close",

		Short:        "FTPClose ISOImage",
		Long:         `FTPClose ISOImage`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return isoImageFTPCloseParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, isoImageFTPCloseParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if isoImageFTPCloseParam.GenerateSkeleton {
				return generateSkeleton(ctx, isoImageFTPCloseParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findISOImageFTPCloseTargets(ctx, isoImageFTPCloseParam)
			if err != nil {
				return err
			}

			// confirm
			if !isoImageFTPCloseParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ftp-close", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				isoImageFTPCloseParam.SetId(id)
				go func(p *params.FTPCloseISOImageParam) {
					err := funcs.ISOImageFTPClose(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(isoImageFTPCloseParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&isoImageFTPCloseParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&isoImageFTPCloseParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&isoImageFTPCloseParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&isoImageFTPCloseParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&isoImageFTPCloseParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&isoImageFTPCloseParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&isoImageFTPCloseParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &isoImageFTPCloseParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(isoImageFTPCloseNormalizeFlagNames)
	buildFlagsUsage(cmd, isoImageFTPCloseFlagOrder(cmd))
	return cmd
}