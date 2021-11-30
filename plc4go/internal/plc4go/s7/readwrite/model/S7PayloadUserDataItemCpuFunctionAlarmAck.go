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
type S7PayloadUserDataItemCpuFunctionAlarmAck struct {
	FunctionId     uint8
	MessageObjects []*AlarmMessageObjectAckType
	Parent         *S7PayloadUserDataItem
}

// The corresponding interface
type IS7PayloadUserDataItemCpuFunctionAlarmAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) CpuFunctionType() uint8 {
	return 0x04
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) CpuSubfunction() uint8 {
	return 0x0b
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) DataLength() uint16 {
	return 0
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.Parent.ReturnCode = returnCode
	m.Parent.TransportSize = transportSize
}

func NewS7PayloadUserDataItemCpuFunctionAlarmAck(functionId uint8, messageObjects []*AlarmMessageObjectAckType, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	child := &S7PayloadUserDataItemCpuFunctionAlarmAck{
		FunctionId:     functionId,
		MessageObjects: messageObjects,
		Parent:         NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7PayloadUserDataItemCpuFunctionAlarmAck(structType interface{}) *S7PayloadUserDataItemCpuFunctionAlarmAck {
	castFunc := func(typ interface{}) *S7PayloadUserDataItemCpuFunctionAlarmAck {
		if casted, ok := typ.(S7PayloadUserDataItemCpuFunctionAlarmAck); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadUserDataItemCpuFunctionAlarmAck); ok {
			return casted
		}
		if casted, ok := typ.(S7PayloadUserDataItem); ok {
			return CastS7PayloadUserDataItemCpuFunctionAlarmAck(casted.Child)
		}
		if casted, ok := typ.(*S7PayloadUserDataItem); ok {
			return CastS7PayloadUserDataItemCpuFunctionAlarmAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionAlarmAck"
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (functionId)
	lengthInBits += 8

	// Implicit Field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for i, element := range m.MessageObjects {
			last := i == len(m.MessageObjects)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionAlarmAckParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionAlarmAck"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (functionId)
	_functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field")
	}
	functionId := _functionId

	// Implicit Field (numberOfObjects) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	_ = numberOfObjects
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field")
	}

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	messageObjects := make([]*AlarmMessageObjectAckType, numberOfObjects)
	for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
		_item, _err := AlarmMessageObjectAckTypeParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field")
		}
		messageObjects[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionAlarmAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionAlarmAck{
		FunctionId:     functionId,
		MessageObjects: messageObjects,
		Parent:         &S7PayloadUserDataItem{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionAlarmAck"); pushErr != nil {
			return pushErr
		}

		// Simple Field (functionId)
		functionId := uint8(m.FunctionId)
		_functionIdErr := writeBuffer.WriteUint8("functionId", 8, (functionId))
		if _functionIdErr != nil {
			return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
		}

		// Implicit Field (numberOfObjects) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numberOfObjects := uint8(uint8(len(m.MessageObjects)))
		_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
		if _numberOfObjectsErr != nil {
			return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
		}

		// Array Field (messageObjects)
		if m.MessageObjects != nil {
			if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.MessageObjects {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
				}
			}
			if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionAlarmAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionAlarmAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
