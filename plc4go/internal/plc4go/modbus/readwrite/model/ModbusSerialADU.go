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
type ModbusSerialADU struct {
	TransactionId uint16
	Length        uint16
	Address       uint8
	Pdu           *ModbusPDU
}

// The corresponding interface
type IModbusSerialADU interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewModbusSerialADU(transactionId uint16, length uint16, address uint8, pdu *ModbusPDU) *ModbusSerialADU {
	return &ModbusSerialADU{TransactionId: transactionId, Length: length, Address: address, Pdu: pdu}
}

func CastModbusSerialADU(structType interface{}) *ModbusSerialADU {
	castFunc := func(typ interface{}) *ModbusSerialADU {
		if casted, ok := typ.(ModbusSerialADU); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusSerialADU); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusSerialADU) GetTypeName() string {
	return "ModbusSerialADU"
}

func (m *ModbusSerialADU) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ModbusSerialADU) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (transactionId)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (length)
	lengthInBits += 16

	// Simple field (address)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.LengthInBits()

	return lengthInBits
}

func (m *ModbusSerialADU) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ModbusSerialADUParse(readBuffer utils.ReadBuffer, response bool) (*ModbusSerialADU, error) {
	if pullErr := readBuffer.PullContext("ModbusSerialADU"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (transactionId)
	_transactionId, _transactionIdErr := readBuffer.ReadUint16("transactionId", 16)
	if _transactionIdErr != nil {
		return nil, errors.Wrap(_transactionIdErr, "Error parsing 'transactionId' field")
	}
	transactionId := _transactionId

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (length)
	_length, _lengthErr := readBuffer.ReadUint16("length", 16)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}
	length := _length

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint8("address", 8)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address

	// Simple Field (pdu)
	if pullErr := readBuffer.PullContext("pdu"); pullErr != nil {
		return nil, pullErr
	}
	_pdu, _pduErr := ModbusPDUParse(readBuffer, response)
	if _pduErr != nil {
		return nil, errors.Wrap(_pduErr, "Error parsing 'pdu' field")
	}
	pdu := CastModbusPDU(_pdu)
	if closeErr := readBuffer.CloseContext("pdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusSerialADU"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewModbusSerialADU(transactionId, length, address, pdu), nil
}

func (m *ModbusSerialADU) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("ModbusSerialADU"); pushErr != nil {
		return pushErr
	}

	// Simple Field (transactionId)
	transactionId := uint16(m.TransactionId)
	_transactionIdErr := writeBuffer.WriteUint16("transactionId", 16, (transactionId))
	if _transactionIdErr != nil {
		return errors.Wrap(_transactionIdErr, "Error serializing 'transactionId' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (length)
	length := uint16(m.Length)
	_lengthErr := writeBuffer.WriteUint16("length", 16, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (address)
	address := uint8(m.Address)
	_addressErr := writeBuffer.WriteUint8("address", 8, (address))
	if _addressErr != nil {
		return errors.Wrap(_addressErr, "Error serializing 'address' field")
	}

	// Simple Field (pdu)
	if pushErr := writeBuffer.PushContext("pdu"); pushErr != nil {
		return pushErr
	}
	_pduErr := m.Pdu.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("pdu"); popErr != nil {
		return popErr
	}
	if _pduErr != nil {
		return errors.Wrap(_pduErr, "Error serializing 'pdu' field")
	}

	if popErr := writeBuffer.PopContext("ModbusSerialADU"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ModbusSerialADU) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
