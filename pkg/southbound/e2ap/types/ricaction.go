// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import e2apies "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-ies"

type RicActionID int32

type RicActionDefinition []byte

type RicEventDefintion []byte

type RicActionDef struct {
	RicActionID         RicActionID
	RicActionType       e2apies.RicactionType
	RicSubsequentAction e2apies.RicsubsequentActionType
	Ricttw              e2apies.RictimeToWait
	RicActionDefinition RicActionDefinition
}
