// Copyright 2016-2020 The Libsacloud Authors
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

package server

import (
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Zone string `request:"-" validate:"required"`

	Name                 string `validate:"required"`
	Description          string `validate:"min=0,max=512"`
	Tags                 types.Tags
	IconID               types.ID
	CPU                  int
	MemoryMB             int
	ServerPlanCommitment types.ECommitment
	ServerPlanGeneration types.EPlanGeneration
	ConnectedSwitches    []*sacloud.ConnectedSwitch
	InterfaceDriver      types.EInterfaceDriver
	WaitDiskMigration    bool
	PrivateHostID        types.ID
}

func (req *CreateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateRequest) ToRequestParameter() (*sacloud.ServerCreateRequest, error) {
	params := &sacloud.ServerCreateRequest{}
	if err := service.RequestConvertTo(req, params); err != nil {
		return nil, err
	}
	return params, nil
}