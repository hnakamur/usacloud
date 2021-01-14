// Copyright 2017-2021 The Usacloud Authors
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

package util

import (
	"fmt"
	"io"

	"github.com/itchyny/gojq"
)

func FilterPrintByGoJQ(errW io.Writer, input interface{}, query string) (result interface{}, err error) {
	q, err := gojq.Parse(query)
	if err != nil {
		if errW != nil {
			fmt.Fprintf(errW, "gojq parse failed: %v\n", err)
		}
		return nil, fmt.Errorf("gojq parse failed: %v", err)
	}
	iter := q.Run(input)
	// We process only the first value here.
	v, ok := iter.Next()
	if !ok {
		return nil, nil
	}
	if err, ok := v.(error); ok {
		if errW != nil {
			fmt.Fprintf(errW, "gojq parse failed: %v\n", err)
		}
		return nil, err
	}
	return v, nil
}
