package secp256r1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_PubKey     protoreflect.MessageDescriptor
	fd_PubKey_key protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_secp256r1_keys_proto_init()
	md_PubKey = File_cosmos_crypto_secp256r1_keys_proto.Messages().ByName("PubKey")
	fd_PubKey_key = md_PubKey.Fields().ByName("key")
}

var _ protoreflect.Message = (*fastReflection_PubKey)(nil)

type fastReflection_PubKey PubKey

func (x *PubKey) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PubKey)(x)
}

func (x *PubKey) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_secp256r1_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PubKey_messageType fastReflection_PubKey_messageType
var _ protoreflect.MessageType = fastReflection_PubKey_messageType{}

type fastReflection_PubKey_messageType struct{}

func (x fastReflection_PubKey_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PubKey)(nil)
}
func (x fastReflection_PubKey_messageType) New() protoreflect.Message {
	return new(fastReflection_PubKey)
}
func (x fastReflection_PubKey_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PubKey
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PubKey) Descriptor() protoreflect.MessageDescriptor {
	return md_PubKey
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PubKey) Type() protoreflect.MessageType {
	return _fastReflection_PubKey_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PubKey) New() protoreflect.Message {
	return new(fastReflection_PubKey)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PubKey) Interface() protoreflect.ProtoMessage {
	return (*PubKey)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PubKey) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Key) != 0 {
		value := protoreflect.ValueOfBytes(x.Key)
		if !f(fd_PubKey_key, value) {
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
func (x *fastReflection_PubKey) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PubKey.key":
		return len(x.Key) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PubKey does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PubKey) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PubKey.key":
		x.Key = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PubKey does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PubKey) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.secp256r1.PubKey.key":
		value := x.Key
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PubKey does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_PubKey) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PubKey.key":
		x.Key = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PubKey does not contain field %s", fd.FullName()))
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
func (x *fastReflection_PubKey) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PubKey.key":
		panic(fmt.Errorf("field key of message cosmos.crypto.secp256r1.PubKey is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PubKey does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PubKey) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PubKey.key":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PubKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PubKey does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PubKey) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.secp256r1.PubKey", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PubKey) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PubKey) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_PubKey) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PubKey) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PubKey)
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
		l = len(x.Key)
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
		x := input.Message.Interface().(*PubKey)
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
		if len(x.Key) > 0 {
			i -= len(x.Key)
			copy(dAtA[i:], x.Key)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Key)))
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
		x := input.Message.Interface().(*PubKey)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PubKey: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PubKey: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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

var (
	md_PrivKey        protoreflect.MessageDescriptor
	fd_PrivKey_secret protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_secp256r1_keys_proto_init()
	md_PrivKey = File_cosmos_crypto_secp256r1_keys_proto.Messages().ByName("PrivKey")
	fd_PrivKey_secret = md_PrivKey.Fields().ByName("secret")
}

var _ protoreflect.Message = (*fastReflection_PrivKey)(nil)

type fastReflection_PrivKey PrivKey

func (x *PrivKey) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PrivKey)(x)
}

func (x *PrivKey) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_secp256r1_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PrivKey_messageType fastReflection_PrivKey_messageType
var _ protoreflect.MessageType = fastReflection_PrivKey_messageType{}

type fastReflection_PrivKey_messageType struct{}

func (x fastReflection_PrivKey_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PrivKey)(nil)
}
func (x fastReflection_PrivKey_messageType) New() protoreflect.Message {
	return new(fastReflection_PrivKey)
}
func (x fastReflection_PrivKey_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PrivKey
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PrivKey) Descriptor() protoreflect.MessageDescriptor {
	return md_PrivKey
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PrivKey) Type() protoreflect.MessageType {
	return _fastReflection_PrivKey_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PrivKey) New() protoreflect.Message {
	return new(fastReflection_PrivKey)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PrivKey) Interface() protoreflect.ProtoMessage {
	return (*PrivKey)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PrivKey) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Secret) != 0 {
		value := protoreflect.ValueOfBytes(x.Secret)
		if !f(fd_PrivKey_secret, value) {
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
func (x *fastReflection_PrivKey) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PrivKey.secret":
		return len(x.Secret) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PrivKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PrivKey does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PrivKey) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PrivKey.secret":
		x.Secret = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PrivKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PrivKey does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PrivKey) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.secp256r1.PrivKey.secret":
		value := x.Secret
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PrivKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PrivKey does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_PrivKey) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PrivKey.secret":
		x.Secret = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PrivKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PrivKey does not contain field %s", fd.FullName()))
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
func (x *fastReflection_PrivKey) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PrivKey.secret":
		panic(fmt.Errorf("field secret of message cosmos.crypto.secp256r1.PrivKey is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PrivKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PrivKey does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PrivKey) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.secp256r1.PrivKey.secret":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.secp256r1.PrivKey"))
		}
		panic(fmt.Errorf("message cosmos.crypto.secp256r1.PrivKey does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PrivKey) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.secp256r1.PrivKey", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PrivKey) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PrivKey) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_PrivKey) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PrivKey) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PrivKey)
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
		l = len(x.Secret)
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
		x := input.Message.Interface().(*PrivKey)
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
		if len(x.Secret) > 0 {
			i -= len(x.Secret)
			copy(dAtA[i:], x.Secret)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Secret)))
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
		x := input.Message.Interface().(*PrivKey)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PrivKey: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PrivKey: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
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
				x.Secret = append(x.Secret[:0], dAtA[iNdEx:postIndex]...)
				if x.Secret == nil {
					x.Secret = []byte{}
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

// Since: cosmos-sdk 0.43

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/crypto/secp256r1/keys.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PubKey defines a secp256r1 ECDSA public key.
type PubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Point on secp256r1 curve in a compressed representation as specified in section
	// 4.3.6 of ANSI X9.62: https://webstore.ansi.org/standards/ascx9/ansix9621998
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PubKey) Reset() {
	*x = PubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_secp256r1_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKey) ProtoMessage() {}

// Deprecated: Use PubKey.ProtoReflect.Descriptor instead.
func (*PubKey) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_secp256r1_keys_proto_rawDescGZIP(), []int{0}
}

func (x *PubKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

// PrivKey defines a secp256r1 ECDSA private key.
type PrivKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// secret number serialized using big-endian encoding
	Secret []byte `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *PrivKey) Reset() {
	*x = PrivKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_secp256r1_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivKey) ProtoMessage() {}

// Deprecated: Use PrivKey.ProtoReflect.Descriptor instead.
func (*PrivKey) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_secp256r1_keys_proto_rawDescGZIP(), []int{1}
}

func (x *PrivKey) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

var File_cosmos_crypto_secp256r1_keys_proto protoreflect.FileDescriptor

var file_cosmos_crypto_secp256r1_keys_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x72, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x72, 0x31, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0b, 0xda, 0xde, 0x1f, 0x07,
	0x65, 0x63, 0x64, 0x73, 0x61, 0x50, 0x4b, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2e, 0x0a, 0x07,
	0x50, 0x72, 0x69, 0x76, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0b, 0xda, 0xde, 0x1f, 0x07, 0x65, 0x63, 0x64,
	0x73, 0x61, 0x53, 0x4b, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0xec, 0x01, 0x0a,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x72, 0x31, 0x42, 0x09, 0x4b, 0x65,
	0x79, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35,
	0x36, 0x72, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x53, 0xaa, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x70, 0x32, 0x35,
	0x36, 0x72, 0x31, 0xca, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x5c, 0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x72, 0x31, 0xe2, 0x02, 0x23,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x53, 0x65,
	0x63, 0x70, 0x32, 0x35, 0x36, 0x72, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x72, 0x31, 0xc8,
	0xe1, 0x1e, 0x00, 0xd8, 0xe1, 0x1e, 0x00, 0xc8, 0xe3, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cosmos_crypto_secp256r1_keys_proto_rawDescOnce sync.Once
	file_cosmos_crypto_secp256r1_keys_proto_rawDescData = file_cosmos_crypto_secp256r1_keys_proto_rawDesc
)

func file_cosmos_crypto_secp256r1_keys_proto_rawDescGZIP() []byte {
	file_cosmos_crypto_secp256r1_keys_proto_rawDescOnce.Do(func() {
		file_cosmos_crypto_secp256r1_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_crypto_secp256r1_keys_proto_rawDescData)
	})
	return file_cosmos_crypto_secp256r1_keys_proto_rawDescData
}

var file_cosmos_crypto_secp256r1_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_crypto_secp256r1_keys_proto_goTypes = []interface{}{
	(*PubKey)(nil),  // 0: cosmos.crypto.secp256r1.PubKey
	(*PrivKey)(nil), // 1: cosmos.crypto.secp256r1.PrivKey
}
var file_cosmos_crypto_secp256r1_keys_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_crypto_secp256r1_keys_proto_init() }
func file_cosmos_crypto_secp256r1_keys_proto_init() {
	if File_cosmos_crypto_secp256r1_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_crypto_secp256r1_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKey); i {
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
		file_cosmos_crypto_secp256r1_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivKey); i {
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
			RawDescriptor: file_cosmos_crypto_secp256r1_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_crypto_secp256r1_keys_proto_goTypes,
		DependencyIndexes: file_cosmos_crypto_secp256r1_keys_proto_depIdxs,
		MessageInfos:      file_cosmos_crypto_secp256r1_keys_proto_msgTypes,
	}.Build()
	File_cosmos_crypto_secp256r1_keys_proto = out.File
	file_cosmos_crypto_secp256r1_keys_proto_rawDesc = nil
	file_cosmos_crypto_secp256r1_keys_proto_goTypes = nil
	file_cosmos_crypto_secp256r1_keys_proto_depIdxs = nil
}
