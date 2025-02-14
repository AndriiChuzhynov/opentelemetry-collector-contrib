// Copyright  The OpenTelemetry Authors
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

package ottlfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/oteltransformationlanguage/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/oteltransformationlanguage/ottl/ottltest"
)

func Test_split(t *testing.T) {
	tests := []struct {
		name      string
		target    ottl.Getter
		delimiter string
		expected  interface{}
	}{
		{
			name: "split string",
			target: &ottl.StandardGetSetter{
				Getter: func(ctx ottl.TransformContext) interface{} {
					return "A|B|C"
				},
			},
			delimiter: "|",
			expected:  []string{"A", "B", "C"},
		},
		{
			name: "split empty string",
			target: &ottl.StandardGetSetter{
				Getter: func(ctx ottl.TransformContext) interface{} {
					return ""
				},
			},
			delimiter: "|",
			expected:  []string{""},
		},
		{
			name: "split empty delimiter",
			target: &ottl.StandardGetSetter{
				Getter: func(ctx ottl.TransformContext) interface{} {
					return "A|B|C"
				},
			},
			delimiter: "",
			expected:  []string{"A", "|", "B", "|", "C"},
		},
		{
			name: "split empty string and empty delimiter",
			target: &ottl.StandardGetSetter{
				Getter: func(ctx ottl.TransformContext) interface{} {
					return ""
				},
			},
			delimiter: "",
			expected:  []string{},
		},
		{
			name: "split non-string",
			target: &ottl.StandardGetSetter{
				Getter: func(ctx ottl.TransformContext) interface{} {
					return 123
				},
			},
			delimiter: "|",
			expected:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := ottltest.TestTransformContext{}

			exprFunc, _ := Split(tt.target, tt.delimiter)
			actual := exprFunc(ctx)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
