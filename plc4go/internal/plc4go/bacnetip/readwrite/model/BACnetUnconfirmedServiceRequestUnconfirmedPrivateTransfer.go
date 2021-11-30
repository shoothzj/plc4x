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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_VENDORIDHEADER uint8 = 0x09
const BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_SERVICENUMBERHEADER uint8 = 0x1A
const BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESOPENINGTAG uint8 = 0x2E
const BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESCLOSINGTAG uint8 = 0x2F

// The data-structure of this message
type BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer struct {
	VendorId      uint8
	ServiceNumber uint16
	Values        []int8
	Parent        *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) ServiceChoice() uint8 {
	return 0x04
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(vendorId uint8, serviceNumber uint16, values []int8) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		VendorId:      vendorId,
		ServiceNumber: serviceNumber,
		Values:        values,
		Parent:        NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (vendorIdHeader)
	lengthInBits += 8

	// Simple field (vendorId)
	lengthInBits += 8

	// Const Field (serviceNumberHeader)
	lengthInBits += 8

	// Simple field (serviceNumber)
	lengthInBits += 16

	// Const Field (listOfValuesOpeningTag)
	lengthInBits += 8

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	// Const Field (listOfValuesClosingTag)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (vendorIdHeader)
	vendorIdHeader, _vendorIdHeaderErr := readBuffer.ReadUint8("vendorIdHeader", 8)
	if _vendorIdHeaderErr != nil {
		return nil, errors.Wrap(_vendorIdHeaderErr, "Error parsing 'vendorIdHeader' field")
	}
	if vendorIdHeader != BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_VENDORIDHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_VENDORIDHEADER) + " but got " + fmt.Sprintf("%d", vendorIdHeader))
	}

	// Simple Field (vendorId)
	_vendorId, _vendorIdErr := readBuffer.ReadUint8("vendorId", 8)
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}
	vendorId := _vendorId

	// Const Field (serviceNumberHeader)
	serviceNumberHeader, _serviceNumberHeaderErr := readBuffer.ReadUint8("serviceNumberHeader", 8)
	if _serviceNumberHeaderErr != nil {
		return nil, errors.Wrap(_serviceNumberHeaderErr, "Error parsing 'serviceNumberHeader' field")
	}
	if serviceNumberHeader != BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_SERVICENUMBERHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_SERVICENUMBERHEADER) + " but got " + fmt.Sprintf("%d", serviceNumberHeader))
	}

	// Simple Field (serviceNumber)
	_serviceNumber, _serviceNumberErr := readBuffer.ReadUint16("serviceNumber", 16)
	if _serviceNumberErr != nil {
		return nil, errors.Wrap(_serviceNumberErr, "Error parsing 'serviceNumber' field")
	}
	serviceNumber := _serviceNumber

	// Const Field (listOfValuesOpeningTag)
	listOfValuesOpeningTag, _listOfValuesOpeningTagErr := readBuffer.ReadUint8("listOfValuesOpeningTag", 8)
	if _listOfValuesOpeningTagErr != nil {
		return nil, errors.Wrap(_listOfValuesOpeningTagErr, "Error parsing 'listOfValuesOpeningTag' field")
	}
	if listOfValuesOpeningTag != BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESOPENINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESOPENINGTAG) + " but got " + fmt.Sprintf("%d", listOfValuesOpeningTag))
	}

	// Array field (values)
	if pullErr := readBuffer.PullContext("values", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	values := make([]int8, 0)
	{
		_valuesLength := uint16(len) - uint16(uint16(8))
		_valuesEndPos := readBuffer.GetPos() + uint16(_valuesLength)
		for readBuffer.GetPos() < _valuesEndPos {
			_item, _err := readBuffer.ReadInt8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'values' field")
			}
			values = append(values, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("values", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (listOfValuesClosingTag)
	listOfValuesClosingTag, _listOfValuesClosingTagErr := readBuffer.ReadUint8("listOfValuesClosingTag", 8)
	if _listOfValuesClosingTagErr != nil {
		return nil, errors.Wrap(_listOfValuesClosingTagErr, "Error parsing 'listOfValuesClosingTag' field")
	}
	if listOfValuesClosingTag != BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESCLOSINGTAG {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer_LISTOFVALUESCLOSINGTAG) + " but got " + fmt.Sprintf("%d", listOfValuesClosingTag))
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		VendorId:      vendorId,
		ServiceNumber: serviceNumber,
		Values:        values,
		Parent:        &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); pushErr != nil {
			return pushErr
		}

		// Const Field (vendorIdHeader)
		_vendorIdHeaderErr := writeBuffer.WriteUint8("vendorIdHeader", 8, 0x09)
		if _vendorIdHeaderErr != nil {
			return errors.Wrap(_vendorIdHeaderErr, "Error serializing 'vendorIdHeader' field")
		}

		// Simple Field (vendorId)
		vendorId := uint8(m.VendorId)
		_vendorIdErr := writeBuffer.WriteUint8("vendorId", 8, (vendorId))
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Const Field (serviceNumberHeader)
		_serviceNumberHeaderErr := writeBuffer.WriteUint8("serviceNumberHeader", 8, 0x1A)
		if _serviceNumberHeaderErr != nil {
			return errors.Wrap(_serviceNumberHeaderErr, "Error serializing 'serviceNumberHeader' field")
		}

		// Simple Field (serviceNumber)
		serviceNumber := uint16(m.ServiceNumber)
		_serviceNumberErr := writeBuffer.WriteUint16("serviceNumber", 16, (serviceNumber))
		if _serviceNumberErr != nil {
			return errors.Wrap(_serviceNumberErr, "Error serializing 'serviceNumber' field")
		}

		// Const Field (listOfValuesOpeningTag)
		_listOfValuesOpeningTagErr := writeBuffer.WriteUint8("listOfValuesOpeningTag", 8, 0x2E)
		if _listOfValuesOpeningTagErr != nil {
			return errors.Wrap(_listOfValuesOpeningTagErr, "Error serializing 'listOfValuesOpeningTag' field")
		}

		// Array Field (values)
		if m.Values != nil {
			if pushErr := writeBuffer.PushContext("values", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Values {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'values' field")
				}
			}
			if popErr := writeBuffer.PopContext("values", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Const Field (listOfValuesClosingTag)
		_listOfValuesClosingTagErr := writeBuffer.WriteUint8("listOfValuesClosingTag", 8, 0x2F)
		if _listOfValuesClosingTagErr != nil {
			return errors.Wrap(_listOfValuesClosingTagErr, "Error serializing 'listOfValuesClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
