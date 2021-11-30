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
type AdsDeleteDeviceNotificationRequest struct {
	NotificationHandle uint32
	Parent             *AdsData
}

// The corresponding interface
type IAdsDeleteDeviceNotificationRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsDeleteDeviceNotificationRequest) CommandId() CommandId {
	return CommandId_ADS_DELETE_DEVICE_NOTIFICATION
}

func (m *AdsDeleteDeviceNotificationRequest) Response() bool {
	return bool(false)
}

func (m *AdsDeleteDeviceNotificationRequest) InitializeParent(parent *AdsData) {
}

func NewAdsDeleteDeviceNotificationRequest(notificationHandle uint32) *AdsData {
	child := &AdsDeleteDeviceNotificationRequest{
		NotificationHandle: notificationHandle,
		Parent:             NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsDeleteDeviceNotificationRequest(structType interface{}) *AdsDeleteDeviceNotificationRequest {
	castFunc := func(typ interface{}) *AdsDeleteDeviceNotificationRequest {
		if casted, ok := typ.(AdsDeleteDeviceNotificationRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsDeleteDeviceNotificationRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsDeleteDeviceNotificationRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsDeleteDeviceNotificationRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsDeleteDeviceNotificationRequest) GetTypeName() string {
	return "AdsDeleteDeviceNotificationRequest"
}

func (m *AdsDeleteDeviceNotificationRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsDeleteDeviceNotificationRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (notificationHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsDeleteDeviceNotificationRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsDeleteDeviceNotificationRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsDeleteDeviceNotificationRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (notificationHandle)
	_notificationHandle, _notificationHandleErr := readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field")
	}
	notificationHandle := _notificationHandle

	if closeErr := readBuffer.CloseContext("AdsDeleteDeviceNotificationRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsDeleteDeviceNotificationRequest{
		NotificationHandle: notificationHandle,
		Parent:             &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsDeleteDeviceNotificationRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDeleteDeviceNotificationRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (notificationHandle)
		notificationHandle := uint32(m.NotificationHandle)
		_notificationHandleErr := writeBuffer.WriteUint32("notificationHandle", 32, (notificationHandle))
		if _notificationHandleErr != nil {
			return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
		}

		if popErr := writeBuffer.PopContext("AdsDeleteDeviceNotificationRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsDeleteDeviceNotificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
