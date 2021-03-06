/**
 *     SpecterGO  Copyright (C) 2018  SpecterTeam
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package protocol

import (
	"errors"

	"github.com/SpecterTeam/go-raknet"
	"github.com/SpecterTeam/go-raknet/binary"
)

const (
	EPacketMinBufferLen                         = 3
	EPacketBitFlagLen                           = 1
	EPacketPayloadLengthLen                     = 2
	EPacketMessageIndexLen                      = 3
	EPacketOrderIndexAndOrderChannelLen         = 4
	EPacketSplitCountAndSplitIDAndSplitIndexLen = 10
)

const (
	ReliabilityPosition = 5 //
)

const (
	FlagReliability = 224 // 0b11100000
	FlagSplit       = 16  // 0b00010000
)

type EncapsulatedPacket struct {
	Buf *binary.RaknetStream

	Reliability  raknet.Reliability
	Split        bool
	MessageIndex binary.Triad
	OrderIndex   binary.Triad
	OrderChannel byte
	SplitCount   int32
	SplitID      uint16
	SplitIndex   int32

	Payload []byte
}

func (epk *EncapsulatedPacket) Encode() error {
	if epk.Buf == nil {
		epk.Buf = binary.NewStream()
	}

	flags := epk.Reliability.ToBinary() << ReliabilityPosition
	if epk.Split {
		flags |= FlagSplit
	}

	err := epk.Buf.PutByte(flags)
	if err != nil {
		return err
	}

	err = epk.Buf.PutShort(uint16(len(epk.Payload) << 3))
	if err != nil {
		return err
	}

	if epk.Reliability.IsReliable() {
		err = epk.Buf.PutLTriad(epk.MessageIndex)
		if err != nil {
			return err
		}
	}

	if epk.Reliability.IsOrdered() || epk.Reliability.IsSequenced() {
		err = epk.Buf.PutLTriad(epk.OrderIndex)
		if err != nil {
			return err
		}

		err = epk.Buf.PutByte(epk.OrderChannel)
		if err != nil {
			return err
		}
	}

	if epk.Split {
		err = epk.Buf.PutInt(epk.SplitCount)
		if err != nil {
			return err
		}

		err = epk.Buf.PutShort(epk.SplitID)
		if err != nil {
			return err
		}

		err = epk.Buf.PutInt(epk.SplitIndex)
		if err != nil {
			return err
		}
	}

	err = epk.Buf.Put(epk.Payload)
	if err != nil {
		return err
	}

	return nil
}

func (epk *EncapsulatedPacket) Decode() error {
	if epk.Buf == nil {
		return errors.New("no sets buffer")
	}

	flags, err := epk.Buf.Byte()
	if err != nil {
		return err
	}

	epk.Reliability = raknet.ReliabilityBinary(flags >> ReliabilityPosition)
	epk.Split = (flags & FlagSplit) > 0

	payloadLen, err := epk.Buf.Short()
	if err != nil {
		return err
	}

	length := int(payloadLen / 8)

	if epk.Reliability.IsReliable() {
		epk.MessageIndex, err = epk.Buf.LTriad()
		if err != nil {
			return err
		}
	}

	if epk.Reliability.IsOrdered() || epk.Reliability.IsSequenced() {
		epk.OrderIndex, err = epk.Buf.LTriad()
		if err != nil {
			return err
		}

		epk.OrderChannel, err = epk.Buf.Byte()
		if err != nil {
			return err
		}
	}

	if epk.Split {
		epk.SplitCount, err = epk.Buf.Int()
		if err != nil {
			return err
		}

		epk.SplitID, err = epk.Buf.Short()
		if err != nil {
			return err
		}

		epk.SplitIndex, err = epk.Buf.Int()
		if err != nil {
			return err
		}
	}

	epk.Payload = epk.Buf.Get(length)

	return nil
}

func (epk *EncapsulatedPacket) CalcSize() int {
	return CalcEPacketSize(epk.Reliability, epk.Split, epk.Payload)
}

func CalcEPacketSize(reliability raknet.Reliability, split bool, payload []byte) int {
	var size int
	size += EPacketBitFlagLen
	size += EPacketPayloadLengthLen

	if reliability.IsReliable() {
		size += EPacketMessageIndexLen
	}

	if reliability.IsOrdered() || reliability.IsSequenced() {
		size += EPacketOrderIndexAndOrderChannelLen
	}

	if split {
		size += EPacketSplitCountAndSplitIDAndSplitIndexLen
	}

	size += len(payload)

	return size
}

func NewCustomPacket(id byte) raknet.Packet {
	return &CustomPacket{
		id: id,
	}
}

type CustomPacket struct {
	BasePacket

	id byte

	Index    binary.Triad
	Messages []*EncapsulatedPacket
}

func (pk *CustomPacket) ID() byte {
	return pk.id
}

func (pk *CustomPacket) Encode() error {
	err := pk.BasePacket.Encode(pk)
	if err != nil {
		return err
	}

	err = pk.PutLTriad(pk.Index)
	if err != nil {
		return err
	}

	for _, epk := range pk.Messages {
		epk.Buf = &pk.RaknetStream

		err = epk.Encode()
		if err != nil {
			return err
		}
	}

	return nil
}

func (pk *CustomPacket) Decode() error {
	err := pk.BasePacket.Decode(pk)
	if err != nil {
		return err
	}

	pk.Index, err = pk.LTriad()
	if err != nil {
		return err
	}

	for pk.Len() >= EPacketMinBufferLen {
		epk := &EncapsulatedPacket{
			Buf: &pk.RaknetStream,
		}

		err = epk.Decode()
		if err != nil {
			return err
		}

		pk.Messages = append(pk.Messages, epk)
	}

	return nil
}

func (pk *CustomPacket) CalcSize() int {
	size := 0
	for _, epk := range pk.Messages {
		size += epk.CalcSize()
	}

	return size
}

func (pk *CustomPacket) New() raknet.Packet {
	return NewCustomPacket(pk.id)
}
