// Copyright 2020 OpenTelemetry Authors
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

// Skip tests on Windows temporarily, see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/11451
//go:build !windows
// +build !windows

package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/config"
)

func TestDefaultComponents(t *testing.T) {
	factories, err := Components()
	assert.NoError(t, err)

	exts := factories.Extensions
	for k, v := range exts {
		assert.Equal(t, k, v.Type())
		assert.Equal(t, config.NewComponentID(k), v.CreateDefaultConfig().ID())
	}

	recvs := factories.Receivers
	for k, v := range recvs {
		assert.Equal(t, k, v.Type())
		assert.Equal(t, config.NewComponentID(k), v.CreateDefaultConfig().ID())
	}

	procs := factories.Processors
	for k, v := range procs {
		assert.Equal(t, k, v.Type())
		assert.Equal(t, config.NewComponentID(k), v.CreateDefaultConfig().ID())
	}

	exps := factories.Exporters
	for k, v := range exps {
		assert.Equal(t, k, v.Type())
		assert.Equal(t, config.NewComponentID(k), v.CreateDefaultConfig().ID())
	}
}
