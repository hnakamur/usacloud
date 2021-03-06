// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InterfaceCreate(ctx command.Context, params *params.CreateInterfaceParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInterfaceAPI()
	p := api.New()

	// set params

	p.SetServerID(params.ServerId)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("InterfaceCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
