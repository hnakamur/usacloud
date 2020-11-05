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

package config

import (
	"errors"

	"github.com/fatih/color"
	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Use(ctx cli.Context, params *params.UseConfigParam) error {
	args := ctx.Args()
	if len(args) == 0 || args[0] == "" {
		return errors.New("profile name required")
	}

	profileName := args[0]
	if err := profile.SetCurrentName(profileName); err != nil {
		return err
	}

	color.New(color.FgHiGreen).Fprintf(ctx.IO().Out(), "\nCurrent profile: %q\n", profileName)
	return nil
}