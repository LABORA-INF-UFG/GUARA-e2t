// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	v2 "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/LABORA-INF-UFG/GUARA-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/LABORA-INF-UFG/GUARA-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeE2connectionUpdateFailurePdu(t *testing.T) {
	procCode := v2.ProcedureCodeIDRICsubscription
	criticality := e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE
	ftg := e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME

	e2apPdu, err := pdubuilder.CreateE2connectionUpdateFailureE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetUnsuccessfulOutcome().GetValue().GetE2ConnectionUpdate().
		SetCause(&e2ap_ies.Cause{
			Cause: &e2ap_ies.Cause_Protocol{
				Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
			},
		}).SetTimeToWait(e2ap_ies.TimeToWait_TIME_TO_WAIT_V5S).SetCriticalityDiagnostics(&procCode, &criticality, &ftg,
		&types.RicRequest{
			RequestorID: 10,
			InstanceID:  20,
		}, []*types.CritDiag{
			{
				TypeOfError:   e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING,
				IECriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				IEId:          v2.ProtocolIeIDRicsubscriptionDetails,
			},
		})

	transactionID, cause, ttw, pr, crit, tm, cdrID, diags, err := DecodeE2connectionUpdateFailurePdu(e2apPdu)
	assert.NilError(t, err)

	assert.Equal(t, int32(e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR), int32(cause.GetProtocol()))
	assert.Equal(t, int32(*ttw), int32(e2ap_ies.TimeToWait_TIME_TO_WAIT_V5S))
	assert.Equal(t, int32(*pr), int32(8))
	assert.Equal(t, int32(*crit), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
	assert.Equal(t, int32(*tm), int32(e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFUL_OUTCOME))
	assert.Equal(t, int32(cdrID.InstanceID), int32(20))
	assert.Equal(t, int32(cdrID.RequestorID), int32(10))
	assert.Equal(t, int32(diags[0].IEId), int32(30))
	assert.Equal(t, int32(diags[0].IECriticality), int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE))
	assert.Equal(t, int32(diags[0].TypeOfError), int32(e2ap_ies.TypeOfError_TYPE_OF_ERROR_MISSING))
	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
}
