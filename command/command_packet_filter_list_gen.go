// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func PacketFilterList(ctx Context, params *ListPacketFilterParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetPacketFilterAPI()

	finder.SetEmpty()

	if !isEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !isEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !isEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}
	if !isEmpty(params.Name) {
		for _, v := range params.Name {
			finder.SetFilterBy("Name", v)
		}
	}
	if !isEmpty(params.Id) {
		for _, v := range params.Id {
			finder.SetFilterMultiBy("ID", v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("PacketFilterList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.PacketFilters {
		list = append(list, &res.PacketFilters[i])
	}

	return ctx.GetOutput().Print(list...)

}
