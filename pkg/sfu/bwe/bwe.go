// Copyright 2023 LiveKit, Inc.
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

package bwe

import (
	"fmt"

	"github.com/livekit/livekit-server/pkg/sfu/ccutils"
	"github.com/pion/rtcp"
)

// ------------------------------------------------

type CongestionState int

const (
	CongestionStateNone CongestionState = iota
	CongestionStateEarlyWarning
	CongestionStateEarlyWarningHangover
	CongestionStateCongested
	CongestionStateCongestedHangover
)

func (c CongestionState) String() string {
	switch c {
	case CongestionStateNone:
		return "NONE"
	case CongestionStateEarlyWarning:
		return "EARLY_WARNING"
	case CongestionStateEarlyWarningHangover:
		return "EARLY_WARNING_HANGOVER"
	case CongestionStateCongested:
		return "CONGESTED"
	case CongestionStateCongestedHangover:
		return "CONGESTED_HANGOVER"
	default:
		return fmt.Sprintf("%d", int(c))
	}
}

// ------------------------------------------------

type ProbeSignal int

const (
	ProbeSignalInconclusive ProbeSignal = iota
	ProbeSignalCongesting
	ProbeSignalClearing
)

func (p ProbeSignal) String() string {
	switch p {
	case ProbeSignalInconclusive:
		return "INCONCLUSIVE"
	case ProbeSignalCongesting:
		return "CONGESTING"
	case ProbeSignalClearing:
		return "CLEARING"
	default:
		return fmt.Sprintf("%d", int(p))
	}
}

// ------------------------------------------------

type BWE interface {
	SetBWEListener(bweListner BWEListener)

	Reset()

	Stop()

	HandleREMB(
		receivedEstimate int64,
		expectedBandwidthUsage int64,
		sentPackets uint32,
		repeatedNacks uint32,
	)

	// TWCC sequence number
	RecordPacketSendAndGetSequenceNumber(atMicro int64, size int, isRTX bool) uint16

	HandleTWCCFeedback(report *rtcp.TransportLayerCC)

	CongestionState() CongestionState

	ProbeClusterStarting(pci ccutils.ProbeClusterInfo)
	ProbeClusterDone(pci ccutils.ProbeClusterInfo) (ProbeSignal, int64)
}

// ------------------------------------------------

type BWEListener interface {
	OnCongestionStateChange(congestionState CongestionState, estimatedAvailableChannelCapacity int64)
}

// ------------------------------------------------
