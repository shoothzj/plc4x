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
type APDUConfirmedRequest struct {
	SegmentedMessage          bool
	MoreFollows               bool
	SegmentedResponseAccepted bool
	MaxSegmentsAccepted       uint8
	MaxApduLengthAccepted     uint8
	InvokeId                  uint8
	SequenceNumber            *uint8
	ProposedWindowSize        *uint8
	ServiceRequest            *BACnetConfirmedServiceRequest
	Parent                    *APDU
}

// The corresponding interface
type IAPDUConfirmedRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUConfirmedRequest) ApduType() uint8 {
	return 0x0
}

func (m *APDUConfirmedRequest) InitializeParent(parent *APDU) {
}

func NewAPDUConfirmedRequest(segmentedMessage bool, moreFollows bool, segmentedResponseAccepted bool, maxSegmentsAccepted uint8, maxApduLengthAccepted uint8, invokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceRequest *BACnetConfirmedServiceRequest) *APDU {
	child := &APDUConfirmedRequest{
		SegmentedMessage:          segmentedMessage,
		MoreFollows:               moreFollows,
		SegmentedResponseAccepted: segmentedResponseAccepted,
		MaxSegmentsAccepted:       maxSegmentsAccepted,
		MaxApduLengthAccepted:     maxApduLengthAccepted,
		InvokeId:                  invokeId,
		SequenceNumber:            sequenceNumber,
		ProposedWindowSize:        proposedWindowSize,
		ServiceRequest:            serviceRequest,
		Parent:                    NewAPDU(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAPDUConfirmedRequest(structType interface{}) *APDUConfirmedRequest {
	castFunc := func(typ interface{}) *APDUConfirmedRequest {
		if casted, ok := typ.(APDUConfirmedRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUConfirmedRequest); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUConfirmedRequest(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUConfirmedRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUConfirmedRequest) GetTypeName() string {
	return "APDUConfirmedRequest"
}

func (m *APDUConfirmedRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *APDUConfirmedRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Simple field (segmentedResponseAccepted)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (maxSegmentsAccepted)
	lengthInBits += 3

	// Simple field (maxApduLengthAccepted)
	lengthInBits += 4

	// Simple field (invokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// Simple field (serviceRequest)
	lengthInBits += m.ServiceRequest.LengthInBits()

	return lengthInBits
}

func (m *APDUConfirmedRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func APDUConfirmedRequestParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDUConfirmedRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (segmentedMessage)
	_segmentedMessage, _segmentedMessageErr := readBuffer.ReadBit("segmentedMessage")
	if _segmentedMessageErr != nil {
		return nil, errors.Wrap(_segmentedMessageErr, "Error parsing 'segmentedMessage' field")
	}
	segmentedMessage := _segmentedMessage

	// Simple Field (moreFollows)
	_moreFollows, _moreFollowsErr := readBuffer.ReadBit("moreFollows")
	if _moreFollowsErr != nil {
		return nil, errors.Wrap(_moreFollowsErr, "Error parsing 'moreFollows' field")
	}
	moreFollows := _moreFollows

	// Simple Field (segmentedResponseAccepted)
	_segmentedResponseAccepted, _segmentedResponseAcceptedErr := readBuffer.ReadBit("segmentedResponseAccepted")
	if _segmentedResponseAcceptedErr != nil {
		return nil, errors.Wrap(_segmentedResponseAcceptedErr, "Error parsing 'segmentedResponseAccepted' field")
	}
	segmentedResponseAccepted := _segmentedResponseAccepted

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (maxSegmentsAccepted)
	_maxSegmentsAccepted, _maxSegmentsAcceptedErr := readBuffer.ReadUint8("maxSegmentsAccepted", 3)
	if _maxSegmentsAcceptedErr != nil {
		return nil, errors.Wrap(_maxSegmentsAcceptedErr, "Error parsing 'maxSegmentsAccepted' field")
	}
	maxSegmentsAccepted := _maxSegmentsAccepted

	// Simple Field (maxApduLengthAccepted)
	_maxApduLengthAccepted, _maxApduLengthAcceptedErr := readBuffer.ReadUint8("maxApduLengthAccepted", 4)
	if _maxApduLengthAcceptedErr != nil {
		return nil, errors.Wrap(_maxApduLengthAcceptedErr, "Error parsing 'maxApduLengthAccepted' field")
	}
	maxApduLengthAccepted := _maxApduLengthAccepted

	// Simple Field (invokeId)
	_invokeId, _invokeIdErr := readBuffer.ReadUint8("invokeId", 8)
	if _invokeIdErr != nil {
		return nil, errors.Wrap(_invokeIdErr, "Error parsing 'invokeId' field")
	}
	invokeId := _invokeId

	// Optional Field (sequenceNumber) (Can be skipped, if a given expression evaluates to false)
	var sequenceNumber *uint8 = nil
	if segmentedMessage {
		_val, _err := readBuffer.ReadUint8("sequenceNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'sequenceNumber' field")
		}
		sequenceNumber = &_val
	}

	// Optional Field (proposedWindowSize) (Can be skipped, if a given expression evaluates to false)
	var proposedWindowSize *uint8 = nil
	if segmentedMessage {
		_val, _err := readBuffer.ReadUint8("proposedWindowSize", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'proposedWindowSize' field")
		}
		proposedWindowSize = &_val
	}

	// Simple Field (serviceRequest)
	if pullErr := readBuffer.PullContext("serviceRequest"); pullErr != nil {
		return nil, pullErr
	}
	_serviceRequest, _serviceRequestErr := BACnetConfirmedServiceRequestParse(readBuffer, uint16(apduLength)-uint16(uint16(uint16(uint16(3))+uint16(uint16(utils.InlineIf(segmentedMessage, func() interface{} { return uint16(uint16(2)) }, func() interface{} { return uint16(uint16(0)) }).(uint16))))))
	if _serviceRequestErr != nil {
		return nil, errors.Wrap(_serviceRequestErr, "Error parsing 'serviceRequest' field")
	}
	serviceRequest := CastBACnetConfirmedServiceRequest(_serviceRequest)
	if closeErr := readBuffer.CloseContext("serviceRequest"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("APDUConfirmedRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUConfirmedRequest{
		SegmentedMessage:          segmentedMessage,
		MoreFollows:               moreFollows,
		SegmentedResponseAccepted: segmentedResponseAccepted,
		MaxSegmentsAccepted:       maxSegmentsAccepted,
		MaxApduLengthAccepted:     maxApduLengthAccepted,
		InvokeId:                  invokeId,
		SequenceNumber:            sequenceNumber,
		ProposedWindowSize:        proposedWindowSize,
		ServiceRequest:            CastBACnetConfirmedServiceRequest(serviceRequest),
		Parent:                    &APDU{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *APDUConfirmedRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUConfirmedRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (segmentedMessage)
		segmentedMessage := bool(m.SegmentedMessage)
		_segmentedMessageErr := writeBuffer.WriteBit("segmentedMessage", (segmentedMessage))
		if _segmentedMessageErr != nil {
			return errors.Wrap(_segmentedMessageErr, "Error serializing 'segmentedMessage' field")
		}

		// Simple Field (moreFollows)
		moreFollows := bool(m.MoreFollows)
		_moreFollowsErr := writeBuffer.WriteBit("moreFollows", (moreFollows))
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
		}

		// Simple Field (segmentedResponseAccepted)
		segmentedResponseAccepted := bool(m.SegmentedResponseAccepted)
		_segmentedResponseAcceptedErr := writeBuffer.WriteBit("segmentedResponseAccepted", (segmentedResponseAccepted))
		if _segmentedResponseAcceptedErr != nil {
			return errors.Wrap(_segmentedResponseAcceptedErr, "Error serializing 'segmentedResponseAccepted' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 2, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (maxSegmentsAccepted)
		maxSegmentsAccepted := uint8(m.MaxSegmentsAccepted)
		_maxSegmentsAcceptedErr := writeBuffer.WriteUint8("maxSegmentsAccepted", 3, (maxSegmentsAccepted))
		if _maxSegmentsAcceptedErr != nil {
			return errors.Wrap(_maxSegmentsAcceptedErr, "Error serializing 'maxSegmentsAccepted' field")
		}

		// Simple Field (maxApduLengthAccepted)
		maxApduLengthAccepted := uint8(m.MaxApduLengthAccepted)
		_maxApduLengthAcceptedErr := writeBuffer.WriteUint8("maxApduLengthAccepted", 4, (maxApduLengthAccepted))
		if _maxApduLengthAcceptedErr != nil {
			return errors.Wrap(_maxApduLengthAcceptedErr, "Error serializing 'maxApduLengthAccepted' field")
		}

		// Simple Field (invokeId)
		invokeId := uint8(m.InvokeId)
		_invokeIdErr := writeBuffer.WriteUint8("invokeId", 8, (invokeId))
		if _invokeIdErr != nil {
			return errors.Wrap(_invokeIdErr, "Error serializing 'invokeId' field")
		}

		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.SequenceNumber != nil {
			sequenceNumber = m.SequenceNumber
			_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, *(sequenceNumber))
			if _sequenceNumberErr != nil {
				return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
			}
		}

		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.ProposedWindowSize != nil {
			proposedWindowSize = m.ProposedWindowSize
			_proposedWindowSizeErr := writeBuffer.WriteUint8("proposedWindowSize", 8, *(proposedWindowSize))
			if _proposedWindowSizeErr != nil {
				return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
			}
		}

		// Simple Field (serviceRequest)
		if pushErr := writeBuffer.PushContext("serviceRequest"); pushErr != nil {
			return pushErr
		}
		_serviceRequestErr := m.ServiceRequest.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("serviceRequest"); popErr != nil {
			return popErr
		}
		if _serviceRequestErr != nil {
			return errors.Wrap(_serviceRequestErr, "Error serializing 'serviceRequest' field")
		}

		if popErr := writeBuffer.PopContext("APDUConfirmedRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUConfirmedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
