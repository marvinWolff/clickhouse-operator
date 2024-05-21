// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
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

package creator

import (
	core "k8s.io/api/core/v1"

	"github.com/altinity/clickhouse-operator/pkg/model/common/creator"
)

func (c *Creator) CreateConfigMap(what creator.ConfigMapType, params ...any) *core.ConfigMap {
	c.cmm.SetCR(c.cr)
	c.cmm.SetTagger(c.tagger)
	c.cmm.SetConfigFilesGenerator(c.configFilesGenerator)
	return c.cmm.CreateConfigMap(what, params...)
}
