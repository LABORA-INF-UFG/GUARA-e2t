// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/LABORA-INF-UFG/GUARA-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	e2ap_ies "github.com/LABORA-INF-UFG/GUARA-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/LABORA-INF-UFG/GUARA-e2t/pkg/southbound/e2ap/types"
	asn1 "github.com/LABORA-INF-UFG/GUARA-utils/api/asn1/v1/bitstring"
	"gotest.tools/assert"
)

func TestE2SetupRequest(t *testing.T) {
	e2ncID1 := CreateE2NodeComponentIDW1(21)
	e2ncID2 := CreateE2NodeComponentIDS1("S1-Component")
	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         "oid2",
	}

	ge2nID, err := CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	})
	assert.NilError(t, err)

	newE2apPdu, err := CreateE2SetupRequestPdu(1, ge2nID, ranFunctionList, []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
			E2NodeComponentID: e2ncID1,
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
				E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
			}},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
			E2NodeComponentID: e2ncID2,
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
				E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
			}},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdate E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())
}

func TestE2SetupRequestCuDuIDs(t *testing.T) {
	enbID, err := CreateEnbIDLongMacro(&asn1.BitString{
		Value: []byte{0x00, 0xA7, 0xF8},
		Len:   21,
	})
	assert.NilError(t, err)
	gEnbID, err := CreateGlobalEnbID([]byte{0xAA, 0xBB, 0xCC}, enbID)
	assert.NilError(t, err)

	gEnGnbID, err := CreateGlobalEnGnbID([]byte{0xFF, 0xCD, 0xBF}, &asn1.BitString{
		Value: []byte{0xFA, 0x2C, 0xD4, 0xF8},
		Len:   29,
	})
	assert.NilError(t, err)

	e2ncID1 := CreateE2NodeComponentIDX2(gEnbID, gEnGnbID)
	e2ncID2 := CreateE2NodeComponentIDNg("NG-Component")

	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         "oid2",
	}

	ge2nID, err := CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	})
	assert.NilError(t, err)
	ge2nID.GetGNb().SetGnbCuUpID(2).SetGnbDuID(13)

	newE2apPdu, err := CreateE2SetupRequestPdu(1, ge2nID, ranFunctionList, []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_X2,
			E2NodeComponentID: e2ncID1,
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x01, 0x02, 0x03},
				E2NodeComponentRequestPart:  []byte{0x04, 0x05, 0x06},
			}},
		{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
			E2NodeComponentID: e2ncID2,
			E2NodeComponentConfiguration: e2ap_ies.E2NodeComponentConfiguration{
				E2NodeComponentResponsePart: []byte{0x07, 0x08, 0x09},
				E2NodeComponentRequestPart:  []byte{0x0A, 0x0B, 0x0C},
			}},
	})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	e2apPdu, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), e2apPdu.String())
}

var rSysBites = []byte{
	0x00, 0x01, 0x00, 0x85, 0xd2, 0x00, 0x00, 0x04, 0x00, 0x31, 0x00, 0x02, 0x00, 0x01, 0x00, 0x03,
	0x00, 0x08, 0x00, 0x13, 0xf1, 0x84, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x84, 0x10, 0x01,
	0x00, 0x08, 0x40, 0x81, 0x7d, 0x00, 0x00, 0x01, 0x81, 0x5a, 0x70, 0x18, 0x4f, 0x52, 0x41, 0x4e,
	0x2d, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x4b, 0x50, 0x4d, 0x00, 0x00, 0x18, 0x31, 0x2e, 0x33, 0x2e,
	0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31, 0x2e,
	0x32, 0x2e, 0x32, 0x2e, 0x32, 0x05, 0x00, 0x4b, 0x50, 0x4d, 0x20, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x00, 0x00, 0x40, 0x00, 0x13, 0xf1, 0x84, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x30, 0x00, 0x13, 0xf1, 0x84, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x01, 0x07, 0x00,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x00,
	0x01, 0x00, 0x03, 0x09, 0x00, 0x45, 0x32, 0x20, 0x4e, 0x6f, 0x64, 0x65, 0x20, 0x4d, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x00, 0x00, 0x00, 0x07, 0x42, 0x60, 0x52, 0x52,
	0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x73,
	0x75, 0x6d, 0x00, 0x00, 0x00, 0x42, 0x80, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x45,
	0x73, 0x74, 0x61, 0x62, 0x53, 0x75, 0x63, 0x63, 0x2e, 0x73, 0x75, 0x6d, 0x00, 0x00, 0x01, 0x41,
	0xa0, 0x52, 0x52, 0x43, 0x2e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x00,
	0x00, 0x02, 0x44, 0x80, 0x52, 0x52, 0x43, 0x2e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41,
	0x74, 0x74, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x00, 0x00, 0x03, 0x43, 0xa0, 0x52, 0x52,
	0x43, 0x2e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e, 0x68, 0x61, 0x6e,
	0x64, 0x6f, 0x76, 0x65, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x00, 0x00, 0x04, 0x43,
	0x40, 0x52, 0x52, 0x43, 0x2e, 0x52, 0x65, 0x45, 0x73, 0x74, 0x61, 0x62, 0x41, 0x74, 0x74, 0x2e,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x00, 0x00, 0x05, 0x41,
	0x60, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x65, 0x61, 0x6e, 0x00, 0x00, 0x06,
	0x41, 0x40, 0x52, 0x52, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x00, 0x00, 0x07,
	0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x18, 0x31, 0x2e, 0x33, 0x2e, 0x36, 0x2e, 0x31,
	0x2e, 0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31, 0x2e, 0x32, 0x2e, 0x32,
	0x2e, 0x32, 0x00, 0x08, 0x40, 0x82, 0x88, 0x00, 0x00, 0x02, 0x82, 0x65, 0x78, 0x05, 0x80, 0x4f,
	0x52, 0x41, 0x4e, 0x2d, 0x45, 0x32, 0x53, 0x4d, 0x2d, 0x52, 0x43, 0x00, 0x00, 0x18, 0x31, 0x2e,
	0x33, 0x2e, 0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e, 0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e,
	0x31, 0x2e, 0x31, 0x2e, 0x32, 0x2e, 0x33, 0x00, 0x80, 0x52, 0x43, 0x20, 0x00, 0x00, 0x02, 0x0b,
	0x00, 0x43, 0x61, 0x6c, 0x6c, 0x20, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x20, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x05, 0x0a,
	0x80, 0x50, 0x44, 0x55, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x00, 0x02, 0x40, 0x00, 0x00, 0x0c, 0x80, 0x50, 0x44,
	0x55, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x20, 0x53, 0x65, 0x74, 0x75, 0x70, 0x00, 0x02, 0x90, 0x75, 0x32, 0x0e, 0x00, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x51, 0x6f, 0x53, 0x20, 0x46, 0x6c, 0x6f, 0x77, 0x73,
	0x20, 0x74, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x73, 0x65, 0x74, 0x75, 0x70, 0x00, 0x90, 0x75, 0x33,
	0x0d, 0x00, 0x51, 0x6f, 0x53, 0x20, 0x66, 0x6c, 0x6f, 0x77, 0x20, 0x73, 0x65, 0x74, 0x75, 0x70,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x00, 0x90, 0x75,
	0x35, 0x03, 0x80, 0x51, 0x6f, 0x53, 0x20, 0x66, 0x6c, 0x6f, 0x77, 0x00, 0x40, 0x00, 0x01, 0x10,
	0x00, 0x50, 0x44, 0x55, 0x20, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x00, 0x02, 0x90, 0x79, 0x1a, 0x10, 0x80, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x51, 0x6f, 0x53, 0x20, 0x46, 0x6c, 0x6f, 0x77, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x64,
	0x64, 0x20, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x00, 0x90, 0x79, 0x1b, 0x11,
	0x00, 0x51, 0x6f, 0x53, 0x20, 0x66, 0x6c, 0x6f, 0x77, 0x20, 0x61, 0x64, 0x64, 0x20, 0x6f, 0x72,
	0x20, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20,
	0x69, 0x74, 0x65, 0x6d, 0x00, 0x90, 0x79, 0x1d, 0x03, 0x80, 0x51, 0x6f, 0x53, 0x20, 0x66, 0x6c,
	0x6f, 0x77, 0x00, 0x00, 0x00, 0x02, 0x0d, 0x80, 0x50, 0x44, 0x55, 0x20, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x20, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x00, 0x80, 0x00, 0x02, 0x0b, 0x00, 0x43, 0x61, 0x6c, 0x6c, 0x20, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x20, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x00, 0x02, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02, 0x00, 0x00, 0x80, 0x00, 0x05, 0x80, 0x44,
	0x52, 0x42, 0x20, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x00, 0x00, 0x80, 0x00, 0x01,
	0x0d, 0x80, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x20, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x00, 0x02,
	0x00, 0x03, 0x00, 0x00, 0x40, 0x00, 0x03, 0x0f, 0x00, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x20, 0x61,
	0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x00, 0x00, 0x80, 0x00, 0x05, 0x80, 0x44, 0x52,
	0x42, 0x20, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x00, 0x00, 0x02, 0x00, 0x05, 0x00,
	0x01, 0x00, 0xc0, 0x00, 0x01, 0x09, 0x80, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x20, 0x42, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x20, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x00, 0x00, 0x00, 0x00, 0x03,
	0x0b, 0x00, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x20, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00,
	0x01, 0x00, 0x00, 0x00, 0x00, 0x18, 0x31, 0x2e, 0x33, 0x2e, 0x36, 0x2e, 0x31, 0x2e, 0x34, 0x2e,
	0x31, 0x2e, 0x35, 0x33, 0x31, 0x34, 0x38, 0x2e, 0x31, 0x2e, 0x31, 0x2e, 0x32, 0x2e, 0x33, 0x00,
	0x32, 0x00, 0x81, 0xa3, 0x00, 0x01, 0x00, 0x33, 0x40, 0x81, 0x06, 0x19, 0xa0, 0x01, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x7d, 0x00, 0x01, 0x00, 0x79, 0x00, 0x00, 0x03, 0x00, 0x4e, 0x00, 0x02, 0x00,
	0x01, 0x00, 0x2a, 0x00, 0x06, 0x80, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x2c, 0x00, 0x62, 0x00,
	0x00, 0x00, 0x2b, 0x00, 0x5c, 0x4a, 0x00, 0x13, 0xf1, 0x84, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00,
	0x01, 0x00, 0x00, 0x01, 0x08, 0x13, 0xf1, 0x84, 0x00, 0x00, 0x00, 0x83, 0x00, 0x07, 0x00, 0x00,
	0x40, 0x60, 0x03, 0x06, 0x09, 0x01, 0x00, 0x0a, 0x2c, 0x2a, 0x01, 0x01, 0x00, 0x00, 0x40, 0x0a,
	0x2c, 0x2a, 0x00, 0x01, 0x00, 0x00, 0x65, 0x0c, 0xa0, 0x0a, 0x10, 0x31, 0x5d, 0xef, 0x41, 0x04,
	0x80, 0x80, 0x80, 0x00, 0x00, 0x00, 0x00, 0x8b, 0x40, 0x01, 0x01, 0x00, 0x03, 0x00, 0x00, 0x00,
	0x10, 0x80, 0x0c, 0x8a, 0x37, 0x01, 0x04, 0xc4, 0x49, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x03, 0x7f, 0x40, 0x01, 0x00, 0x7b, 0x00, 0x00, 0x03, 0x00, 0x4e, 0x00, 0x02, 0x00, 0x01, 0x00,
	0x03, 0x00, 0x60, 0x00, 0x00, 0x00, 0x04, 0x00, 0x5a, 0x20, 0x13, 0xf1, 0x84, 0x00, 0x00, 0x00,
	0x00, 0x10, 0x00, 0x00, 0x00, 0x76, 0x00, 0x4b, 0x04, 0x00, 0x20, 0x3c, 0x38, 0x50, 0xa1, 0x6e,
	0x24, 0xa2, 0x03, 0xc8, 0x92, 0x9b, 0xdf, 0xfe, 0x15, 0xb6, 0x85, 0x01, 0x1a, 0x10, 0x46, 0x85,
	0x18, 0x28, 0xf8, 0x04, 0xd2, 0x43, 0xec, 0x6c, 0x14, 0x01, 0xb0, 0x00, 0x04, 0x0e, 0x60, 0x70,
	0x1d, 0xe0, 0x05, 0x02, 0x40, 0x4e, 0x0a, 0x80, 0x19, 0x03, 0xe0, 0x80, 0x00, 0x0c, 0x13, 0x60,
	0x79, 0x2f, 0xe8, 0xa0, 0x90, 0x27, 0x0a, 0x80, 0x32, 0x0f, 0x84, 0x17, 0xea, 0x1c, 0x02, 0x47,
	0x40, 0x04, 0x00, 0x00, 0xaa, 0x00, 0x0a, 0x80, 0x00, 0x00, 0x00, 0xc7, 0x40, 0x03, 0x11, 0x01,
	0x00, 0x00, 0x33, 0x40, 0x80, 0x91, 0x00, 0x01, 0x20, 0x52, 0x61, 0x64, 0x69, 0x73, 0x79, 0x73,
	0x41, 0x4d, 0x46, 0x00, 0x2c, 0x00, 0x15, 0x00, 0x28, 0x00, 0x00, 0x03, 0x00, 0x1b, 0x00, 0x08,
	0x00, 0x13, 0xf1, 0x84, 0x00, 0x00, 0x00, 0x00, 0x00, 0x66, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00,
	0x01, 0x00, 0x13, 0xf1, 0x84, 0x00, 0x00, 0x10, 0x18, 0x03, 0x06, 0x09, 0x00, 0x15, 0x40, 0x01,
	0x00, 0x55, 0x20, 0x15, 0x00, 0x51, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x0c, 0x04, 0x80, 0x52,
	0x61, 0x64, 0x69, 0x73, 0x79, 0x73, 0x41, 0x4d, 0x46, 0x00, 0x60, 0x00, 0x12, 0x00, 0x40, 0x13,
	0xf1, 0x84, 0x09, 0x01, 0x83, 0x03, 0x80, 0x52, 0x53, 0x59, 0x53, 0x41, 0x4d, 0x46, 0x31, 0x00,
	0x56, 0x40, 0x01, 0xff, 0x00, 0x50, 0x00, 0x1f, 0x00, 0x13, 0xf1, 0x84, 0x00, 0x04, 0x10, 0x18,
	0x03, 0x06, 0x09, 0x10, 0x18, 0x03, 0x06, 0x0a, 0x10, 0x08, 0x00, 0x00, 0x03, 0x10, 0x08, 0x00,
	0x00, 0x04, 0x10, 0x08, 0x00, 0x00, 0x05}

func TestDecodeE2SetupRequestRadisys(t *testing.T) {
	e2apPdu, err := encoder.PerDecodeE2ApPdu(rSysBites)
	assert.NilError(t, err)
	t.Logf("Decoded message is:\n%v\n", e2apPdu.String())
}
