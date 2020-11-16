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

package vdef

import (
	"fmt"
	"strings"
)

var validatorAliases = map[string]string{
	"description": "max=512",
	"tags":        "max=10,dive,max=32",
	//"disk_plan":       fmt.Sprintf("oneof=%s", joinWithSpace(types.DiskPlanStrings)),
	//"disk_connection": fmt.Sprintf("oneof=%s", joinWithSpace(types.DiskConnectionStrings)),
	//"os_type":         fmt.Sprintf("oneof=%s", joinWithSpace(ostype.OSTypeShortNames)),
	// "zone": ... // NOTE: 実行時に登録される
}

func ValidatorAliases(zones []string) map[string]string {
	aliases := validatorAliases
	aliases["zone"] = fmt.Sprintf("required,oneof=%s", joinWithSpace(zones))
	return aliases
}

func joinWithSpace(values []string) string {
	return strings.Join(values, " ")
}
