// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	asn1 "github.com/LABORA-INF-UFG/GUARA-aper/api/asn1/v1/bitstring"
)

type TnlInformation struct {
	TnlPort    *asn1.BitString // optional structure
	TnlAddress asn1.BitString
}

type TnlAssociationRemovalItem struct {
	TnlInformation    TnlInformation
	TnlInformationRic TnlInformation
}
