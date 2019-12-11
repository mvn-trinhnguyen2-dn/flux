// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type IntegerLiteral struct {
	_tab flatbuffers.Table
}

func GetRootAsIntegerLiteral(buf []byte, offset flatbuffers.UOffsetT) *IntegerLiteral {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IntegerLiteral{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *IntegerLiteral) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IntegerLiteral) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IntegerLiteral) Loc(obj *SourceLocation) *SourceLocation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SourceLocation)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *IntegerLiteral) Value() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *IntegerLiteral) MutateValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *IntegerLiteral) TypType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *IntegerLiteral) MutateTypType(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *IntegerLiteral) Typ(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func IntegerLiteralStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func IntegerLiteralAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func IntegerLiteralAddValue(builder *flatbuffers.Builder, value int64) {
	builder.PrependInt64Slot(1, value, 0)
}
func IntegerLiteralAddTypType(builder *flatbuffers.Builder, typType byte) {
	builder.PrependByteSlot(2, typType, 0)
}
func IntegerLiteralAddTyp(builder *flatbuffers.Builder, typ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(typ), 0)
}
func IntegerLiteralEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}