// Copyright (C) 2020 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpplink

import (
	"fmt"

	"github.com/projectcalico/vpp-dataplane/v3/vpplink/generated/bindings/rdma"
	"github.com/projectcalico/vpp-dataplane/v3/vpplink/types"
)

func (v *VppLink) CreateRDMA(intf *types.RDMAInterface) (swIfIndex uint32, err error) {
	client := rdma.NewServiceClient(v.GetConnection())

	response, err := client.RdmaCreateV2(v.GetContext(), &rdma.RdmaCreateV2{
		HostIf:  intf.HostInterfaceName,
		RxqNum:  uint16(intf.NumRxQueues),
		RxqSize: uint16(intf.RxQueueSize),
		TxqSize: uint16(intf.TxQueueSize),
	})
	if err != nil {
		return 0, fmt.Errorf("failed to create RDMA interface: %w", err)
	}
	intf.SwIfIndex = uint32(response.SwIfIndex)
	return uint32(response.SwIfIndex), nil
}
