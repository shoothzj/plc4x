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
type BACnetComplexTagPropertyIdentifier struct {
	Data          []byte
	SlimmedLength uint16
	Value         BACnetPropertyIdentifier
	Parent        *BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagPropertyIdentifier interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagPropertyIdentifier) DataType() BACnetDataType {
	return BACnetDataType_BACNET_PROPERTY_IDENTIFIER
}

func (m *BACnetComplexTagPropertyIdentifier) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.TagClass = tagClass
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagPropertyIdentifier(data []byte, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagPropertyIdentifier{
		Data:   data,
		Parent: NewBACnetComplexTag(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetComplexTagPropertyIdentifier(structType interface{}) *BACnetComplexTagPropertyIdentifier {
	castFunc := func(typ interface{}) *BACnetComplexTagPropertyIdentifier {
		if casted, ok := typ.(BACnetComplexTagPropertyIdentifier); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagPropertyIdentifier); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagPropertyIdentifier(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagPropertyIdentifier(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagPropertyIdentifier) GetTypeName() string {
	return "BACnetComplexTagPropertyIdentifier"
}

func (m *BACnetComplexTagPropertyIdentifier) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagPropertyIdentifier) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetComplexTagPropertyIdentifier) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagPropertyIdentifierParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, actualLength uint32) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagPropertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_slimmedLength := actualLength
	slimmedLength := uint16(_slimmedLength)
	// Byte Array field (data)
	numberOfBytesdata := int(slimmedLength)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	// Virtual field
	_value := MapToPropertyIdentifier(data)
	value := BACnetPropertyIdentifier(_value)

	if closeErr := readBuffer.CloseContext("BACnetComplexTagPropertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagPropertyIdentifier{
		Data:          data,
		SlimmedLength: slimmedLength,
		Value:         value,
		Parent:        &BACnetComplexTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetComplexTagPropertyIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagPropertyIdentifier"); pushErr != nil {
			return pushErr
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagPropertyIdentifier"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagPropertyIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}