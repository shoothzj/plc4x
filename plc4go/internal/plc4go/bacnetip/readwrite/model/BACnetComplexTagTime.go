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
type BACnetComplexTagTime struct {
	Hour                 int8
	Minute               int8
	Second               int8
	Fractional           int8
	Wildcard             int8
	HourIsWildcard       bool
	MinuteIsWildcard     bool
	SecondIsWildcard     bool
	FractionalIsWildcard bool
	Parent               *BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagTime interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagTime) DataType() BACnetDataType {
	return BACnetDataType_TIME
}

func (m *BACnetComplexTagTime) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.TagClass = tagClass
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagTime(hour int8, minute int8, second int8, fractional int8, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagTime{
		Hour:       hour,
		Minute:     minute,
		Second:     second,
		Fractional: fractional,
		Parent:     NewBACnetComplexTag(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetComplexTagTime(structType interface{}) *BACnetComplexTagTime {
	castFunc := func(typ interface{}) *BACnetComplexTagTime {
		if casted, ok := typ.(BACnetComplexTagTime); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagTime); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagTime(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagTime(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagTime) GetTypeName() string {
	return "BACnetComplexTagTime"
}

func (m *BACnetComplexTagTime) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagTime) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Simple field (hour)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (minute)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (second)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (fractional)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetComplexTagTime) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagTimeParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagTime"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_wildcard := 0xFF
	wildcard := int8(_wildcard)

	// Simple Field (hour)
	_hour, _hourErr := readBuffer.ReadInt8("hour", 8)
	if _hourErr != nil {
		return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field")
	}
	hour := _hour

	// Virtual field
	_hourIsWildcard := bool((hour) == (wildcard))
	hourIsWildcard := bool(_hourIsWildcard)

	// Simple Field (minute)
	_minute, _minuteErr := readBuffer.ReadInt8("minute", 8)
	if _minuteErr != nil {
		return nil, errors.Wrap(_minuteErr, "Error parsing 'minute' field")
	}
	minute := _minute

	// Virtual field
	_minuteIsWildcard := bool((minute) == (wildcard))
	minuteIsWildcard := bool(_minuteIsWildcard)

	// Simple Field (second)
	_second, _secondErr := readBuffer.ReadInt8("second", 8)
	if _secondErr != nil {
		return nil, errors.Wrap(_secondErr, "Error parsing 'second' field")
	}
	second := _second

	// Virtual field
	_secondIsWildcard := bool((second) == (wildcard))
	secondIsWildcard := bool(_secondIsWildcard)

	// Simple Field (fractional)
	_fractional, _fractionalErr := readBuffer.ReadInt8("fractional", 8)
	if _fractionalErr != nil {
		return nil, errors.Wrap(_fractionalErr, "Error parsing 'fractional' field")
	}
	fractional := _fractional

	// Virtual field
	_fractionalIsWildcard := bool((fractional) == (wildcard))
	fractionalIsWildcard := bool(_fractionalIsWildcard)

	if closeErr := readBuffer.CloseContext("BACnetComplexTagTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagTime{
		Hour:                 hour,
		Minute:               minute,
		Second:               second,
		Fractional:           fractional,
		Wildcard:             wildcard,
		HourIsWildcard:       hourIsWildcard,
		MinuteIsWildcard:     minuteIsWildcard,
		SecondIsWildcard:     secondIsWildcard,
		FractionalIsWildcard: fractionalIsWildcard,
		Parent:               &BACnetComplexTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetComplexTagTime) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (hour)
		hour := int8(m.Hour)
		_hourErr := writeBuffer.WriteInt8("hour", 8, (hour))
		if _hourErr != nil {
			return errors.Wrap(_hourErr, "Error serializing 'hour' field")
		}

		// Simple Field (minute)
		minute := int8(m.Minute)
		_minuteErr := writeBuffer.WriteInt8("minute", 8, (minute))
		if _minuteErr != nil {
			return errors.Wrap(_minuteErr, "Error serializing 'minute' field")
		}

		// Simple Field (second)
		second := int8(m.Second)
		_secondErr := writeBuffer.WriteInt8("second", 8, (second))
		if _secondErr != nil {
			return errors.Wrap(_secondErr, "Error serializing 'second' field")
		}

		// Simple Field (fractional)
		fractional := int8(m.Fractional)
		_fractionalErr := writeBuffer.WriteInt8("fractional", 8, (fractional))
		if _fractionalErr != nil {
			return errors.Wrap(_fractionalErr, "Error serializing 'fractional' field")
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
