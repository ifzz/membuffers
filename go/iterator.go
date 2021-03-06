// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package membuffers

import "math/big"

type Iterator struct {
	cursor    Offset
	endCursor Offset
	fieldType FieldType
	m         *InternalMessage
}

func (i *Iterator) HasNext() bool {
	return i.cursor < i.endCursor
}

func (i *Iterator) NextBool() bool {
	if i.cursor+FieldSizes[TypeBool] > i.endCursor {
		i.cursor = i.endCursor
		return false
	}
	res := i.m.GetBoolInOffset(i.cursor)
	i.cursor += FieldSizes[TypeBool]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeBoolArray)
	return res
}

func (i *Iterator) NextUint8() uint8 {
	if i.cursor+FieldSizes[TypeUint8] > i.endCursor {
		i.cursor = i.endCursor
		return 0
	}
	res := i.m.GetUint8InOffset(i.cursor)
	i.cursor += FieldSizes[TypeUint8]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeUint8Array)
	return res
}

func (i *Iterator) NextUint16() uint16 {
	if i.cursor+FieldSizes[TypeUint16] > i.endCursor {
		i.cursor = i.endCursor
		return 0
	}
	res := i.m.GetUint16InOffset(i.cursor)
	i.cursor += FieldSizes[TypeUint16]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeUint16Array)
	return res
}

func (i *Iterator) NextUint32() uint32 {
	if i.cursor+FieldSizes[TypeUint32] > i.endCursor {
		i.cursor = i.endCursor
		return 0
	}
	res := i.m.GetUint32InOffset(i.cursor)
	i.cursor += FieldSizes[TypeUint32]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeUint32Array)
	return res
}

func (i *Iterator) NextUint64() uint64 {
	if i.cursor+FieldSizes[TypeUint64] > i.endCursor {
		i.cursor = i.endCursor
		return 0
	}
	res := i.m.GetUint64InOffset(i.cursor)
	i.cursor += FieldSizes[TypeUint64]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeUint64Array)
	return res
}

func (i *Iterator) NextMessage() (buf []byte, size Offset) {
	if i.cursor+FieldSizes[TypeMessage] > i.endCursor {
		i.cursor = i.endCursor
		return []byte{}, 0
	}
	resSize := i.m.GetOffsetInOffset(i.cursor)
	i.cursor += FieldSizes[TypeMessage]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeMessage)
	if i.cursor+resSize > i.endCursor {
		i.cursor = i.endCursor
		return []byte{}, 0
	}
	resBuf := i.m.bytes[i.cursor : i.cursor+resSize]
	i.cursor += resSize
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeMessageArray)
	return resBuf, resSize
}

func (i *Iterator) NextBytes() []byte {
	if i.cursor+FieldSizes[TypeBytes] > i.endCursor {
		i.cursor = i.endCursor
		return []byte{}
	}
	resSize := i.m.GetOffsetInOffset(i.cursor)
	i.cursor += FieldSizes[TypeBytes]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeBytes)
	if i.cursor+resSize > i.endCursor {
		i.cursor = i.endCursor
		return []byte{}
	}
	resBuf := i.m.bytes[i.cursor : i.cursor+resSize]
	i.cursor += resSize
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeBytesArray)
	return resBuf
}

func (i *Iterator) NextBytes20() (out [20]byte) {
	if i.cursor+FieldSizes[TypeBytes20] > i.endCursor {
		i.cursor = i.endCursor
		return
	}

	out = GetBytes20(i.m.bytes[i.cursor:])
	i.cursor += FieldSizes[TypeBytes20]
	return
}

func (i *Iterator) NextBytes32() (out [32]byte) {
	if i.cursor+FieldSizes[TypeBytes32] > i.endCursor {
		i.cursor = i.endCursor
		return
	}

	out = GetBytes32(i.m.bytes[i.cursor:])
	i.cursor += FieldSizes[TypeBytes32]
	return
}

func (i *Iterator) NextUint256() (out *big.Int) {
	if i.cursor+FieldSizes[TypeUint256] > i.endCursor {
		i.cursor = i.endCursor
		return
	}

	out = GetUint256(i.m.bytes[i.cursor:])
	i.cursor += FieldSizes[TypeUint256]
	return
}

func (i *Iterator) NextString() string {
	b := i.NextBytes()
	return byteSliceToString(b)
}
