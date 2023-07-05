// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bean

import (
	"context"

	"github.com/secretflow/kuscia/pkg/datamesh/config"
	"github.com/secretflow/kuscia/pkg/utils/nlog"
	"github.com/secretflow/kuscia/pkg/web/errorcode"
	"github.com/secretflow/kuscia/pkg/web/framework"
	frameworkconfig "github.com/secretflow/kuscia/pkg/web/framework/config"
)

// operatorBean
// 1 register default datasource
// 2 check the status of domain data periodically
type operatorBean struct {
	frameworkconfig.FlagEnvConfigLoader
	config config.DataMeshConfig
}

func NewOperatorBean(config *config.DataMeshConfig) *operatorBean { // nolint: golint
	return &operatorBean{
		FlagEnvConfigLoader: frameworkconfig.FlagEnvConfigLoader{
			EnableTLSFlag: true,
		},
		config: *config,
	}
}

func (s *operatorBean) Validate(errs *errorcode.Errs) {

}

func (s *operatorBean) Init(e framework.ConfBeanRegistry) error {
	// tls config from config file
	nlog.Infof("operatorBean init")
	return nil
}

// Start grpcServerBean
func (s *operatorBean) Start(ctx context.Context, e framework.ConfBeanRegistry) (err error) {

	return
}

func (s *operatorBean) ServerName() string {
	return "DataMeshOperator"
}