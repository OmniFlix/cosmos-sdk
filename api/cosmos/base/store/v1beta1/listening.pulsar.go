package storev1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_StoreKVPair           protoreflect.MessageDescriptor
	fd_StoreKVPair_store_key protoreflect.FieldDescriptor
	fd_StoreKVPair_delete    protoreflect.FieldDescriptor
	fd_StoreKVPair_key       protoreflect.FieldDescriptor
	fd_StoreKVPair_value     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_store_v1beta1_listening_proto_init()
	md_StoreKVPair = File_cosmos_base_store_v1beta1_listening_proto.Messages().ByName("StoreKVPair")
	fd_StoreKVPair_store_key = md_StoreKVPair.Fields().ByName("store_key")
	fd_StoreKVPair_delete = md_StoreKVPair.Fields().ByName("delete")
	fd_StoreKVPair_key = md_StoreKVPair.Fields().ByName("key")
	fd_StoreKVPair_value = md_StoreKVPair.Fields().ByName("value")
}

var _ protoreflect.Message = (*fastReflection_StoreKVPair)(nil)

type fastReflection_StoreKVPair StoreKVPair

func (x *StoreKVPair) ProtoReflect() protoreflect.Message {
	return (*fastReflection_StoreKVPair)(x)
}

func (x *StoreKVPair) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_store_v1beta1_listening_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_StoreKVPair_messageType fastReflection_StoreKVPair_messageType
var _ protoreflect.MessageType = fastReflection_StoreKVPair_messageType{}

type fastReflection_StoreKVPair_messageType struct{}

func (x fastReflection_StoreKVPair_messageType) Zero() protoreflect.Message {
	return (*fastReflection_StoreKVPair)(nil)
}
func (x fastReflection_StoreKVPair_messageType) New() protoreflect.Message {
	return new(fastReflection_StoreKVPair)
}
func (x fastReflection_StoreKVPair_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_StoreKVPair
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_StoreKVPair) Descriptor() protoreflect.MessageDescriptor {
	return md_StoreKVPair
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_StoreKVPair) Type() protoreflect.MessageType {
	return _fastReflection_StoreKVPair_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_StoreKVPair) New() protoreflect.Message {
	return new(fastReflection_StoreKVPair)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_StoreKVPair) Interface() protoreflect.ProtoMessage {
	return (*StoreKVPair)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_StoreKVPair) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.StoreKey != "" {
		value := protoreflect.ValueOfString(x.StoreKey)
		if !f(fd_StoreKVPair_store_key, value) {
			return
		}
	}
	if x.Delete != false {
		value := protoreflect.ValueOfBool(x.Delete)
		if !f(fd_StoreKVPair_delete, value) {
			return
		}
	}
	if len(x.Key) != 0 {
		value := protoreflect.ValueOfBytes(x.Key)
		if !f(fd_StoreKVPair_key, value) {
			return
		}
	}
	if len(x.Value) != 0 {
		value := protoreflect.ValueOfBytes(x.Value)
		if !f(fd_StoreKVPair_value, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_StoreKVPair) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.StoreKVPair.store_key":
		return x.StoreKey != ""
	case "cosmos.base.store.v1beta1.StoreKVPair.delete":
		return x.Delete != false
	case "cosmos.base.store.v1beta1.StoreKVPair.key":
		return len(x.Key) != 0
	case "cosmos.base.store.v1beta1.StoreKVPair.value":
		return len(x.Value) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.StoreKVPair"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.StoreKVPair does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoreKVPair) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.StoreKVPair.store_key":
		x.StoreKey = ""
	case "cosmos.base.store.v1beta1.StoreKVPair.delete":
		x.Delete = false
	case "cosmos.base.store.v1beta1.StoreKVPair.key":
		x.Key = nil
	case "cosmos.base.store.v1beta1.StoreKVPair.value":
		x.Value = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.StoreKVPair"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.StoreKVPair does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_StoreKVPair) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.store.v1beta1.StoreKVPair.store_key":
		value := x.StoreKey
		return protoreflect.ValueOfString(value)
	case "cosmos.base.store.v1beta1.StoreKVPair.delete":
		value := x.Delete
		return protoreflect.ValueOfBool(value)
	case "cosmos.base.store.v1beta1.StoreKVPair.key":
		value := x.Key
		return protoreflect.ValueOfBytes(value)
	case "cosmos.base.store.v1beta1.StoreKVPair.value":
		value := x.Value
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.StoreKVPair"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.StoreKVPair does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoreKVPair) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.StoreKVPair.store_key":
		x.StoreKey = value.Interface().(string)
	case "cosmos.base.store.v1beta1.StoreKVPair.delete":
		x.Delete = value.Bool()
	case "cosmos.base.store.v1beta1.StoreKVPair.key":
		x.Key = value.Bytes()
	case "cosmos.base.store.v1beta1.StoreKVPair.value":
		x.Value = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.StoreKVPair"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.StoreKVPair does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoreKVPair) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.StoreKVPair.store_key":
		panic(fmt.Errorf("field store_key of message cosmos.base.store.v1beta1.StoreKVPair is not mutable"))
	case "cosmos.base.store.v1beta1.StoreKVPair.delete":
		panic(fmt.Errorf("field delete of message cosmos.base.store.v1beta1.StoreKVPair is not mutable"))
	case "cosmos.base.store.v1beta1.StoreKVPair.key":
		panic(fmt.Errorf("field key of message cosmos.base.store.v1beta1.StoreKVPair is not mutable"))
	case "cosmos.base.store.v1beta1.StoreKVPair.value":
		panic(fmt.Errorf("field value of message cosmos.base.store.v1beta1.StoreKVPair is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.StoreKVPair"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.StoreKVPair does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_StoreKVPair) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.StoreKVPair.store_key":
		return protoreflect.ValueOfString("")
	case "cosmos.base.store.v1beta1.StoreKVPair.delete":
		return protoreflect.ValueOfBool(false)
	case "cosmos.base.store.v1beta1.StoreKVPair.key":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.base.store.v1beta1.StoreKVPair.value":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.StoreKVPair"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.StoreKVPair does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_StoreKVPair) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.store.v1beta1.StoreKVPair", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_StoreKVPair) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StoreKVPair) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_StoreKVPair) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_StoreKVPair) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*StoreKVPair)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.StoreKey)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Delete {
			n += 2
		}
		l = len(x.Key)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Value)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*StoreKVPair)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Value) > 0 {
			i -= len(x.Value)
			copy(dAtA[i:], x.Value)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Value)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Key) > 0 {
			i -= len(x.Key)
			copy(dAtA[i:], x.Key)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Key)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Delete {
			i--
			if x.Delete {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
		}
		if len(x.StoreKey) > 0 {
			i -= len(x.StoreKey)
			copy(dAtA[i:], x.StoreKey)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.StoreKey)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*StoreKVPair)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StoreKVPair: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StoreKVPair: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StoreKey", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.StoreKey = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Delete", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Delete = bool(v != 0)
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Key = append(x.Key[:0], dAtA[iNdEx:postIndex]...)
				if x.Key == nil {
					x.Key = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Value = append(x.Value[:0], dAtA[iNdEx:postIndex]...)
				if x.Value == nil {
					x.Value = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/base/store/v1beta1/listening.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// StoreKVPair is a KVStore KVPair used for listening to state changes (Sets and Deletes)
// It optionally includes the StoreKey for the originating KVStore and a Boolean flag to distinguish between Sets and
// Deletes
//
// Since: cosmos-sdk 0.43
type StoreKVPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreKey string `protobuf:"bytes,1,opt,name=store_key,json=storeKey,proto3" json:"store_key,omitempty"` // the store key for the KVStore this pair originates from
	Delete   bool   `protobuf:"varint,2,opt,name=delete,proto3" json:"delete,omitempty"`                    // true indicates a delete operation, false indicates a set operation
	Key      []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value    []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StoreKVPair) Reset() {
	*x = StoreKVPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_store_v1beta1_listening_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreKVPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreKVPair) ProtoMessage() {}

// Deprecated: Use StoreKVPair.ProtoReflect.Descriptor instead.
func (*StoreKVPair) Descriptor() ([]byte, []int) {
	return file_cosmos_base_store_v1beta1_listening_proto_rawDescGZIP(), []int{0}
}

func (x *StoreKVPair) GetStoreKey() string {
	if x != nil {
		return x.StoreKey
	}
	return ""
}

func (x *StoreKVPair) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

func (x *StoreKVPair) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *StoreKVPair) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_cosmos_base_store_v1beta1_listening_proto protoreflect.FileDescriptor

var file_cosmos_base_store_v1beta1_listening_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x6a, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4b,
	0x56, 0x50, 0x61, 0x69, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0xff, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x42, 0x53, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xca, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x25,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a,
	0x42, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_base_store_v1beta1_listening_proto_rawDescOnce sync.Once
	file_cosmos_base_store_v1beta1_listening_proto_rawDescData = file_cosmos_base_store_v1beta1_listening_proto_rawDesc
)

func file_cosmos_base_store_v1beta1_listening_proto_rawDescGZIP() []byte {
	file_cosmos_base_store_v1beta1_listening_proto_rawDescOnce.Do(func() {
		file_cosmos_base_store_v1beta1_listening_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_base_store_v1beta1_listening_proto_rawDescData)
	})
	return file_cosmos_base_store_v1beta1_listening_proto_rawDescData
}

var file_cosmos_base_store_v1beta1_listening_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_base_store_v1beta1_listening_proto_goTypes = []interface{}{
	(*StoreKVPair)(nil), // 0: cosmos.base.store.v1beta1.StoreKVPair
}
var file_cosmos_base_store_v1beta1_listening_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_base_store_v1beta1_listening_proto_init() }
func file_cosmos_base_store_v1beta1_listening_proto_init() {
	if File_cosmos_base_store_v1beta1_listening_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_base_store_v1beta1_listening_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreKVPair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_base_store_v1beta1_listening_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_base_store_v1beta1_listening_proto_goTypes,
		DependencyIndexes: file_cosmos_base_store_v1beta1_listening_proto_depIdxs,
		MessageInfos:      file_cosmos_base_store_v1beta1_listening_proto_msgTypes,
	}.Build()
	File_cosmos_base_store_v1beta1_listening_proto = out.File
	file_cosmos_base_store_v1beta1_listening_proto_rawDesc = nil
	file_cosmos_base_store_v1beta1_listening_proto_goTypes = nil
	file_cosmos_base_store_v1beta1_listening_proto_depIdxs = nil
}
