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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type S7PayloadUserDataItemCpuFunctionMsgSubscription struct {
	Subscription uint8
	MagicKey     string
	Alarmtype    *AlarmStateType
	Reserve      *uint8
	Parent       *S7PayloadUserDataItem
}

// The corresponding interface
type IS7PayloadUserDataItemCpuFunctionMsgSubscription interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) CpuFunctionType() uint8 {
	return 0x04
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) CpuSubfunction() uint8 {
	return 0x02
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) DataLength() uint16 {
	return 0
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.Parent.ReturnCode = returnCode
	m.Parent.TransportSize = transportSize
}

func NewS7PayloadUserDataItemCpuFunctionMsgSubscription(Subscription uint8, magicKey string, Alarmtype *AlarmStateType, Reserve *uint8, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	child := &S7PayloadUserDataItemCpuFunctionMsgSubscription{
		Subscription: Subscription,
		MagicKey:     magicKey,
		Alarmtype:    Alarmtype,
		Reserve:      Reserve,
		Parent:       NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7PayloadUserDataItemCpuFunctionMsgSubscription(structType interface{}) *S7PayloadUserDataItemCpuFunctionMsgSubscription {
	castFunc := func(typ interface{}) *S7PayloadUserDataItemCpuFunctionMsgSubscription {
		if casted, ok := typ.(S7PayloadUserDataItemCpuFunctionMsgSubscription); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadUserDataItemCpuFunctionMsgSubscription); ok {
			return casted
		}
		if casted, ok := typ.(S7PayloadUserDataItem); ok {
			return CastS7PayloadUserDataItemCpuFunctionMsgSubscription(casted.Child)
		}
		if casted, ok := typ.(*S7PayloadUserDataItem); ok {
			return CastS7PayloadUserDataItemCpuFunctionMsgSubscription(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionMsgSubscription"
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (Subscription)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (magicKey)
	lengthInBits += 64

	// Optional Field (Alarmtype)
	if m.Alarmtype != nil {
		lengthInBits += 8
	}

	// Optional Field (Reserve)
	if m.Reserve != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionMsgSubscriptionParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionMsgSubscription"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (Subscription)
	_Subscription, _SubscriptionErr := readBuffer.ReadUint8("Subscription", 8)
	if _SubscriptionErr != nil {
		return nil, errors.Wrap(_SubscriptionErr, "Error parsing 'Subscription' field")
	}
	Subscription := _Subscription

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (magicKey)
	_magicKey, _magicKeyErr := readBuffer.ReadString("magicKey", uint32(64))
	if _magicKeyErr != nil {
		return nil, errors.Wrap(_magicKeyErr, "Error parsing 'magicKey' field")
	}
	magicKey := _magicKey

	// Optional Field (Alarmtype) (Can be skipped, if a given expression evaluates to false)
	var Alarmtype *AlarmStateType = nil
	if bool((Subscription) >= (128)) {
		if pullErr := readBuffer.PullContext("Alarmtype"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := AlarmStateTypeParse(readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'Alarmtype' field")
		}
		Alarmtype = &_val
		if closeErr := readBuffer.CloseContext("Alarmtype"); closeErr != nil {
			return nil, closeErr
		}
	}

	// Optional Field (Reserve) (Can be skipped, if a given expression evaluates to false)
	var Reserve *uint8 = nil
	if bool((Subscription) >= (128)) {
		_val, _err := readBuffer.ReadUint8("Reserve", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'Reserve' field")
		}
		Reserve = &_val
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionMsgSubscription"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserDataItemCpuFunctionMsgSubscription{
		Subscription: Subscription,
		MagicKey:     magicKey,
		Alarmtype:    Alarmtype,
		Reserve:      Reserve,
		Parent:       &S7PayloadUserDataItem{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionMsgSubscription"); pushErr != nil {
			return pushErr
		}

		// Simple Field (Subscription)
		Subscription := uint8(m.Subscription)
		_SubscriptionErr := writeBuffer.WriteUint8("Subscription", 8, (Subscription))
		if _SubscriptionErr != nil {
			return errors.Wrap(_SubscriptionErr, "Error serializing 'Subscription' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (magicKey)
		magicKey := string(m.MagicKey)
		_magicKeyErr := writeBuffer.WriteString("magicKey", uint32(64), "UTF-8", (magicKey))
		if _magicKeyErr != nil {
			return errors.Wrap(_magicKeyErr, "Error serializing 'magicKey' field")
		}

		// Optional Field (Alarmtype) (Can be skipped, if the value is null)
		var Alarmtype *AlarmStateType = nil
		if m.Alarmtype != nil {
			if pushErr := writeBuffer.PushContext("Alarmtype"); pushErr != nil {
				return pushErr
			}
			Alarmtype = m.Alarmtype
			_AlarmtypeErr := Alarmtype.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("Alarmtype"); popErr != nil {
				return popErr
			}
			if _AlarmtypeErr != nil {
				return errors.Wrap(_AlarmtypeErr, "Error serializing 'Alarmtype' field")
			}
		}

		// Optional Field (Reserve) (Can be skipped, if the value is null)
		var Reserve *uint8 = nil
		if m.Reserve != nil {
			Reserve = m.Reserve
			_ReserveErr := writeBuffer.WriteUint8("Reserve", 8, *(Reserve))
			if _ReserveErr != nil {
				return errors.Wrap(_ReserveErr, "Error serializing 'Reserve' field")
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionMsgSubscription"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserDataItemCpuFunctionMsgSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
