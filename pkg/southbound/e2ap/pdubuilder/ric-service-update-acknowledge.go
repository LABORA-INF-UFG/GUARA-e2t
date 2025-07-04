// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"

	"github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2appducontents "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/LABORA-INF-UFG/GUARA-e2t/pkg/southbound/e2ap/types"
)

func CreateRicServiceUpdateAcknowledgeE2apPdu(trID int32, rfAccepted types.RanFunctionRevisions) (*e2appdudescriptions.E2ApPdu, error) {

	if rfAccepted == nil {
		return nil, fmt.Errorf("RanFunctionsAccepetd was not passed - it is mandatory parameter")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDRICserviceUpdate),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
					SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicServiceUpdate{
						RicServiceUpdate: &e2appducontents.RicserviceUpdateAcknowledge{
							ProtocolIes: make([]*e2appducontents.RicserviceUpdateAcknowledgeIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetSuccessfulOutcome().GetValue().GetRicServiceUpdate().SetTransactionID(trID).SetRanFunctionsAccepted(rfAccepted)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
