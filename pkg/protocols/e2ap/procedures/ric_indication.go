// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package procedures

import (
	"context"

	"github.com/LABORA-INF-UFG/GUARA-utils/pkg/errors"

	e2api "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2"

	"github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2appducontents "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
)

// RICIndication is a RIC indication procedure
type RICIndication interface {
	RICIndication(ctx context.Context, request *e2appducontents.Ricindication) (err error)
}

func NewRICIndicationInitiator(dispatcher Dispatcher) *RICIndicationInitiator {
	return &RICIndicationInitiator{
		dispatcher: dispatcher,
	}
}

type RICIndicationInitiator struct {
	dispatcher Dispatcher
}

func (p *RICIndicationInitiator) Initiate(_ context.Context, request *e2appducontents.Ricindication) (err error) {
	pdu := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(e2api.ProcedureCodeIDRICindication),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicIndication{
						RicIndication: request,
					},
				},
			},
		},
	}
	if err := pdu.Validate(); err != nil {
		return errors.NewInvalid("E2AP PDU validation failed: %v", err)
	}
	return p.dispatcher(pdu)
}

func (p *RICIndicationInitiator) Matches(_ *e2appdudescriptions.E2ApPdu) bool {
	return false
}

func (p *RICIndicationInitiator) Handle(_ *e2appdudescriptions.E2ApPdu) {

}

func (p *RICIndicationInitiator) Close() error {
	return nil
}

var _ ElementaryProcedure = &RICIndicationInitiator{}

func NewRICIndicationProcedure(dispatcher Dispatcher, handler RICIndication) *RICIndicationProcedure {
	return &RICIndicationProcedure{
		dispatcher: dispatcher,
		handler:    handler,
	}
}

type RICIndicationProcedure struct {
	dispatcher Dispatcher
	handler    RICIndication
}

func (p *RICIndicationProcedure) Matches(pdu *e2appdudescriptions.E2ApPdu) bool {
	switch msg := pdu.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_InitiatingMessage:
		switch ret := msg.InitiatingMessage.Value.ImValues.(type) {
		case *e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicIndication:
			return ret.RicIndication != nil
		default:
			return false
		}
	default:
		return false
	}
}

func (p *RICIndicationProcedure) Handle(pdu *e2appdudescriptions.E2ApPdu) {
	err := p.handler.RICIndication(context.Background(), pdu.GetInitiatingMessage().GetValue().GetRicIndication())
	if err != nil {
		log.Errorf("RIC Indication procedure failed: %v", err)
	}
}

func (p *RICIndicationProcedure) Close() error {
	return nil
}

var _ ElementaryProcedure = &RICIndicationProcedure{}
