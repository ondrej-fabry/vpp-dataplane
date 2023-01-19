// Copyright (C) 2019 Cisco Systems Inc.
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

	"github.com/pkg/errors"
	"github.com/projectcalico/vpp-dataplane/vpplink/generated/bindings/ipsec"
)

func (v *VppLink) SetIPsecAsyncMode(enable bool) error {
	v.Lock()
	defer v.Unlock()

	response := &ipsec.IpsecSetAsyncModeReply{}

	request := &ipsec.IpsecSetAsyncMode{
		AsyncEnable: enable,
	}
	var err = v.GetChannel().SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrap(err, "IPsec async mode enable failed")
	} else if response.Retval != 0 {
		return fmt.Errorf("IPsec async mode enable failed with retval: %d", response.Retval)
	}
	return nil
}
