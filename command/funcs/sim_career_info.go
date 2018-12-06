package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMCareerInfo(ctx command.Context, params *params.CareerInfoSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMCareerInfo is failed: %s", e)
	}

	careerInfo, err := api.GetNetworkOperator(p.ID)
	if err != nil {
		return fmt.Errorf("SIMCareerInfo is failed: %s", err)
	}

	var list []interface{}
	for _, v := range careerInfo.NetworkOperatorConfigs {
		if v.Allow {
			list = append(list, v)
		}
	}
	return ctx.GetOutput().Print(list...)

}
