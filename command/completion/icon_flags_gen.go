// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func IconListCompleteFlags(ctx command.Context, params *params.ListIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["Icon"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Icon"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "scope":
		param := define.Resources["Icon"].Commands["list"].BuildedParams().Get("scope")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags", "selector":
		param := define.Resources["Icon"].Commands["list"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["Icon"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["Icon"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["Icon"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func IconCreateCompleteFlags(ctx command.Context, params *params.CreateIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "image":
		param := define.Resources["Icon"].Commands["create"].BuildedParams().Get("image")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["Icon"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["Icon"].Commands["create"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func IconReadCompleteFlags(ctx command.Context, params *params.ReadIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["Icon"].Commands["read"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Icon"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func IconUpdateCompleteFlags(ctx command.Context, params *params.UpdateIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["Icon"].Commands["update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["Icon"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["Icon"].Commands["update"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Icon"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func IconDeleteCompleteFlags(ctx command.Context, params *params.DeleteIconParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["Icon"].Commands["delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Icon"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
