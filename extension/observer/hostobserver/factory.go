// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hostobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
)

const (
	// The value of extension "type" in configuration.
	typeStr config.Type = "host_observer"

	defaultCollectionInterval = 10
)

// NewFactory creates a factory for HostObserver extension.
func NewFactory() component.ExtensionFactory {
	return component.NewExtensionFactory(
		typeStr,
		createDefaultConfig,
		createExtension)
}

func createDefaultConfig() config.Extension {
	return &Config{
		ExtensionSettings: config.NewExtensionSettings(config.NewComponentID(typeStr)),
		RefreshInterval:   defaultCollectionInterval * time.Second,
	}
}

func createExtension(
	_ context.Context,
	params component.ExtensionCreateSettings,
	cfg config.Extension,
) (component.Extension, error) {
	config := cfg.(*Config)
	return newObserver(params.Logger, config)
}
