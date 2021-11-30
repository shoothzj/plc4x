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
type DateAndTime struct {
	Year    uint8
	Month   uint8
	Day     uint8
	Hour    uint8
	Minutes uint8
	Seconds uint8
	Msec    uint16
	Dow     uint8
}

// The corresponding interface
type IDateAndTime interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewDateAndTime(year uint8, month uint8, day uint8, hour uint8, minutes uint8, seconds uint8, msec uint16, dow uint8) *DateAndTime {
	return &DateAndTime{Year: year, Month: month, Day: day, Hour: hour, Minutes: minutes, Seconds: seconds, Msec: msec, Dow: dow}
}

func CastDateAndTime(structType interface{}) *DateAndTime {
	castFunc := func(typ interface{}) *DateAndTime {
		if casted, ok := typ.(DateAndTime); ok {
			return &casted
		}
		if casted, ok := typ.(*DateAndTime); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DateAndTime) GetTypeName() string {
	return "DateAndTime"
}

func (m *DateAndTime) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DateAndTime) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (year)
	lengthInBits += uint16(int32(1) * 8)

	// Manual Field (month)
	lengthInBits += uint16(int32(1) * 8)

	// Manual Field (day)
	lengthInBits += uint16(int32(1) * 8)

	// Manual Field (hour)
	lengthInBits += uint16(int32(1) * 8)

	// Manual Field (minutes)
	lengthInBits += uint16(int32(1) * 8)

	// Manual Field (seconds)
	lengthInBits += uint16(int32(1) * 8)

	// Manual Field (msec)
	lengthInBits += uint16(int32(2) * 8)

	// Simple field (dow)
	lengthInBits += 4

	return lengthInBits
}

func (m *DateAndTime) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DateAndTimeParse(readBuffer utils.ReadBuffer) (*DateAndTime, error) {
	if pullErr := readBuffer.PullContext("DateAndTime"); pullErr != nil {
		return nil, pullErr
	}

	// Manual Field (year)
	if pullErr := readBuffer.PullContext("year"); pullErr != nil {
		return nil, pullErr
	}
	year, _yearErr := BcdToInt(readBuffer)
	if _yearErr != nil {
		return nil, errors.Wrap(_yearErr, "Error parsing 'year' field")
	}
	if closeErr := readBuffer.CloseContext("year"); closeErr != nil {
		return nil, closeErr
	}

	// Manual Field (month)
	if pullErr := readBuffer.PullContext("month"); pullErr != nil {
		return nil, pullErr
	}
	month, _monthErr := BcdToInt(readBuffer)
	if _monthErr != nil {
		return nil, errors.Wrap(_monthErr, "Error parsing 'month' field")
	}
	if closeErr := readBuffer.CloseContext("month"); closeErr != nil {
		return nil, closeErr
	}

	// Manual Field (day)
	if pullErr := readBuffer.PullContext("day"); pullErr != nil {
		return nil, pullErr
	}
	day, _dayErr := BcdToInt(readBuffer)
	if _dayErr != nil {
		return nil, errors.Wrap(_dayErr, "Error parsing 'day' field")
	}
	if closeErr := readBuffer.CloseContext("day"); closeErr != nil {
		return nil, closeErr
	}

	// Manual Field (hour)
	if pullErr := readBuffer.PullContext("hour"); pullErr != nil {
		return nil, pullErr
	}
	hour, _hourErr := BcdToInt(readBuffer)
	if _hourErr != nil {
		return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field")
	}
	if closeErr := readBuffer.CloseContext("hour"); closeErr != nil {
		return nil, closeErr
	}

	// Manual Field (minutes)
	if pullErr := readBuffer.PullContext("minutes"); pullErr != nil {
		return nil, pullErr
	}
	minutes, _minutesErr := BcdToInt(readBuffer)
	if _minutesErr != nil {
		return nil, errors.Wrap(_minutesErr, "Error parsing 'minutes' field")
	}
	if closeErr := readBuffer.CloseContext("minutes"); closeErr != nil {
		return nil, closeErr
	}

	// Manual Field (seconds)
	if pullErr := readBuffer.PullContext("seconds"); pullErr != nil {
		return nil, pullErr
	}
	seconds, _secondsErr := BcdToInt(readBuffer)
	if _secondsErr != nil {
		return nil, errors.Wrap(_secondsErr, "Error parsing 'seconds' field")
	}
	if closeErr := readBuffer.CloseContext("seconds"); closeErr != nil {
		return nil, closeErr
	}

	// Manual Field (msec)
	if pullErr := readBuffer.PullContext("msec"); pullErr != nil {
		return nil, pullErr
	}
	msec, _msecErr := S7msecToInt(readBuffer)
	if _msecErr != nil {
		return nil, errors.Wrap(_msecErr, "Error parsing 'msec' field")
	}
	if closeErr := readBuffer.CloseContext("msec"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (dow)
	_dow, _dowErr := readBuffer.ReadUint8("dow", 4)
	if _dowErr != nil {
		return nil, errors.Wrap(_dowErr, "Error parsing 'dow' field")
	}
	dow := _dow

	if closeErr := readBuffer.CloseContext("DateAndTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewDateAndTime(year, month, day, hour, minutes, seconds, msec, dow), nil
}

func (m *DateAndTime) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("DateAndTime"); pushErr != nil {
		return pushErr
	}

	// Manual Field (year)
	if pushErr := writeBuffer.PushContext("dow"); pushErr != nil {
		return pushErr
	}
	_yearErr := ByteToBcd(writeBuffer, m.Year)
	if popErr := writeBuffer.PopContext("year"); popErr != nil {
		return popErr
	}
	if _yearErr != nil {
		return errors.Wrap(_yearErr, "Error serializing 'year' field")
	}

	// Manual Field (month)
	if pushErr := writeBuffer.PushContext("dow"); pushErr != nil {
		return pushErr
	}
	_monthErr := ByteToBcd(writeBuffer, m.Month)
	if popErr := writeBuffer.PopContext("month"); popErr != nil {
		return popErr
	}
	if _monthErr != nil {
		return errors.Wrap(_monthErr, "Error serializing 'month' field")
	}

	// Manual Field (day)
	if pushErr := writeBuffer.PushContext("dow"); pushErr != nil {
		return pushErr
	}
	_dayErr := ByteToBcd(writeBuffer, m.Day)
	if popErr := writeBuffer.PopContext("day"); popErr != nil {
		return popErr
	}
	if _dayErr != nil {
		return errors.Wrap(_dayErr, "Error serializing 'day' field")
	}

	// Manual Field (hour)
	if pushErr := writeBuffer.PushContext("dow"); pushErr != nil {
		return pushErr
	}
	_hourErr := ByteToBcd(writeBuffer, m.Hour)
	if popErr := writeBuffer.PopContext("hour"); popErr != nil {
		return popErr
	}
	if _hourErr != nil {
		return errors.Wrap(_hourErr, "Error serializing 'hour' field")
	}

	// Manual Field (minutes)
	if pushErr := writeBuffer.PushContext("dow"); pushErr != nil {
		return pushErr
	}
	_minutesErr := ByteToBcd(writeBuffer, m.Minutes)
	if popErr := writeBuffer.PopContext("minutes"); popErr != nil {
		return popErr
	}
	if _minutesErr != nil {
		return errors.Wrap(_minutesErr, "Error serializing 'minutes' field")
	}

	// Manual Field (seconds)
	if pushErr := writeBuffer.PushContext("dow"); pushErr != nil {
		return pushErr
	}
	_secondsErr := ByteToBcd(writeBuffer, m.Seconds)
	if popErr := writeBuffer.PopContext("seconds"); popErr != nil {
		return popErr
	}
	if _secondsErr != nil {
		return errors.Wrap(_secondsErr, "Error serializing 'seconds' field")
	}

	// Manual Field (msec)
	if pushErr := writeBuffer.PushContext("dow"); pushErr != nil {
		return pushErr
	}
	_msecErr := IntToS7msec(writeBuffer, m.Msec)
	if popErr := writeBuffer.PopContext("msec"); popErr != nil {
		return popErr
	}
	if _msecErr != nil {
		return errors.Wrap(_msecErr, "Error serializing 'msec' field")
	}

	// Simple Field (dow)
	dow := uint8(m.Dow)
	_dowErr := writeBuffer.WriteUint8("dow", 4, (dow))
	if _dowErr != nil {
		return errors.Wrap(_dowErr, "Error serializing 'dow' field")
	}

	if popErr := writeBuffer.PopContext("DateAndTime"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DateAndTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
