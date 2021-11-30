/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataGroupValueWrite struct {
	DataFirstByte int8
	Data          []byte
	Parent        *ApduData
}

// The corresponding interface
type IApduDataGroupValueWrite interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataGroupValueWrite) ApciType() uint8 {
	return 0x2
}

func (m *ApduDataGroupValueWrite) InitializeParent(parent *ApduData) {
}

func NewApduDataGroupValueWrite(dataFirstByte int8, data []byte) *ApduData {
	child := &ApduDataGroupValueWrite{
		DataFirstByte: dataFirstByte,
		Data:          data,
		Parent:        NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataGroupValueWrite(structType interface{}) *ApduDataGroupValueWrite {
	castFunc := func(typ interface{}) *ApduDataGroupValueWrite {
		if casted, ok := typ.(ApduDataGroupValueWrite); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataGroupValueWrite); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataGroupValueWrite(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataGroupValueWrite(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataGroupValueWrite) GetTypeName() string {
	return "ApduDataGroupValueWrite"
}

func (m *ApduDataGroupValueWrite) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataGroupValueWrite) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (dataFirstByte)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataGroupValueWrite) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataGroupValueWriteParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataGroupValueWrite"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (dataFirstByte)
	_dataFirstByte, _dataFirstByteErr := readBuffer.ReadInt8("dataFirstByte", 6)
	if _dataFirstByteErr != nil {
		return nil, errors.Wrap(_dataFirstByteErr, "Error parsing 'dataFirstByte' field")
	}
	dataFirstByte := _dataFirstByte
	// Byte Array field (data)
	numberOfBytesdata := int(utils.InlineIf(bool(bool((dataLength) < (1))), func() interface{} { return uint16(uint16(0)) }, func() interface{} { return uint16(uint16(dataLength) - uint16(uint16(1))) }).(uint16))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("ApduDataGroupValueWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataGroupValueWrite{
		DataFirstByte: dataFirstByte,
		Data:          data,
		Parent:        &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataGroupValueWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataGroupValueWrite"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dataFirstByte)
		dataFirstByte := int8(m.DataFirstByte)
		_dataFirstByteErr := writeBuffer.WriteInt8("dataFirstByte", 6, (dataFirstByte))
		if _dataFirstByteErr != nil {
			return errors.Wrap(_dataFirstByteErr, "Error serializing 'dataFirstByte' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataGroupValueWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataGroupValueWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
