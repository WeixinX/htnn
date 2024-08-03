// Copyright The HTNN Authors.
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

package sentinel

import (
	"mosn.io/htnn/api/pkg/filtermanager/api"
	"mosn.io/htnn/api/pkg/plugins"
	"mosn.io/htnn/plugins/plugins/sentinel/rules"
	"mosn.io/htnn/types/plugins/sentinel"

	sentinelApi "github.com/alibaba/sentinel-golang/api"
	sentinelConf "github.com/alibaba/sentinel-golang/core/config"
)

func init() {
	plugins.RegisterPlugin(sentinel.Name, &plugin{})
}

type plugin struct {
	sentinel.Plugin
}

func (p *plugin) Factory() api.FilterFactory {
	return factory
}

func (p *plugin) Config() api.PluginConfig {
	return &config{}
}

type config struct {
	sentinel.Config
}

func (conf *config) Init(cb api.ConfigCallbackHandler) error {
	sconf := sentinelConf.NewDefaultConfig()
	err := sentinelApi.InitWithConfig(sconf)
	if err != nil {
		return err
	}

	_, err = rules.Load(conf.Type, conf.Rule)
	if err != nil {
		return err
	}
	return nil
}
