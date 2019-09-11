/**
 * Copyright 2018 the membuffers authors
 * This file is part of the membuffers library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

const FieldTypes = Object.freeze({
  TypeMessage: 1,
  TypeBytes: 2,
  TypeString: 3,
  TypeUnion: 4,
  TypeUint8: 11,
  TypeUint16: 12,
  TypeUint32: 13,
  TypeUint64: 14,
  TypeUint8Array: 21,
  TypeUint16Array: 22,
  TypeUint32Array: 23,
  TypeUint64Array: 24,
  TypeMessageArray: 31,
  TypeBytesArray: 32,
  TypeStringArray: 33,
});

const FieldSizes = Object.freeze({
  [FieldTypes.TypeMessage]: 4,
  [FieldTypes.TypeBytes]: 4,
  [FieldTypes.TypeString]: 4,
  [FieldTypes.TypeUnion]: 2,
  [FieldTypes.TypeUint8]: 1,
  [FieldTypes.TypeUint16]: 2,
  [FieldTypes.TypeUint32]: 4,
  [FieldTypes.TypeUint64]: 8,
  [FieldTypes.TypeUint8Array]: 4,
  [FieldTypes.TypeUint16Array]: 4,
  [FieldTypes.TypeUint32Array]: 4,
  [FieldTypes.TypeUint64Array]: 4,
  [FieldTypes.TypeMessageArray]: 4,
  [FieldTypes.TypeBytesArray]: 4,
  [FieldTypes.TypeStringArray]: 4,
});

const FieldAlignment = Object.freeze({
  [FieldTypes.TypeMessage]: 4,
  [FieldTypes.TypeBytes]: 4,
  [FieldTypes.TypeString]: 4,
  [FieldTypes.TypeUnion]: 2,
  [FieldTypes.TypeUint8]: 1,
  [FieldTypes.TypeUint16]: 2,
  [FieldTypes.TypeUint32]: 4,
  [FieldTypes.TypeUint64]: 4,
  [FieldTypes.TypeUint8Array]: 4,
  [FieldTypes.TypeUint16Array]: 4,
  [FieldTypes.TypeUint32Array]: 4,
  [FieldTypes.TypeUint64Array]: 4,
  [FieldTypes.TypeMessageArray]: 4,
  [FieldTypes.TypeBytesArray]: 4,
  [FieldTypes.TypeStringArray]: 4,
});

const FieldDynamic = Object.freeze({
  [FieldTypes.TypeMessage]: true,
  [FieldTypes.TypeBytes]: true,
  [FieldTypes.TypeString]: true,
  [FieldTypes.TypeUnion]: true,
  [FieldTypes.TypeUint8]: false,
  [FieldTypes.TypeUint16]: false,
  [FieldTypes.TypeUint32]: false,
  [FieldTypes.TypeUint64]: false,
  [FieldTypes.TypeUint8Array]: true,
  [FieldTypes.TypeUint16Array]: true,
  [FieldTypes.TypeUint32Array]: true,
  [FieldTypes.TypeUint64Array]: true,
  [FieldTypes.TypeMessageArray]: true,
  [FieldTypes.TypeBytesArray]: true,
  [FieldTypes.TypeStringArray]: true,
});

const FieldDynamicContentAlignment = Object.freeze({
  [FieldTypes.TypeMessage]: 4,
  [FieldTypes.TypeBytes]: 1,
  [FieldTypes.TypeString]: 1,
  [FieldTypes.TypeUnion]: 0,
  [FieldTypes.TypeUint8]: 0,
  [FieldTypes.TypeUint16]: 0,
  [FieldTypes.TypeUint32]: 0,
  [FieldTypes.TypeUint64]: 0,
  [FieldTypes.TypeUint8Array]: 1,
  [FieldTypes.TypeUint16Array]: 2,
  [FieldTypes.TypeUint32Array]: 4,
  [FieldTypes.TypeUint64Array]: 4,
  [FieldTypes.TypeMessageArray]: 4,
  [FieldTypes.TypeBytesArray]: 4,
  [FieldTypes.TypeStringArray]: 4,
});

module.exports = {FieldAlignment, FieldDynamic, FieldDynamicContentAlignment, FieldSizes, FieldTypes};
