// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vppcalls

import (
	govppapi "git.fd.io/govpp.git/api"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/bin_api/l2"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/bin_api/vpe"
)

// CheckMsgCompatibilityForBridgeDomains checks if CRSs are compatible with VPP in runtime.
func CheckMsgCompatibilityForBridgeDomains(log logging.Logger, vppChan *govppapi.Channel) error {
	msgs := []govppapi.Message{
		&l2.BridgeDomainAddDel{},
		&l2.BridgeDomainAddDelReply{},
		&l2.L2fibAddDel{},
		&l2.L2fibAddDelReply{},
		&vpe.BdIPMacAddDel{},
		&vpe.BdIPMacAddDelReply{},
		&vpe.SwInterfaceSetL2Bridge{},
		&vpe.SwInterfaceSetL2BridgeReply{},
	}
	err := vppChan.CheckMessageCompatibility(msgs...)
	if err != nil {
		log.Error(err)
	}
	return err
}

// CheckMsgCompatibilityForL2FIB checks if CRSs are compatible with VPP in runtime.
func CheckMsgCompatibilityForL2FIB(log logging.Logger, vppChan *govppapi.Channel) error {
	msgs := []govppapi.Message{
		&l2.BridgeDomainDump{},
		&l2.BridgeDomainDetails{},
		&l2.L2FibTableDump{},
		&l2.L2FibTableDetails{},
		&l2.L2fibAddDel{},
		&l2.L2fibAddDelReply{},
	}
	err := vppChan.CheckMessageCompatibility(msgs...)
	if err != nil {
		log.Error(err)
	}
	return err
}

// CheckMsgCompatibilityForL2XConnect checks if CRSs are compatible with VPP in runtime.
func CheckMsgCompatibilityForL2XConnect(log logging.Logger, vppChan *govppapi.Channel) error {
	msgs := []govppapi.Message{
		&l2.L2XconnectDump{},
		&l2.L2XconnectDetails{},
		&vpe.SwInterfaceSetL2Xconnect{},
		&vpe.SwInterfaceSetL2XconnectReply{},
	}
	err := vppChan.CheckMessageCompatibility(msgs...)
	if err != nil {
		log.Error(err)
	}
	return err
}
