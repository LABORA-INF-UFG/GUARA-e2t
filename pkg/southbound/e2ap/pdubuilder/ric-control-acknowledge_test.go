// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/LABORA-INF-UFG/GUARA-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	"github.com/LABORA-INF-UFG/GUARA-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestRicControlAcknowledge(t *testing.T) {
	ricRequestID := types.RicRequest{
		RequestorID: 21,
		InstanceID:  22,
	}
	var ranFuncID types.RanFunctionID = 9
	var ricCallPrID types.RicCallProcessID = []byte("123")
	var ricCtrlOut types.RicControlOutcome = []byte("456")
	newE2apPdu, err := CreateRicControlAcknowledgeE2apPdu(ricRequestID,
		ranFuncID)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetSuccessfulOutcome().GetValue().GetRicControl().
		SetRicControlOutcome(ricCtrlOut).SetRicCallProcessID(ricCallPrID)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("RicControlAcknowledge E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}
