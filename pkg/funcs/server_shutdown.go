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

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func ServerShutdown(ctx cli.Context, params *params.ShutdownServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerShutdown is failed: %s", e)
	}

	if p.IsDown() {
		return nil // already downed.
	}

	err := progress.ExecWithProgress(
		fmt.Sprintf("Still waiting for Shutdown[ID:%d]...", params.Id),
		fmt.Sprintf("Shutdown server[ID:%d]", params.Id),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			var err error
			_, err = api.Shutdown(params.Id)
			if err != nil {
				errChan <- err
				return
			}

			err = api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("ServerShutdown is failed: %s", err)
	}

	return nil
}
