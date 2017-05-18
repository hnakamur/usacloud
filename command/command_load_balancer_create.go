package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command/internal"
)

func LoadBalancerCreate(ctx Context, params *CreateLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()

	// validate
	if params.HighAvailability && params.Ipaddress2 == "" {
		return fmt.Errorf("%q: is required when using high-availability flag", "ipaddress2")
	}

	p := &sacloud.CreateLoadBalancerValue{
		SwitchID:     fmt.Sprintf("%d", params.SwitchId),
		VRID:         params.VRID,
		IPAddress1:   params.Ipaddress1,
		MaskLen:      params.NwMaskLen,
		DefaultRoute: params.DefaultRoute,
		Name:         params.Name,
		Description:  params.Description,
		Tags:         params.Tags,
		Icon:         sacloud.NewResource(params.IconId),
	}

	switch params.Plan {
	case "standard":
		p.Plan = sacloud.LoadBalancerPlanStandard
	case "highspec":
		p.Plan = sacloud.LoadBalancerPlanPremium
	}

	var lb *sacloud.LoadBalancer
	var err error
	if params.HighAvailability {
		//冗長構成
		lb, err = sacloud.CreateNewLoadBalancerDouble(&sacloud.CreateDoubleLoadBalancerValue{
			CreateLoadBalancerValue: p,
			IPAddress2:              params.Ipaddress2,
		}, nil)

		if err != nil {
			return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
		}

	} else {
		lb, err = sacloud.CreateNewLoadBalancerSingle(p, nil)
		if err != nil {
			return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
		}
	}

	// call Create(id)
	res, err := api.Create(lb)
	if err != nil {
		return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
	}

	// wait for boot
	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewProgress(
		fmt.Sprintf("Still creating[ID:%d]...", res.ID),
		fmt.Sprintf("Create LoadBalancer[ID:%d]", res.ID),
		GlobalOption.Progress)

	go func() {
		spinner.Start()
		// call manipurate functions
		err := api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
		if err != nil {
			errChan <- err
			return
		}
		err = api.SleepUntilUp(res.ID, client.DefaultTimeoutDuration)
		if err != nil {
			errChan <- err
			return
		}
		compChan <- true
	}()

boot:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break boot
		case err := <-errChan:
			return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(res)

}