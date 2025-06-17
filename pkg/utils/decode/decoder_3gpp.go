// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package decode

import (
	asn1 "github.com/LABORA-INF-UFG/GUARA-aper/api/asn1/v1/bitstring"
	"math"
)

const ByteBits = 8.0

func Asn1BitstringToUint64(source *asn1.BitString) *uint64 {
	result := uint64(0)
	arrayLen := int(math.Ceil(float64(source.Len) / ByteBits))
	for i := 0; i < arrayLen; i++ {
		result = result + uint64(source.Value[i])<<(i*ByteBits)
	}
	return &result
}

func Asn1BytesToUint64(source []byte) *uint64 {
	result := uint64(0)
	for i := 0; i < len(source); i++ {
		result = result + uint64(source[i])<<(i*ByteBits)
	}
	return &result
}
