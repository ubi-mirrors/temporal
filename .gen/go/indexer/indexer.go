// The MIT License (MIT)
// 
// Copyright (c) 2017 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by thriftrw v1.20.0. DO NOT EDIT.
// @generated

package indexer

import (
	bytes "bytes"
	base64 "encoding/base64"
	json "encoding/json"
	fmt "fmt"
	shared "github.com/uber/cadence/.gen/go/shared"
	multierr "go.uber.org/multierr"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	math "math"
	strconv "strconv"
	strings "strings"
)

type Field struct {
	Type       *FieldType `json:"type,omitempty"`
	StringData *string    `json:"stringData,omitempty"`
	IntData    *int64     `json:"intData,omitempty"`
	BoolData   *bool      `json:"boolData,omitempty"`
	BinaryData []byte     `json:"binaryData,omitempty"`
}

// ToWire translates a Field struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Field) ToWire() (wire.Value, error) {
	var (
		fields [5]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Type != nil {
		w, err = v.Type.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}
	if v.StringData != nil {
		w, err = wire.NewValueString(*(v.StringData)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 20, Value: w}
		i++
	}
	if v.IntData != nil {
		w, err = wire.NewValueI64(*(v.IntData)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 30, Value: w}
		i++
	}
	if v.BoolData != nil {
		w, err = wire.NewValueBool(*(v.BoolData)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 40, Value: w}
		i++
	}
	if v.BinaryData != nil {
		w, err = wire.NewValueBinary(v.BinaryData), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 50, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _FieldType_Read(w wire.Value) (FieldType, error) {
	var v FieldType
	err := v.FromWire(w)
	return v, err
}

// FromWire deserializes a Field struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Field struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Field
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Field) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 10:
			if field.Value.Type() == wire.TI32 {
				var x FieldType
				x, err = _FieldType_Read(field.Value)
				v.Type = &x
				if err != nil {
					return err
				}

			}
		case 20:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.StringData = &x
				if err != nil {
					return err
				}

			}
		case 30:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.IntData = &x
				if err != nil {
					return err
				}

			}
		case 40:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.BoolData = &x
				if err != nil {
					return err
				}

			}
		case 50:
			if field.Value.Type() == wire.TBinary {
				v.BinaryData, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Field
// struct.
func (v *Field) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [5]string
	i := 0
	if v.Type != nil {
		fields[i] = fmt.Sprintf("Type: %v", *(v.Type))
		i++
	}
	if v.StringData != nil {
		fields[i] = fmt.Sprintf("StringData: %v", *(v.StringData))
		i++
	}
	if v.IntData != nil {
		fields[i] = fmt.Sprintf("IntData: %v", *(v.IntData))
		i++
	}
	if v.BoolData != nil {
		fields[i] = fmt.Sprintf("BoolData: %v", *(v.BoolData))
		i++
	}
	if v.BinaryData != nil {
		fields[i] = fmt.Sprintf("BinaryData: %v", v.BinaryData)
		i++
	}

	return fmt.Sprintf("Field{%v}", strings.Join(fields[:i], ", "))
}

func _FieldType_EqualsPtr(lhs, rhs *FieldType) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _I64_EqualsPtr(lhs, rhs *int64) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _Bool_EqualsPtr(lhs, rhs *bool) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Field match the
// provided Field.
//
// This function performs a deep comparison.
func (v *Field) Equals(rhs *Field) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_FieldType_EqualsPtr(v.Type, rhs.Type) {
		return false
	}
	if !_String_EqualsPtr(v.StringData, rhs.StringData) {
		return false
	}
	if !_I64_EqualsPtr(v.IntData, rhs.IntData) {
		return false
	}
	if !_Bool_EqualsPtr(v.BoolData, rhs.BoolData) {
		return false
	}
	if !((v.BinaryData == nil && rhs.BinaryData == nil) || (v.BinaryData != nil && rhs.BinaryData != nil && bytes.Equal(v.BinaryData, rhs.BinaryData))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Field.
func (v *Field) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Type != nil {
		err = multierr.Append(err, enc.AddObject("type", *v.Type))
	}
	if v.StringData != nil {
		enc.AddString("stringData", *v.StringData)
	}
	if v.IntData != nil {
		enc.AddInt64("intData", *v.IntData)
	}
	if v.BoolData != nil {
		enc.AddBool("boolData", *v.BoolData)
	}
	if v.BinaryData != nil {
		enc.AddString("binaryData", base64.StdEncoding.EncodeToString(v.BinaryData))
	}
	return err
}

// GetType returns the value of Type if it is set or its
// zero value if it is unset.
func (v *Field) GetType() (o FieldType) {
	if v != nil && v.Type != nil {
		return *v.Type
	}

	return
}

// IsSetType returns true if Type is not nil.
func (v *Field) IsSetType() bool {
	return v != nil && v.Type != nil
}

// GetStringData returns the value of StringData if it is set or its
// zero value if it is unset.
func (v *Field) GetStringData() (o string) {
	if v != nil && v.StringData != nil {
		return *v.StringData
	}

	return
}

// IsSetStringData returns true if StringData is not nil.
func (v *Field) IsSetStringData() bool {
	return v != nil && v.StringData != nil
}

// GetIntData returns the value of IntData if it is set or its
// zero value if it is unset.
func (v *Field) GetIntData() (o int64) {
	if v != nil && v.IntData != nil {
		return *v.IntData
	}

	return
}

// IsSetIntData returns true if IntData is not nil.
func (v *Field) IsSetIntData() bool {
	return v != nil && v.IntData != nil
}

// GetBoolData returns the value of BoolData if it is set or its
// zero value if it is unset.
func (v *Field) GetBoolData() (o bool) {
	if v != nil && v.BoolData != nil {
		return *v.BoolData
	}

	return
}

// IsSetBoolData returns true if BoolData is not nil.
func (v *Field) IsSetBoolData() bool {
	return v != nil && v.BoolData != nil
}

// GetBinaryData returns the value of BinaryData if it is set or its
// zero value if it is unset.
func (v *Field) GetBinaryData() (o []byte) {
	if v != nil && v.BinaryData != nil {
		return v.BinaryData
	}

	return
}

// IsSetBinaryData returns true if BinaryData is not nil.
func (v *Field) IsSetBinaryData() bool {
	return v != nil && v.BinaryData != nil
}

type FieldType int32

const (
	FieldTypeString FieldType = 0
	FieldTypeInt    FieldType = 1
	FieldTypeBool   FieldType = 2
	FieldTypeBinary FieldType = 3
)

// FieldType_Values returns all recognized values of FieldType.
func FieldType_Values() []FieldType {
	return []FieldType{
		FieldTypeString,
		FieldTypeInt,
		FieldTypeBool,
		FieldTypeBinary,
	}
}

// UnmarshalText tries to decode FieldType from a byte slice
// containing its name.
//
//   var v FieldType
//   err := v.UnmarshalText([]byte("String"))
func (v *FieldType) UnmarshalText(value []byte) error {
	switch s := string(value); s {
	case "String":
		*v = FieldTypeString
		return nil
	case "Int":
		*v = FieldTypeInt
		return nil
	case "Bool":
		*v = FieldTypeBool
		return nil
	case "Binary":
		*v = FieldTypeBinary
		return nil
	default:
		val, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return fmt.Errorf("unknown enum value %q for %q: %v", s, "FieldType", err)
		}
		*v = FieldType(val)
		return nil
	}
}

// MarshalText encodes FieldType to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v FieldType) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("String"), nil
	case 1:
		return []byte("Int"), nil
	case 2:
		return []byte("Bool"), nil
	case 3:
		return []byte("Binary"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of FieldType.
// Enums are logged as objects, where the value is logged with key "value", and
// if this value's name is known, the name is logged with key "name".
func (v FieldType) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "String")
	case 1:
		enc.AddString("name", "Int")
	case 2:
		enc.AddString("name", "Bool")
	case 3:
		enc.AddString("name", "Binary")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v FieldType) Ptr() *FieldType {
	return &v
}

// ToWire translates FieldType into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v FieldType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes FieldType from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return FieldType(0), err
//   }
//
//   var v FieldType
//   if err := v.FromWire(x); err != nil {
//     return FieldType(0), err
//   }
//   return v, nil
func (v *FieldType) FromWire(w wire.Value) error {
	*v = (FieldType)(w.GetI32())
	return nil
}

// String returns a readable string representation of FieldType.
func (v FieldType) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "String"
	case 1:
		return "Int"
	case 2:
		return "Bool"
	case 3:
		return "Binary"
	}
	return fmt.Sprintf("FieldType(%d)", w)
}

// Equals returns true if this FieldType value matches the provided
// value.
func (v FieldType) Equals(rhs FieldType) bool {
	return v == rhs
}

// MarshalJSON serializes FieldType into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v FieldType) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"String\""), nil
	case 1:
		return ([]byte)("\"Int\""), nil
	case 2:
		return ([]byte)("\"Bool\""), nil
	case 3:
		return ([]byte)("\"Binary\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode FieldType from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *FieldType) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "FieldType")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "FieldType")
		}
		*v = (FieldType)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "FieldType")
	}
}

type Message struct {
	MessageType *MessageType      `json:"messageType,omitempty"`
	DomainID    *string           `json:"domainID,omitempty"`
	WorkflowID  *string           `json:"workflowID,omitempty"`
	RunID       *string           `json:"runID,omitempty"`
	Version     *int64            `json:"version,omitempty"`
	Fields      map[string]*Field `json:"fields,omitempty"`
}

type _Map_String_Field_MapItemList map[string]*Field

func (m _Map_String_Field_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}

		vw, err := v.ToWire()
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_Field_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_Field_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_Field_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_String_Field_MapItemList) Close() {}

// ToWire translates a Message struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Message) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.MessageType != nil {
		w, err = v.MessageType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 10, Value: w}
		i++
	}
	if v.DomainID != nil {
		w, err = wire.NewValueString(*(v.DomainID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 20, Value: w}
		i++
	}
	if v.WorkflowID != nil {
		w, err = wire.NewValueString(*(v.WorkflowID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 30, Value: w}
		i++
	}
	if v.RunID != nil {
		w, err = wire.NewValueString(*(v.RunID)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 40, Value: w}
		i++
	}
	if v.Version != nil {
		w, err = wire.NewValueI64(*(v.Version)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 50, Value: w}
		i++
	}
	if v.Fields != nil {
		w, err = wire.NewValueMap(_Map_String_Field_MapItemList(v.Fields)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 60, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _MessageType_Read(w wire.Value) (MessageType, error) {
	var v MessageType
	err := v.FromWire(w)
	return v, err
}

func _Field_Read(w wire.Value) (*Field, error) {
	var v Field
	err := v.FromWire(w)
	return &v, err
}

func _Map_String_Field_Read(m wire.MapItemList) (map[string]*Field, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}

	if m.ValueType() != wire.TStruct {
		return nil, nil
	}

	o := make(map[string]*Field, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}

		v, err := _Field_Read(x.Value)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

// FromWire deserializes a Message struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Message struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Message
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Message) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 10:
			if field.Value.Type() == wire.TI32 {
				var x MessageType
				x, err = _MessageType_Read(field.Value)
				v.MessageType = &x
				if err != nil {
					return err
				}

			}
		case 20:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.DomainID = &x
				if err != nil {
					return err
				}

			}
		case 30:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.WorkflowID = &x
				if err != nil {
					return err
				}

			}
		case 40:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.RunID = &x
				if err != nil {
					return err
				}

			}
		case 50:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.Version = &x
				if err != nil {
					return err
				}

			}
		case 60:
			if field.Value.Type() == wire.TMap {
				v.Fields, err = _Map_String_Field_Read(field.Value.GetMap())
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Message
// struct.
func (v *Message) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [6]string
	i := 0
	if v.MessageType != nil {
		fields[i] = fmt.Sprintf("MessageType: %v", *(v.MessageType))
		i++
	}
	if v.DomainID != nil {
		fields[i] = fmt.Sprintf("DomainID: %v", *(v.DomainID))
		i++
	}
	if v.WorkflowID != nil {
		fields[i] = fmt.Sprintf("WorkflowID: %v", *(v.WorkflowID))
		i++
	}
	if v.RunID != nil {
		fields[i] = fmt.Sprintf("RunID: %v", *(v.RunID))
		i++
	}
	if v.Version != nil {
		fields[i] = fmt.Sprintf("Version: %v", *(v.Version))
		i++
	}
	if v.Fields != nil {
		fields[i] = fmt.Sprintf("Fields: %v", v.Fields)
		i++
	}

	return fmt.Sprintf("Message{%v}", strings.Join(fields[:i], ", "))
}

func _MessageType_EqualsPtr(lhs, rhs *MessageType) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

func _Map_String_Field_Equals(lhs, rhs map[string]*Field) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !lv.Equals(rv) {
			return false
		}
	}
	return true
}

// Equals returns true if all the fields of this Message match the
// provided Message.
//
// This function performs a deep comparison.
func (v *Message) Equals(rhs *Message) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_MessageType_EqualsPtr(v.MessageType, rhs.MessageType) {
		return false
	}
	if !_String_EqualsPtr(v.DomainID, rhs.DomainID) {
		return false
	}
	if !_String_EqualsPtr(v.WorkflowID, rhs.WorkflowID) {
		return false
	}
	if !_String_EqualsPtr(v.RunID, rhs.RunID) {
		return false
	}
	if !_I64_EqualsPtr(v.Version, rhs.Version) {
		return false
	}
	if !((v.Fields == nil && rhs.Fields == nil) || (v.Fields != nil && rhs.Fields != nil && _Map_String_Field_Equals(v.Fields, rhs.Fields))) {
		return false
	}

	return true
}

type _Map_String_Field_Zapper map[string]*Field

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of _Map_String_Field_Zapper.
func (m _Map_String_Field_Zapper) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	for k, v := range m {
		err = multierr.Append(err, enc.AddObject((string)(k), v))
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Message.
func (v *Message) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.MessageType != nil {
		err = multierr.Append(err, enc.AddObject("messageType", *v.MessageType))
	}
	if v.DomainID != nil {
		enc.AddString("domainID", *v.DomainID)
	}
	if v.WorkflowID != nil {
		enc.AddString("workflowID", *v.WorkflowID)
	}
	if v.RunID != nil {
		enc.AddString("runID", *v.RunID)
	}
	if v.Version != nil {
		enc.AddInt64("version", *v.Version)
	}
	if v.Fields != nil {
		err = multierr.Append(err, enc.AddObject("fields", (_Map_String_Field_Zapper)(v.Fields)))
	}
	return err
}

// GetMessageType returns the value of MessageType if it is set or its
// zero value if it is unset.
func (v *Message) GetMessageType() (o MessageType) {
	if v != nil && v.MessageType != nil {
		return *v.MessageType
	}

	return
}

// IsSetMessageType returns true if MessageType is not nil.
func (v *Message) IsSetMessageType() bool {
	return v != nil && v.MessageType != nil
}

// GetDomainID returns the value of DomainID if it is set or its
// zero value if it is unset.
func (v *Message) GetDomainID() (o string) {
	if v != nil && v.DomainID != nil {
		return *v.DomainID
	}

	return
}

// IsSetDomainID returns true if DomainID is not nil.
func (v *Message) IsSetDomainID() bool {
	return v != nil && v.DomainID != nil
}

// GetWorkflowID returns the value of WorkflowID if it is set or its
// zero value if it is unset.
func (v *Message) GetWorkflowID() (o string) {
	if v != nil && v.WorkflowID != nil {
		return *v.WorkflowID
	}

	return
}

// IsSetWorkflowID returns true if WorkflowID is not nil.
func (v *Message) IsSetWorkflowID() bool {
	return v != nil && v.WorkflowID != nil
}

// GetRunID returns the value of RunID if it is set or its
// zero value if it is unset.
func (v *Message) GetRunID() (o string) {
	if v != nil && v.RunID != nil {
		return *v.RunID
	}

	return
}

// IsSetRunID returns true if RunID is not nil.
func (v *Message) IsSetRunID() bool {
	return v != nil && v.RunID != nil
}

// GetVersion returns the value of Version if it is set or its
// zero value if it is unset.
func (v *Message) GetVersion() (o int64) {
	if v != nil && v.Version != nil {
		return *v.Version
	}

	return
}

// IsSetVersion returns true if Version is not nil.
func (v *Message) IsSetVersion() bool {
	return v != nil && v.Version != nil
}

// GetFields returns the value of Fields if it is set or its
// zero value if it is unset.
func (v *Message) GetFields() (o map[string]*Field) {
	if v != nil && v.Fields != nil {
		return v.Fields
	}

	return
}

// IsSetFields returns true if Fields is not nil.
func (v *Message) IsSetFields() bool {
	return v != nil && v.Fields != nil
}

type MessageType int32

const (
	MessageTypeIndex  MessageType = 0
	MessageTypeDelete MessageType = 1
)

// MessageType_Values returns all recognized values of MessageType.
func MessageType_Values() []MessageType {
	return []MessageType{
		MessageTypeIndex,
		MessageTypeDelete,
	}
}

// UnmarshalText tries to decode MessageType from a byte slice
// containing its name.
//
//   var v MessageType
//   err := v.UnmarshalText([]byte("Index"))
func (v *MessageType) UnmarshalText(value []byte) error {
	switch s := string(value); s {
	case "Index":
		*v = MessageTypeIndex
		return nil
	case "Delete":
		*v = MessageTypeDelete
		return nil
	default:
		val, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return fmt.Errorf("unknown enum value %q for %q: %v", s, "MessageType", err)
		}
		*v = MessageType(val)
		return nil
	}
}

// MarshalText encodes MessageType to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v MessageType) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("Index"), nil
	case 1:
		return []byte("Delete"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of MessageType.
// Enums are logged as objects, where the value is logged with key "value", and
// if this value's name is known, the name is logged with key "name".
func (v MessageType) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "Index")
	case 1:
		enc.AddString("name", "Delete")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v MessageType) Ptr() *MessageType {
	return &v
}

// ToWire translates MessageType into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v MessageType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes MessageType from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return MessageType(0), err
//   }
//
//   var v MessageType
//   if err := v.FromWire(x); err != nil {
//     return MessageType(0), err
//   }
//   return v, nil
func (v *MessageType) FromWire(w wire.Value) error {
	*v = (MessageType)(w.GetI32())
	return nil
}

// String returns a readable string representation of MessageType.
func (v MessageType) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "Index"
	case 1:
		return "Delete"
	}
	return fmt.Sprintf("MessageType(%d)", w)
}

// Equals returns true if this MessageType value matches the provided
// value.
func (v MessageType) Equals(rhs MessageType) bool {
	return v == rhs
}

// MarshalJSON serializes MessageType into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v MessageType) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"Index\""), nil
	case 1:
		return ([]byte)("\"Delete\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode MessageType from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *MessageType) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "MessageType")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "MessageType")
		}
		*v = (MessageType)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "MessageType")
	}
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "indexer",
	Package:  "github.com/uber/cadence/.gen/go/indexer",
	FilePath: "indexer.thrift",
	SHA1:     "464fff1b44654c3f41da6b7a95de46593a9776d6",
	Includes: []*thriftreflect.ThriftModule{
		shared.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "// Copyright (c) 2017 Uber Technologies, Inc.\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\nnamespace java com.uber.cadence.indexer\n\ninclude \"shared.thrift\"\n\nenum MessageType {\n  Index\n  Delete\n}\n\nenum FieldType {\n  String\n  Int\n  Bool\n  Binary\n}\n\nstruct Field {\n  10: optional FieldType type\n  20: optional string stringData\n  30: optional i64 (js.type = \"Long\") intData\n  40: optional bool boolData\n  50: optional binary binaryData\n}\n\nstruct Message {\n  10: optional MessageType messageType\n  20: optional string domainID\n  30: optional string workflowID\n  40: optional string runID\n  50: optional i64 (js.type = \"Long\") version\n  60: optional map<string,Field> fields\n}"
