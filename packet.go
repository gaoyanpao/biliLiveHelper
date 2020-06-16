package biliLiveHelper

import (
	"bytes"
	"encoding/binary"
	"errors"
	json "github.com/json-iterator/go"
)

const (
	LenPacketLength = 4
	LenHeaderLength = 2
	LenProtocolVersion = 2
	LenOperation = 4
	LenSequenceID = 4
)

type BiliHeader struct {
	PacketLength    uint32
	HeaderLength    uint16
	ProtocolVersion uint16
	Operation       uint32
	SequenceID      uint32
}

type DataPacket struct {
	BiliHeader
	Body []byte
}

func NewPacket(ver uint16, operation uint32) *DataPacket {
	p := &DataPacket{}
	p.HeaderLength, p.PacketLength = 16, 16
	p.SequenceID = 1
	p.ProtocolVersion = ver
	p.Operation = operation
	return p
}

func NewHeartBeatPacket() *DataPacket {
	packet := NewPacket(1, OpHeartbeatSend)
	packet.WriteBody(HeartBeatBody)
	return packet
}

func (p *DataPacket) WriteJSONBody(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil{
		return err
	}
	p.Body = jsonData
	p.setPacketLength()
	return nil
}

func (p *DataPacket) WriteBody(data []byte)  {
	p.Body = data
	p.setPacketLength()
}

func (p *DataPacket) setPacketLength() {
	p.PacketLength = uint32(len(p.Body)) + uint32(p.HeaderLength)
}

func (p *DataPacket) Bytes() []byte {
	headerBytes := new(bytes.Buffer)
	binary.Write(headerBytes, binary.BigEndian, p.PacketLength)
	binary.Write(headerBytes, binary.BigEndian, p.HeaderLength)
	binary.Write(headerBytes, binary.BigEndian, p.ProtocolVersion)
	binary.Write(headerBytes, binary.BigEndian, p.Operation)
	binary.Write(headerBytes, binary.BigEndian, p.SequenceID)
	socketData := append(headerBytes.Bytes(), p.Body...)
	return socketData
}

func PacketFromBytes(data []byte) ([]*DataPacket, error) {
	DataLen := len(data)
	if DataLen < 16 {
		return nil, errors.New("包数据长度不足")
	}
	var packetList []*DataPacket
	offset := 0
	for offset < DataLen {
		var packetLength []byte
		var headerLength []byte
		var protocolVersion []byte
		var operation []byte
		var sequenceID []byte
		var body []byte
		packetLength, offset = readBytes(data, offset, LenPacketLength)
		headerLength, offset = readBytes(data, offset, LenHeaderLength)
		protocolVersion, offset = readBytes(data, offset, LenProtocolVersion)
		operation, offset = readBytes(data, offset, LenOperation)
		sequenceID, offset = readBytes(data, offset, LenSequenceID)
		packet := &DataPacket{}
		packet.PacketLength = binary.BigEndian.Uint32(packetLength)
		packet.HeaderLength = binary.BigEndian.Uint16(headerLength)
		packet.ProtocolVersion = binary.BigEndian.Uint16(protocolVersion)
		packet.Operation = binary.BigEndian.Uint32(operation)
		packet.SequenceID = binary.BigEndian.Uint32(sequenceID)
		body, offset = readBytes(data, offset, int(packet.PacketLength) - int(packet.HeaderLength))
		packet.Body = body
		packetList = append(packetList, packet)
	}
	return packetList, nil
}

func readBytes(b []byte, offset int, l int) ([]byte, int) {
	if l >= len(b) - offset {
		return b[offset:], len(b)
	}
	end := offset + l
	return b[offset:end], end
}