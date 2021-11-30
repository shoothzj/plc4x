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
type LDataExtended struct {
	GroupAddress        bool
	HopCount            uint8
	ExtendedFrameFormat uint8
	SourceAddress       *KnxAddress
	DestinationAddress  []byte
	Apdu                *Apdu
	Parent              *LDataFrame
}

// The corresponding interface
type ILDataExtended interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LDataExtended) NotAckFrame() bool {
	return bool(true)
}

func (m *LDataExtended) Polling() bool {
	return bool(false)
}

func (m *LDataExtended) InitializeParent(parent *LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) {
	m.Parent.FrameType = frameType
	m.Parent.NotRepeated = notRepeated
	m.Parent.Priority = priority
	m.Parent.AcknowledgeRequested = acknowledgeRequested
	m.Parent.ErrorFlag = errorFlag
}

func NewLDataExtended(groupAddress bool, hopCount uint8, extendedFrameFormat uint8, sourceAddress *KnxAddress, destinationAddress []byte, apdu *Apdu, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *LDataFrame {
	child := &LDataExtended{
		GroupAddress:        groupAddress,
		HopCount:            hopCount,
		ExtendedFrameFormat: extendedFrameFormat,
		SourceAddress:       sourceAddress,
		DestinationAddress:  destinationAddress,
		Apdu:                apdu,
		Parent:              NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLDataExtended(structType interface{}) *LDataExtended {
	castFunc := func(typ interface{}) *LDataExtended {
		if casted, ok := typ.(LDataExtended); ok {
			return &casted
		}
		if casted, ok := typ.(*LDataExtended); ok {
			return casted
		}
		if casted, ok := typ.(LDataFrame); ok {
			return CastLDataExtended(casted.Child)
		}
		if casted, ok := typ.(*LDataFrame); ok {
			return CastLDataExtended(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LDataExtended) GetTypeName() string {
	return "LDataExtended"
}

func (m *LDataExtended) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LDataExtended) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (groupAddress)
	lengthInBits += 1

	// Simple field (hopCount)
	lengthInBits += 3

	// Simple field (extendedFrameFormat)
	lengthInBits += 4

	// Simple field (sourceAddress)
	lengthInBits += m.SourceAddress.LengthInBits()

	// Array field
	if len(m.DestinationAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.DestinationAddress))
	}

	// Implicit Field (dataLength)
	lengthInBits += 8

	// Simple field (apdu)
	lengthInBits += m.Apdu.LengthInBits()

	return lengthInBits
}

func (m *LDataExtended) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LDataExtendedParse(readBuffer utils.ReadBuffer) (*LDataFrame, error) {
	if pullErr := readBuffer.PullContext("LDataExtended"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (groupAddress)
	_groupAddress, _groupAddressErr := readBuffer.ReadBit("groupAddress")
	if _groupAddressErr != nil {
		return nil, errors.Wrap(_groupAddressErr, "Error parsing 'groupAddress' field")
	}
	groupAddress := _groupAddress

	// Simple Field (hopCount)
	_hopCount, _hopCountErr := readBuffer.ReadUint8("hopCount", 3)
	if _hopCountErr != nil {
		return nil, errors.Wrap(_hopCountErr, "Error parsing 'hopCount' field")
	}
	hopCount := _hopCount

	// Simple Field (extendedFrameFormat)
	_extendedFrameFormat, _extendedFrameFormatErr := readBuffer.ReadUint8("extendedFrameFormat", 4)
	if _extendedFrameFormatErr != nil {
		return nil, errors.Wrap(_extendedFrameFormatErr, "Error parsing 'extendedFrameFormat' field")
	}
	extendedFrameFormat := _extendedFrameFormat

	// Simple Field (sourceAddress)
	if pullErr := readBuffer.PullContext("sourceAddress"); pullErr != nil {
		return nil, pullErr
	}
	_sourceAddress, _sourceAddressErr := KnxAddressParse(readBuffer)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field")
	}
	sourceAddress := CastKnxAddress(_sourceAddress)
	if closeErr := readBuffer.CloseContext("sourceAddress"); closeErr != nil {
		return nil, closeErr
	}
	// Byte Array field (destinationAddress)
	numberOfBytesdestinationAddress := int(uint16(2))
	destinationAddress, _readArrayErr := readBuffer.ReadByteArray("destinationAddress", numberOfBytesdestinationAddress)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'destinationAddress' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := readBuffer.ReadUint8("dataLength", 8)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field")
	}

	// Simple Field (apdu)
	if pullErr := readBuffer.PullContext("apdu"); pullErr != nil {
		return nil, pullErr
	}
	_apdu, _apduErr := ApduParse(readBuffer, dataLength)
	if _apduErr != nil {
		return nil, errors.Wrap(_apduErr, "Error parsing 'apdu' field")
	}
	apdu := CastApdu(_apdu)
	if closeErr := readBuffer.CloseContext("apdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("LDataExtended"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &LDataExtended{
		GroupAddress:        groupAddress,
		HopCount:            hopCount,
		ExtendedFrameFormat: extendedFrameFormat,
		SourceAddress:       CastKnxAddress(sourceAddress),
		DestinationAddress:  destinationAddress,
		Apdu:                CastApdu(apdu),
		Parent:              &LDataFrame{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LDataExtended) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataExtended"); pushErr != nil {
			return pushErr
		}

		// Simple Field (groupAddress)
		groupAddress := bool(m.GroupAddress)
		_groupAddressErr := writeBuffer.WriteBit("groupAddress", (groupAddress))
		if _groupAddressErr != nil {
			return errors.Wrap(_groupAddressErr, "Error serializing 'groupAddress' field")
		}

		// Simple Field (hopCount)
		hopCount := uint8(m.HopCount)
		_hopCountErr := writeBuffer.WriteUint8("hopCount", 3, (hopCount))
		if _hopCountErr != nil {
			return errors.Wrap(_hopCountErr, "Error serializing 'hopCount' field")
		}

		// Simple Field (extendedFrameFormat)
		extendedFrameFormat := uint8(m.ExtendedFrameFormat)
		_extendedFrameFormatErr := writeBuffer.WriteUint8("extendedFrameFormat", 4, (extendedFrameFormat))
		if _extendedFrameFormatErr != nil {
			return errors.Wrap(_extendedFrameFormatErr, "Error serializing 'extendedFrameFormat' field")
		}

		// Simple Field (sourceAddress)
		if pushErr := writeBuffer.PushContext("sourceAddress"); pushErr != nil {
			return pushErr
		}
		_sourceAddressErr := m.SourceAddress.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("sourceAddress"); popErr != nil {
			return popErr
		}
		if _sourceAddressErr != nil {
			return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
		}

		// Array Field (destinationAddress)
		if m.DestinationAddress != nil {
			// Byte Array field (destinationAddress)
			_writeArrayErr := writeBuffer.WriteByteArray("destinationAddress", m.DestinationAddress)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'destinationAddress' field")
			}
		}

		// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		dataLength := uint8(uint8(m.Apdu.LengthInBytes()) - uint8(uint8(1)))
		_dataLengthErr := writeBuffer.WriteUint8("dataLength", 8, (dataLength))
		if _dataLengthErr != nil {
			return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
		}

		// Simple Field (apdu)
		if pushErr := writeBuffer.PushContext("apdu"); pushErr != nil {
			return pushErr
		}
		_apduErr := m.Apdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("apdu"); popErr != nil {
			return popErr
		}
		if _apduErr != nil {
			return errors.Wrap(_apduErr, "Error serializing 'apdu' field")
		}

		if popErr := writeBuffer.PopContext("LDataExtended"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *LDataExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
