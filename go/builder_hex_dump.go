// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package membuffers

import (
	"fmt"
	"math/big"
)

const HEX_DUMP_INDENT = "    "

func hexDumpAlignOffsetToType(prefix string, offsetFromStart Offset, off Offset, fieldType FieldType) Offset {
	newOff := alignOffsetToType(off, fieldType)
	if newOff > off {
		buf := make([]byte, newOff-off)
		fmt.Printf("%s%x // padding (offset 0x%x, size: 0x%x)\n", prefix, buf, offsetFromStart+off, len(buf))
	}
	return newOff
}

func hexDumpAlignDynamicFieldContentOffset(prefix string, offsetFromStart Offset, off Offset, fieldType FieldType) Offset {
	newOff := alignDynamicFieldContentOffset(off, fieldType)
	if newOff > off {
		buf := make([]byte, newOff-off)
		fmt.Printf("%s%x // padding (offset 0x%x, size: 0x%x)\n", prefix, buf, offsetFromStart+off, len(buf))
	}
	return newOff
}

func HexDumpRawInLines(raw []byte, lineLen int) {
	lineEnd := lineLen
	for lineEnd < len(raw) {
		fmt.Printf("%x\n", raw[lineEnd-lineLen:lineEnd])
		lineEnd += lineLen
	}
	fmt.Printf("%x\n", raw[lineEnd-lineLen:])
	fmt.Println()
	fmt.Printf("(total size: 0x%x)\n", len(raw))
}

func (w *InternalBuilder) HexDumpBool(prefix string, offsetFromStart Offset, fieldName string, v bool) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBool)
	fmt.Printf("%s%v // %s: bool (offset 0x%x, size: 0x%x)\n", prefix, v, fieldName, offsetFromStart+w.size, 1)
	w.size += FieldSizes[TypeBool]
}

func (w *InternalBuilder) HexDumpUint8(prefix string, offsetFromStart Offset, fieldName string, v uint8) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint8)
	buf := make([]byte, FieldSizes[TypeUint8])
	WriteUint8(buf, v)
	fmt.Printf("%s%x // %s: uint8 (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint8]
}

func (w *InternalBuilder) HexDumpUint16(prefix string, offsetFromStart Offset, fieldName string, v uint16) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint16)
	buf := make([]byte, FieldSizes[TypeUint16])
	WriteUint16(buf, v)
	fmt.Printf("%s%x // %s: uint16 (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint16]
}

func (w *InternalBuilder) HexDumpUint32(prefix string, offsetFromStart Offset, fieldName string, v uint32) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint32)
	buf := make([]byte, FieldSizes[TypeUint32])
	WriteUint32(buf, v)
	fmt.Printf("%s%x // %s: uint32 (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint32]
}

func (w *InternalBuilder) HexDumpUint64(prefix string, offsetFromStart Offset, fieldName string, v uint64) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint64)
	buf := make([]byte, FieldSizes[TypeUint64])
	WriteUint64(buf, v)
	fmt.Printf("%s%x // %s: uint64 (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint64]
}

func (w *InternalBuilder) HexDumpBytes(prefix string, offsetFromStart Offset, fieldName string, v []byte) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytes)
	buf := make([]byte, FieldSizes[TypeBytes])
	WriteOffset(buf, Offset(len(v)))
	fmt.Printf("%s%x // %s: bytes size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeBytes]
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytes)
	if v != nil {
		fmt.Printf("%s%x // %s: bytes content (offset 0x%x, size: 0x%x)\n", prefix+HEX_DUMP_INDENT, v, fieldName, offsetFromStart+w.size, len(v))
		w.size += Offset(len(v))
	}
}

func (w *InternalBuilder) HexDumpBytes20(prefix string, offsetFromStart Offset, fieldName string, v [20]byte) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytes20)
	fmt.Printf("%s%x // %s: bytes20 content (offset 0x%x, size: 0x%x)\n", prefix, v, fieldName, offsetFromStart+w.size, len(v))
	w.size += Offset(len(v))
}

func (w *InternalBuilder) HexDumpBytes32(prefix string, offsetFromStart Offset, fieldName string, v [32]byte) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytes32)
	fmt.Printf("%s%x // %s: bytes32 content (offset 0x%x, size: 0x%x)\n", prefix, v, fieldName, offsetFromStart+w.size, len(v))
	w.size += Offset(len(v))
}

func (w *InternalBuilder) HexDumpUint256(prefix string, offsetFromStart Offset, fieldName string, v *big.Int) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint256)
	buf := make([]byte, FieldSizes[TypeUint256])
	WriteUint256(buf, v)
	fmt.Printf("%s%x // %s: uint256 content (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, FieldSizes[TypeUint256])
	w.size += FieldSizes[TypeUint256]
}

func (w *InternalBuilder) HexDumpString(prefix string, offsetFromStart Offset, fieldName string, v string) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytes)
	buf := make([]byte, FieldSizes[TypeBytes])
	WriteOffset(buf, Offset(len(v)))
	fmt.Printf("%s%x // %s: string size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeBytes]
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytes)
	if v != "" {
		fmt.Printf("%s%x // %s: string content (offset 0x%x, size: 0x%x)\n", prefix+HEX_DUMP_INDENT, []byte(v), fieldName, offsetFromStart+w.size, len(v))
		w.size += Offset(len(v))
	}
}

func (w *InternalBuilder) HexDumpUnionIndex(prefix string, offsetFromStart Offset, fieldName string, unionIndex uint16) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUnion)
	buf := make([]byte, FieldSizes[TypeUnion])
	WriteUnionType(buf, unionIndex)
	fmt.Printf("%s%x // %s: union type (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUnion]
}

func (w *InternalBuilder) HexDumpBoolArray(prefix string, offsetFromStart Offset, fieldName string, v []bool) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBoolArray)
	buf := make([]byte, FieldSizes[TypeBoolArray])
	WriteOffset(buf, Offset(len(v))*FieldSizes[TypeBool])
	fmt.Printf("%s%x // %s: bool array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeBoolArray]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeBoolArray)
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpBool(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpUint8Array(prefix string, offsetFromStart Offset, fieldName string, v []uint8) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint8Array)
	buf := make([]byte, FieldSizes[TypeUint8Array])
	WriteOffset(buf, Offset(len(v))*FieldSizes[TypeUint8])
	fmt.Printf("%s%x // %s: uint8 array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint8Array]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeUint8Array)
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpUint8(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpUint16Array(prefix string, offsetFromStart Offset, fieldName string, v []uint16) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint16Array)
	buf := make([]byte, FieldSizes[TypeUint16Array])
	WriteOffset(buf, Offset(len(v))*FieldSizes[TypeUint16])
	fmt.Printf("%s%x // %s: uint16 array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint16Array]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeUint16Array)
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpUint16(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpUint32Array(prefix string, offsetFromStart Offset, fieldName string, v []uint32) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint32Array)
	buf := make([]byte, FieldSizes[TypeUint32Array])
	WriteOffset(buf, Offset(len(v))*FieldSizes[TypeUint32])
	fmt.Printf("%s%x // %s: uint32 array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint32Array]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeUint32Array)
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpUint32(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpUint64Array(prefix string, offsetFromStart Offset, fieldName string, v []uint64) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint64Array)
	buf := make([]byte, FieldSizes[TypeUint64Array])
	WriteOffset(buf, Offset(len(v))*FieldSizes[TypeUint64])
	fmt.Printf("%s%x // %s: uint64 array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint64Array]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeUint64Array)
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpUint64(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpBytesArray(prefix string, offsetFromStart Offset, fieldName string, v [][]byte) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytesArray)
	sizePlaceholderOffset := w.size
	w.size += FieldSizes[TypeBytesArray]
	w.size = alignDynamicFieldContentOffset(w.size, TypeBytesArray)
	contentSizeStartOffset := w.size
	for _, vv := range v {
		w.WriteBytes(nil, vv)
	}
	contentSize := w.size - contentSizeStartOffset
	w.size = sizePlaceholderOffset
	buf := make([]byte, FieldSizes[TypeBytesArray])
	WriteOffset(buf, contentSize)
	fmt.Printf("%s%x // %s: bytes array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeBytesArray]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeBytesArray)
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpBytes(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpBytes20Array(prefix string, offsetFromStart Offset, fieldName string, v [][20]byte) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytes20Array)
	buf := make([]byte, FieldSizes[TypeBytes20Array])
	WriteOffset(buf, Offset(len(v)*20))
	fmt.Printf("%s%x // %s: bytes20 array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeBytes20Array]
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpBytes20(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpBytes32Array(prefix string, offsetFromStart Offset, fieldName string, v [][32]byte) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeBytes32Array)
	buf := make([]byte, FieldSizes[TypeBytes32Array])
	WriteOffset(buf, Offset(len(v)*32))
	fmt.Printf("%s%x // %s: bytes32 array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeBytes32Array]
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpBytes32(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpUint256Array(prefix string, offsetFromStart Offset, fieldName string, v []*big.Int) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeUint256Array)
	buf := make([]byte, FieldSizes[TypeUint256Array])
	WriteOffset(buf, Offset(len(v)*32))
	fmt.Printf("%s%x // %s: bytes32 array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeUint256Array]
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpUint256(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpStringArray(prefix string, offsetFromStart Offset, fieldName string, v []string) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeStringArray)
	sizePlaceholderOffset := w.size
	w.size += FieldSizes[TypeStringArray]
	w.size = alignDynamicFieldContentOffset(w.size, TypeStringArray)
	contentSizeStartOffset := w.size
	for _, vv := range v {
		w.WriteString(nil, vv)
	}
	contentSize := w.size - contentSizeStartOffset
	w.size = sizePlaceholderOffset
	buf := make([]byte, FieldSizes[TypeStringArray])
	WriteOffset(buf, contentSize)
	fmt.Printf("%s%x // %s: string array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeStringArray]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeStringArray)
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpString(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
}

func (w *InternalBuilder) HexDumpMessage(prefix string, offsetFromStart Offset, fieldName string, v MessageWriter) (err error) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeMessage)
	sizePlaceholderOffset := w.size
	w.size += FieldSizes[TypeMessage]
	w.size = alignDynamicFieldContentOffset(w.size, TypeMessage)
	err = v.Write(nil)
	contentSize := v.GetSize()
	w.size = sizePlaceholderOffset
	buf := make([]byte, FieldSizes[TypeMessage])
	WriteOffset(buf, contentSize)
	fmt.Printf("%s%x // %s: message size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeMessage]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeMessage)
	err = v.HexDump(prefix+HEX_DUMP_INDENT, offsetFromStart+w.size)
	w.size += contentSize
	return
}

func (w *InternalBuilder) HexDumpMessageArray(prefix string, offsetFromStart Offset, fieldName string, v []MessageWriter) (err error) {
	w.size = hexDumpAlignOffsetToType(prefix, offsetFromStart, w.size, TypeMessageArray)
	sizePlaceholderOffset := w.size
	w.size += FieldSizes[TypeMessageArray]
	w.size = alignDynamicFieldContentOffset(w.size, TypeMessageArray)
	contentSizeStartOffset := w.size
	for _, vv := range v {
		err = w.WriteMessage(nil, vv)
		if err != nil {
			return
		}
	}
	contentSize := w.size - contentSizeStartOffset
	w.size = sizePlaceholderOffset
	buf := make([]byte, FieldSizes[TypeMessageArray])
	WriteOffset(buf, contentSize)
	fmt.Printf("%s%x // %s: message array size (offset 0x%x, size: 0x%x)\n", prefix, buf, fieldName, offsetFromStart+w.size, len(buf))
	w.size += FieldSizes[TypeMessageArray]
	w.size = hexDumpAlignDynamicFieldContentOffset(prefix, offsetFromStart, w.size, TypeMessageArray)
	for index, vv := range v {
		fieldNameIndex := fmt.Sprintf("%s #%d", fieldName, index)
		w.HexDumpMessage(prefix+HEX_DUMP_INDENT, offsetFromStart, fieldNameIndex, vv)
	}
	return nil
}
