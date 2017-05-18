package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func LoadBalancerReset(ctx Context, params *ResetLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerReset is failed: %s", e)
	}

	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewProgress(
		fmt.Sprintf("Still resetting[ID:%d]...", params.Id),
		fmt.Sprintf("Reset LoadBalancer[ID:%d]", params.Id),
		GlobalOption.Progress)

	go func() {

		spinner.Start()
		// call manipurate functions
		_, err := api.RebootForce(params.Id)
		if err != nil {
			errChan <- err
			return
		}
		err = api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
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
			return fmt.Errorf("LoadBalancerReset is failed: %s", err)
		}
	}

	return nil
}