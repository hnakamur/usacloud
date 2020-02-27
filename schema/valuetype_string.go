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

// Code generated by "stringer -type=ValueType value_type.go"; DO NOT EDIT.

package schema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TypeInvalid-0]
	_ = x[TypeBool-1]
	_ = x[TypeInt-2]
	_ = x[TypeInt64-3]
	_ = x[TypeFloat-4]
	_ = x[TypeString-5]
	_ = x[TypeIntList-6]
	_ = x[TypeStringList-7]
	_ = x[TypeId-8]
}

const _ValueType_name = "TypeInvalidTypeBoolTypeIntTypeInt64TypeFloatTypeStringTypeIntListTypeStringListTypeId"

var _ValueType_index = [...]uint8{0, 11, 19, 26, 35, 44, 54, 65, 79, 85}

func (i ValueType) String() string {
	if i < 0 || i >= ValueType(len(_ValueType_index)-1) {
		return "ValueType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ValueType_name[_ValueType_index[i]:_ValueType_index[i+1]]
}
