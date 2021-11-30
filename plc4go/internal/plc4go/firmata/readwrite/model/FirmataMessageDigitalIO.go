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
type FirmataMessageDigitalIO struct {
	PinBlock uint8
	Data     []int8
	Parent   *FirmataMessage
}

// The corresponding interface
type IFirmataMessageDigitalIO interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataMessageDigitalIO) MessageType() uint8 {
	return 0x9
}

func (m *FirmataMessageDigitalIO) InitializeParent(parent *FirmataMessage) {
}

func NewFirmataMessageDigitalIO(pinBlock uint8, data []int8) *FirmataMessage {
	child := &FirmataMessageDigitalIO{
		PinBlock: pinBlock,
		Data:     data,
		Parent:   NewFirmataMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastFirmataMessageDigitalIO(structType interface{}) *FirmataMessageDigitalIO {
	castFunc := func(typ interface{}) *FirmataMessageDigitalIO {
		if casted, ok := typ.(FirmataMessageDigitalIO); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataMessageDigitalIO); ok {
			return casted
		}
		if casted, ok := typ.(FirmataMessage); ok {
			return CastFirmataMessageDigitalIO(casted.Child)
		}
		if casted, ok := typ.(*FirmataMessage); ok {
			return CastFirmataMessageDigitalIO(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataMessageDigitalIO) GetTypeName() string {
	return "FirmataMessageDigitalIO"
}

func (m *FirmataMessageDigitalIO) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *FirmataMessageDigitalIO) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (pinBlock)
	lengthInBits += 4

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *FirmataMessageDigitalIO) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmataMessageDigitalIOParse(readBuffer utils.ReadBuffer, response bool) (*FirmataMessage, error) {
	if pullErr := readBuffer.PullContext("FirmataMessageDigitalIO"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (pinBlock)
	_pinBlock, _pinBlockErr := readBuffer.ReadUint8("pinBlock", 4)
	if _pinBlockErr != nil {
		return nil, errors.Wrap(_pinBlockErr, "Error parsing 'pinBlock' field")
	}
	pinBlock := _pinBlock

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	data := make([]int8, uint16(2))
	for curItem := uint16(0); curItem < uint16(uint16(2)); curItem++ {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("FirmataMessageDigitalIO"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataMessageDigitalIO{
		PinBlock: pinBlock,
		Data:     data,
		Parent:   &FirmataMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *FirmataMessageDigitalIO) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageDigitalIO"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pinBlock)
		pinBlock := uint8(m.PinBlock)
		_pinBlockErr := writeBuffer.WriteUint8("pinBlock", 4, (pinBlock))
		if _pinBlockErr != nil {
			return errors.Wrap(_pinBlockErr, "Error serializing 'pinBlock' field")
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("FirmataMessageDigitalIO"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataMessageDigitalIO) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
