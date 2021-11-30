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
type S7PayloadAlarmAckInd struct {
	AlarmMessage *AlarmMessageAckPushType
	Parent       *S7PayloadUserDataItem
}

// The corresponding interface
type IS7PayloadAlarmAckInd interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadAlarmAckInd) CpuFunctionType() uint8 {
	return 0x00
}

func (m *S7PayloadAlarmAckInd) CpuSubfunction() uint8 {
	return 0x0c
}

func (m *S7PayloadAlarmAckInd) DataLength() uint16 {
	return 0
}

func (m *S7PayloadAlarmAckInd) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.Parent.ReturnCode = returnCode
	m.Parent.TransportSize = transportSize
}

func NewS7PayloadAlarmAckInd(alarmMessage *AlarmMessageAckPushType, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	child := &S7PayloadAlarmAckInd{
		AlarmMessage: alarmMessage,
		Parent:       NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7PayloadAlarmAckInd(structType interface{}) *S7PayloadAlarmAckInd {
	castFunc := func(typ interface{}) *S7PayloadAlarmAckInd {
		if casted, ok := typ.(S7PayloadAlarmAckInd); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadAlarmAckInd); ok {
			return casted
		}
		if casted, ok := typ.(S7PayloadUserDataItem); ok {
			return CastS7PayloadAlarmAckInd(casted.Child)
		}
		if casted, ok := typ.(*S7PayloadUserDataItem); ok {
			return CastS7PayloadAlarmAckInd(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadAlarmAckInd) GetTypeName() string {
	return "S7PayloadAlarmAckInd"
}

func (m *S7PayloadAlarmAckInd) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadAlarmAckInd) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.LengthInBits()

	return lengthInBits
}

func (m *S7PayloadAlarmAckInd) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadAlarmAckIndParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7PayloadAlarmAckInd"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (alarmMessage)
	if pullErr := readBuffer.PullContext("alarmMessage"); pullErr != nil {
		return nil, pullErr
	}
	_alarmMessage, _alarmMessageErr := AlarmMessageAckPushTypeParse(readBuffer)
	if _alarmMessageErr != nil {
		return nil, errors.Wrap(_alarmMessageErr, "Error parsing 'alarmMessage' field")
	}
	alarmMessage := CastAlarmMessageAckPushType(_alarmMessage)
	if closeErr := readBuffer.CloseContext("alarmMessage"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadAlarmAckInd"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadAlarmAckInd{
		AlarmMessage: CastAlarmMessageAckPushType(alarmMessage),
		Parent:       &S7PayloadUserDataItem{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7PayloadAlarmAckInd) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadAlarmAckInd"); pushErr != nil {
			return pushErr
		}

		// Simple Field (alarmMessage)
		if pushErr := writeBuffer.PushContext("alarmMessage"); pushErr != nil {
			return pushErr
		}
		_alarmMessageErr := m.AlarmMessage.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("alarmMessage"); popErr != nil {
			return popErr
		}
		if _alarmMessageErr != nil {
			return errors.Wrap(_alarmMessageErr, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadAlarmAckInd"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadAlarmAckInd) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
