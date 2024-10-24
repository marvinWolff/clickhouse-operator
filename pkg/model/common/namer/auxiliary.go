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

package namer

import (
	"fmt"

	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
)

// IsAutoGeneratedShardName checks whether provided name is auto-generated
func IsAutoGeneratedShardName(name string, shard api.IShard, index int) bool {
	return name == createShardName(shard, index)
}

// IsAutoGeneratedReplicaName checks whether provided name is auto-generated
func IsAutoGeneratedReplicaName(name string, replica api.IReplica, index int) bool {
	return name == createReplicaName(replica, index)
}

// IsAutoGeneratedHostName checks whether name is auto-generated
func IsAutoGeneratedHostName(
	name string,
	host *api.Host,
	shard api.IShard,
	shardIndex int,
	replica api.IReplica,
	replicaIndex int,
) bool {
	switch {
	case name == createHostName(host, shard, shardIndex, replica, replicaIndex):
		// Current version of the name
		return true
	case name == fmt.Sprintf("%d-%d", shardIndex, replicaIndex):
		// old version - index-index
		return true
	case name == fmt.Sprintf("%d", shardIndex):
		// old version - index
		return true
	case name == fmt.Sprintf("%d", replicaIndex):
		// old version - index
		return true
	default:
		return false
	}
}
