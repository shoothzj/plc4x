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
type TunnelingRequest struct {
	TunnelingRequestDataBlock *TunnelingRequestDataBlock
	Cemi                      *CEMI
	Parent                    *KnxNetIpMessage
}

// The corresponding interface
type ITunnelingRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *TunnelingRequest) MsgType() uint16 {
	return 0x0420
}

func (m *TunnelingRequest) InitializeParent(parent *KnxNetIpMessage) {
}

func NewTunnelingRequest(tunnelingRequestDataBlock *TunnelingRequestDataBlock, cemi *CEMI) *KnxNetIpMessage {
	child := &TunnelingRequest{
		TunnelingRequestDataBlock: tunnelingRequestDataBlock,
		Cemi:                      cemi,
		Parent:                    NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastTunnelingRequest(structType interface{}) *TunnelingRequest {
	castFunc := func(typ interface{}) *TunnelingRequest {
		if casted, ok := typ.(TunnelingRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*TunnelingRequest); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastTunnelingRequest(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastTunnelingRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *TunnelingRequest) GetTypeName() string {
	return "TunnelingRequest"
}

func (m *TunnelingRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *TunnelingRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (tunnelingRequestDataBlock)
	lengthInBits += m.TunnelingRequestDataBlock.LengthInBits()

	// Simple field (cemi)
	lengthInBits += m.Cemi.LengthInBits()

	return lengthInBits
}

func (m *TunnelingRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func TunnelingRequestParse(readBuffer utils.ReadBuffer, totalLength uint16) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("TunnelingRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (tunnelingRequestDataBlock)
	if pullErr := readBuffer.PullContext("tunnelingRequestDataBlock"); pullErr != nil {
		return nil, pullErr
	}
	tunnelingRequestDataBlock, _tunnelingRequestDataBlockErr := TunnelingRequestDataBlockParse(readBuffer)
	if _tunnelingRequestDataBlockErr != nil {
		return nil, errors.Wrap(_tunnelingRequestDataBlockErr, "Error parsing 'tunnelingRequestDataBlock' field")
	}
	if closeErr := readBuffer.CloseContext("tunnelingRequestDataBlock"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (cemi)
	if pullErr := readBuffer.PullContext("cemi"); pullErr != nil {
		return nil, pullErr
	}
	cemi, _cemiErr := CEMIParse(readBuffer, uint8(totalLength)-uint8(uint8(uint8(uint8(6))+uint8(tunnelingRequestDataBlock.LengthInBytes()))))
	if _cemiErr != nil {
		return nil, errors.Wrap(_cemiErr, "Error parsing 'cemi' field")
	}
	if closeErr := readBuffer.CloseContext("cemi"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("TunnelingRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &TunnelingRequest{
		TunnelingRequestDataBlock: tunnelingRequestDataBlock,
		Cemi:                      cemi,
		Parent:                    &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *TunnelingRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TunnelingRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (tunnelingRequestDataBlock)
		if pushErr := writeBuffer.PushContext("tunnelingRequestDataBlock"); pushErr != nil {
			return pushErr
		}
		_tunnelingRequestDataBlockErr := m.TunnelingRequestDataBlock.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("tunnelingRequestDataBlock"); popErr != nil {
			return popErr
		}
		if _tunnelingRequestDataBlockErr != nil {
			return errors.Wrap(_tunnelingRequestDataBlockErr, "Error serializing 'tunnelingRequestDataBlock' field")
		}

		// Simple Field (cemi)
		if pushErr := writeBuffer.PushContext("cemi"); pushErr != nil {
			return pushErr
		}
		_cemiErr := m.Cemi.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("cemi"); popErr != nil {
			return popErr
		}
		if _cemiErr != nil {
			return errors.Wrap(_cemiErr, "Error serializing 'cemi' field")
		}

		if popErr := writeBuffer.PopContext("TunnelingRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *TunnelingRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
