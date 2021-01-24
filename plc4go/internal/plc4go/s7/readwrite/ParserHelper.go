//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package readwrite

import (
    "errors"
    "github.com/apache/plc4x/plc4go/internal/plc4go/s7/readwrite/model"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type S7ParserHelper struct {
}

func (m S7ParserHelper) Parse(typeName string, arguments []string, io *utils.ReadBuffer) (spi.Message, error) {
    switch typeName {
    case "SzlId":
        return model.SzlIdParse(io)
    case "S7Message":
        return model.S7MessageParse(io)
    case "S7VarPayloadStatusItem":
        return model.S7VarPayloadStatusItemParse(io)
    case "S7Parameter":
        messageType, err := utils.StrToUint8(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.S7ParameterParse(io, messageType)
    case "SzlDataTreeItem":
        return model.SzlDataTreeItemParse(io)
    case "COTPPacket":
        cotpLen, err := utils.StrToUint16(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.COTPPacketParse(io, cotpLen)
    case "S7PayloadUserDataItem":
        cpuFunctionType, err := utils.StrToUint8(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.S7PayloadUserDataItemParse(io, cpuFunctionType)
    case "COTPParameter":
        rest, err := utils.StrToUint8(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.COTPParameterParse(io, rest)
    case "TPKTPacket":
        return model.TPKTPacketParse(io)
    case "S7Payload":
        messageType, err := utils.StrToUint8(arguments[0])
        if err != nil {
            return nil, err
        }
        var parameter IS7Parameter
        return model.S7PayloadParse(io, messageType, parameter)
    case "S7VarRequestParameterItem":
        return model.S7VarRequestParameterItemParse(io)
    case "S7VarPayloadDataItem":
        lastItem, err := utils.StrToBool(arguments[0])
        if err != nil {
            return nil, err
        }
        return model.S7VarPayloadDataItemParse(io, lastItem)
    case "S7Address":
        return model.S7AddressParse(io)
    case "S7ParameterUserDataItem":
        return model.S7ParameterUserDataItemParse(io)
    }
    return nil, errors.New("Unsupported type " + typeName)
}
