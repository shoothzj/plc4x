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
type BACnetConfirmedServiceRequestAtomicWriteFile struct {
	DeviceIdentifier  *BACnetTagApplicationObjectIdentifier
	OpeningTag        *BACnetComplexTagNull
	FileStartPosition *BACnetTagApplicationSignedInteger
	FileData          *BACnetTagApplicationOctetString
	ClosingTag        *BACnetComplexTagNull
	Parent            *BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestAtomicWriteFile interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestAtomicWriteFile) ServiceChoice() uint8 {
	return 0x07
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestAtomicWriteFile(deviceIdentifier *BACnetTagApplicationObjectIdentifier, openingTag *BACnetComplexTagNull, fileStartPosition *BACnetTagApplicationSignedInteger, fileData *BACnetTagApplicationOctetString, closingTag *BACnetComplexTagNull) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestAtomicWriteFile{
		DeviceIdentifier:  deviceIdentifier,
		OpeningTag:        openingTag,
		FileStartPosition: fileStartPosition,
		FileData:          fileData,
		ClosingTag:        closingTag,
		Parent:            NewBACnetConfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetConfirmedServiceRequestAtomicWriteFile(structType interface{}) *BACnetConfirmedServiceRequestAtomicWriteFile {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestAtomicWriteFile {
		if casted, ok := typ.(BACnetConfirmedServiceRequestAtomicWriteFile); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestAtomicWriteFile); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestAtomicWriteFile(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestAtomicWriteFile(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicWriteFile"
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.LengthInBits()

	// Optional Field (openingTag)
	if m.OpeningTag != nil {
		lengthInBits += (*m.OpeningTag).LengthInBits()
	}

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.LengthInBits()

	// Simple field (fileData)
	lengthInBits += m.FileData.LengthInBits()

	// Optional Field (closingTag)
	if m.ClosingTag != nil {
		lengthInBits += (*m.ClosingTag).LengthInBits()
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestAtomicWriteFileParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicWriteFile"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (deviceIdentifier)
	if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_deviceIdentifier, _deviceIdentifierErr := BACnetTagParse(readBuffer)
	if _deviceIdentifierErr != nil {
		return nil, errors.Wrap(_deviceIdentifierErr, "Error parsing 'deviceIdentifier' field")
	}
	deviceIdentifier := CastBACnetTagApplicationObjectIdentifier(_deviceIdentifier)
	if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (openingTag) (Can be skipped, if a given expression evaluates to false)
	var openingTag *BACnetComplexTagNull = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetComplexTagParse(readBuffer, uint8(0), BACnetDataType_NULL)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'openingTag' field")
		case _err == utils.ParseAssertError:
			readBuffer.SetPos(currentPos)
		default:
			openingTag = CastBACnetComplexTagNull(_val)
			if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (fileStartPosition)
	if pullErr := readBuffer.PullContext("fileStartPosition"); pullErr != nil {
		return nil, pullErr
	}
	_fileStartPosition, _fileStartPositionErr := BACnetTagParse(readBuffer)
	if _fileStartPositionErr != nil {
		return nil, errors.Wrap(_fileStartPositionErr, "Error parsing 'fileStartPosition' field")
	}
	fileStartPosition := CastBACnetTagApplicationSignedInteger(_fileStartPosition)
	if closeErr := readBuffer.CloseContext("fileStartPosition"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (fileData)
	if pullErr := readBuffer.PullContext("fileData"); pullErr != nil {
		return nil, pullErr
	}
	_fileData, _fileDataErr := BACnetTagParse(readBuffer)
	if _fileDataErr != nil {
		return nil, errors.Wrap(_fileDataErr, "Error parsing 'fileData' field")
	}
	fileData := CastBACnetTagApplicationOctetString(_fileData)
	if closeErr := readBuffer.CloseContext("fileData"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (closingTag) (Can be skipped, if a given expression evaluates to false)
	var closingTag *BACnetComplexTagNull = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetComplexTagParse(readBuffer, uint8(0), BACnetDataType_NULL)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'closingTag' field")
		case _err == utils.ParseAssertError:
			readBuffer.SetPos(currentPos)
		default:
			closingTag = CastBACnetComplexTagNull(_val)
			if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicWriteFile"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestAtomicWriteFile{
		DeviceIdentifier:  CastBACnetTagApplicationObjectIdentifier(deviceIdentifier),
		OpeningTag:        CastBACnetComplexTagNull(openingTag),
		FileStartPosition: CastBACnetTagApplicationSignedInteger(fileStartPosition),
		FileData:          CastBACnetTagApplicationOctetString(fileData),
		ClosingTag:        CastBACnetComplexTagNull(closingTag),
		Parent:            &BACnetConfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicWriteFile"); pushErr != nil {
			return pushErr
		}

		// Simple Field (deviceIdentifier)
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return pushErr
		}
		_deviceIdentifierErr := m.DeviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return popErr
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}

		// Optional Field (openingTag) (Can be skipped, if the value is null)
		var openingTag *BACnetComplexTagNull = nil
		if m.OpeningTag != nil {
			if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
				return pushErr
			}
			openingTag = m.OpeningTag
			_openingTagErr := openingTag.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
				return popErr
			}
			if _openingTagErr != nil {
				return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
			}
		}

		// Simple Field (fileStartPosition)
		if pushErr := writeBuffer.PushContext("fileStartPosition"); pushErr != nil {
			return pushErr
		}
		_fileStartPositionErr := m.FileStartPosition.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fileStartPosition"); popErr != nil {
			return popErr
		}
		if _fileStartPositionErr != nil {
			return errors.Wrap(_fileStartPositionErr, "Error serializing 'fileStartPosition' field")
		}

		// Simple Field (fileData)
		if pushErr := writeBuffer.PushContext("fileData"); pushErr != nil {
			return pushErr
		}
		_fileDataErr := m.FileData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fileData"); popErr != nil {
			return popErr
		}
		if _fileDataErr != nil {
			return errors.Wrap(_fileDataErr, "Error serializing 'fileData' field")
		}

		// Optional Field (closingTag) (Can be skipped, if the value is null)
		var closingTag *BACnetComplexTagNull = nil
		if m.ClosingTag != nil {
			if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
				return pushErr
			}
			closingTag = m.ClosingTag
			_closingTagErr := closingTag.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
				return popErr
			}
			if _closingTagErr != nil {
				return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicWriteFile"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestAtomicWriteFile) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
