package stakingv1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_QueryValidatorsRequest            protoreflect.MessageDescriptor
	fd_QueryValidatorsRequest_status     protoreflect.FieldDescriptor
	fd_QueryValidatorsRequest_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryValidatorsRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryValidatorsRequest")
	fd_QueryValidatorsRequest_status = md_QueryValidatorsRequest.Fields().ByName("status")
	fd_QueryValidatorsRequest_pagination = md_QueryValidatorsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorsRequest)(nil)

type fastReflection_QueryValidatorsRequest QueryValidatorsRequest

func (x *QueryValidatorsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorsRequest)(x)
}

func (x *QueryValidatorsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorsRequest_messageType fastReflection_QueryValidatorsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorsRequest_messageType{}

type fastReflection_QueryValidatorsRequest_messageType struct{}

func (x fastReflection_QueryValidatorsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorsRequest)(nil)
}
func (x fastReflection_QueryValidatorsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorsRequest)
}
func (x fastReflection_QueryValidatorsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Status != "" {
		value := protoreflect.ValueOfString(x.Status)
		if !f(fd_QueryValidatorsRequest_status, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryValidatorsRequest_pagination, value) {
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
func (x *fastReflection_QueryValidatorsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.status":
		return x.Status != ""
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.status":
		x.Status = ""
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.status":
		value := x.Status
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryValidatorsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.status":
		x.Status = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryValidatorsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.status":
		panic(fmt.Errorf("field status of message cosmos.staking.v1beta1.QueryValidatorsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.status":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryValidatorsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryValidatorsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryValidatorsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorsRequest)
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
		l = len(x.Status)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryValidatorsRequest)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Status) > 0 {
			i -= len(x.Status)
			copy(dAtA[i:], x.Status)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Status)))
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
		x := input.Message.Interface().(*QueryValidatorsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
				x.Status = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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

var _ protoreflect.List = (*_QueryValidatorsResponse_1_list)(nil)

type _QueryValidatorsResponse_1_list struct {
	list *[]*Validator
}

func (x *_QueryValidatorsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryValidatorsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryValidatorsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	(*x.list)[i] = concreteValue
}

func (x *_QueryValidatorsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryValidatorsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(Validator)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryValidatorsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryValidatorsResponse_1_list) NewElement() protoreflect.Value {
	v := new(Validator)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryValidatorsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryValidatorsResponse            protoreflect.MessageDescriptor
	fd_QueryValidatorsResponse_validators protoreflect.FieldDescriptor
	fd_QueryValidatorsResponse_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryValidatorsResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryValidatorsResponse")
	fd_QueryValidatorsResponse_validators = md_QueryValidatorsResponse.Fields().ByName("validators")
	fd_QueryValidatorsResponse_pagination = md_QueryValidatorsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorsResponse)(nil)

type fastReflection_QueryValidatorsResponse QueryValidatorsResponse

func (x *QueryValidatorsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorsResponse)(x)
}

func (x *QueryValidatorsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorsResponse_messageType fastReflection_QueryValidatorsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorsResponse_messageType{}

type fastReflection_QueryValidatorsResponse_messageType struct{}

func (x fastReflection_QueryValidatorsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorsResponse)(nil)
}
func (x fastReflection_QueryValidatorsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorsResponse)
}
func (x fastReflection_QueryValidatorsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Validators) != 0 {
		value := protoreflect.ValueOfList(&_QueryValidatorsResponse_1_list{list: &x.Validators})
		if !f(fd_QueryValidatorsResponse_validators, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryValidatorsResponse_pagination, value) {
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
func (x *fastReflection_QueryValidatorsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.validators":
		return len(x.Validators) != 0
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.validators":
		x.Validators = nil
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.validators":
		if len(x.Validators) == 0 {
			return protoreflect.ValueOfList(&_QueryValidatorsResponse_1_list{})
		}
		listValue := &_QueryValidatorsResponse_1_list{list: &x.Validators}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryValidatorsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.validators":
		lv := value.List()
		clv := lv.(*_QueryValidatorsResponse_1_list)
		x.Validators = *clv.list
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryValidatorsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.validators":
		if x.Validators == nil {
			x.Validators = []*Validator{}
		}
		value := &_QueryValidatorsResponse_1_list{list: &x.Validators}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.validators":
		list := []*Validator{}
		return protoreflect.ValueOfList(&_QueryValidatorsResponse_1_list{list: &list})
	case "cosmos.staking.v1beta1.QueryValidatorsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryValidatorsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryValidatorsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorsResponse)
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
		if len(x.Validators) > 0 {
			for _, e := range x.Validators {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryValidatorsResponse)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Validators) > 0 {
			for iNdEx := len(x.Validators) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Validators[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*QueryValidatorsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Validators = append(x.Validators, &Validator{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validators[len(x.Validators)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryValidatorRequest                protoreflect.MessageDescriptor
	fd_QueryValidatorRequest_validator_addr protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryValidatorRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryValidatorRequest")
	fd_QueryValidatorRequest_validator_addr = md_QueryValidatorRequest.Fields().ByName("validator_addr")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorRequest)(nil)

type fastReflection_QueryValidatorRequest QueryValidatorRequest

func (x *QueryValidatorRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorRequest)(x)
}

func (x *QueryValidatorRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorRequest_messageType fastReflection_QueryValidatorRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorRequest_messageType{}

type fastReflection_QueryValidatorRequest_messageType struct{}

func (x fastReflection_QueryValidatorRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorRequest)(nil)
}
func (x fastReflection_QueryValidatorRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorRequest)
}
func (x fastReflection_QueryValidatorRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorRequest) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddr != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddr)
		if !f(fd_QueryValidatorRequest_validator_addr, value) {
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
func (x *fastReflection_QueryValidatorRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorRequest.validator_addr":
		return x.ValidatorAddr != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorRequest.validator_addr":
		x.ValidatorAddr = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorRequest.validator_addr":
		value := x.ValidatorAddr
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryValidatorRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorRequest.validator_addr":
		x.ValidatorAddr = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryValidatorRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorRequest.validator_addr":
		panic(fmt.Errorf("field validator_addr of message cosmos.staking.v1beta1.QueryValidatorRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorRequest.validator_addr":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryValidatorRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryValidatorRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorRequest)
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
		l = len(x.ValidatorAddr)
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
		x := input.Message.Interface().(*QueryValidatorRequest)
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
		if len(x.ValidatorAddr) > 0 {
			i -= len(x.ValidatorAddr)
			copy(dAtA[i:], x.ValidatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddr)))
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
		x := input.Message.Interface().(*QueryValidatorRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
				x.ValidatorAddr = string(dAtA[iNdEx:postIndex])
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
	md_QueryValidatorResponse           protoreflect.MessageDescriptor
	fd_QueryValidatorResponse_validator protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryValidatorResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryValidatorResponse")
	fd_QueryValidatorResponse_validator = md_QueryValidatorResponse.Fields().ByName("validator")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorResponse)(nil)

type fastReflection_QueryValidatorResponse QueryValidatorResponse

func (x *QueryValidatorResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorResponse)(x)
}

func (x *QueryValidatorResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorResponse_messageType fastReflection_QueryValidatorResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorResponse_messageType{}

type fastReflection_QueryValidatorResponse_messageType struct{}

func (x fastReflection_QueryValidatorResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorResponse)(nil)
}
func (x fastReflection_QueryValidatorResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorResponse)
}
func (x fastReflection_QueryValidatorResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorResponse) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Validator != nil {
		value := protoreflect.ValueOfMessage(x.Validator.ProtoReflect())
		if !f(fd_QueryValidatorResponse_validator, value) {
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
func (x *fastReflection_QueryValidatorResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorResponse.validator":
		return x.Validator != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorResponse.validator":
		x.Validator = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorResponse.validator":
		value := x.Validator
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryValidatorResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorResponse.validator":
		x.Validator = value.Message().Interface().(*Validator)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryValidatorResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorResponse.validator":
		if x.Validator == nil {
			x.Validator = new(Validator)
		}
		return protoreflect.ValueOfMessage(x.Validator.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorResponse.validator":
		m := new(Validator)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryValidatorResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryValidatorResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorResponse)
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
		if x.Validator != nil {
			l = options.Size(x.Validator)
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
		x := input.Message.Interface().(*QueryValidatorResponse)
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
		if x.Validator != nil {
			encoded, err := options.Marshal(x.Validator)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*QueryValidatorResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Validator == nil {
					x.Validator = &Validator{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validator); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryValidatorDelegationsRequest                protoreflect.MessageDescriptor
	fd_QueryValidatorDelegationsRequest_validator_addr protoreflect.FieldDescriptor
	fd_QueryValidatorDelegationsRequest_pagination     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryValidatorDelegationsRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryValidatorDelegationsRequest")
	fd_QueryValidatorDelegationsRequest_validator_addr = md_QueryValidatorDelegationsRequest.Fields().ByName("validator_addr")
	fd_QueryValidatorDelegationsRequest_pagination = md_QueryValidatorDelegationsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorDelegationsRequest)(nil)

type fastReflection_QueryValidatorDelegationsRequest QueryValidatorDelegationsRequest

func (x *QueryValidatorDelegationsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorDelegationsRequest)(x)
}

func (x *QueryValidatorDelegationsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorDelegationsRequest_messageType fastReflection_QueryValidatorDelegationsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorDelegationsRequest_messageType{}

type fastReflection_QueryValidatorDelegationsRequest_messageType struct{}

func (x fastReflection_QueryValidatorDelegationsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorDelegationsRequest)(nil)
}
func (x fastReflection_QueryValidatorDelegationsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorDelegationsRequest)
}
func (x fastReflection_QueryValidatorDelegationsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorDelegationsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorDelegationsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorDelegationsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorDelegationsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorDelegationsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorDelegationsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorDelegationsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorDelegationsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorDelegationsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorDelegationsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddr != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddr)
		if !f(fd_QueryValidatorDelegationsRequest_validator_addr, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryValidatorDelegationsRequest_pagination, value) {
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
func (x *fastReflection_QueryValidatorDelegationsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.validator_addr":
		return x.ValidatorAddr != ""
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorDelegationsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.validator_addr":
		x.ValidatorAddr = ""
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorDelegationsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.validator_addr":
		value := x.ValidatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryValidatorDelegationsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.validator_addr":
		x.ValidatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryValidatorDelegationsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.validator_addr":
		panic(fmt.Errorf("field validator_addr of message cosmos.staking.v1beta1.QueryValidatorDelegationsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorDelegationsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.validator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorDelegationsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryValidatorDelegationsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorDelegationsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorDelegationsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryValidatorDelegationsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorDelegationsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorDelegationsRequest)
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
		l = len(x.ValidatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryValidatorDelegationsRequest)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ValidatorAddr) > 0 {
			i -= len(x.ValidatorAddr)
			copy(dAtA[i:], x.ValidatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddr)))
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
		x := input.Message.Interface().(*QueryValidatorDelegationsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorDelegationsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorDelegationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
				x.ValidatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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

var _ protoreflect.List = (*_QueryValidatorDelegationsResponse_1_list)(nil)

type _QueryValidatorDelegationsResponse_1_list struct {
	list *[]*DelegationResponse
}

func (x *_QueryValidatorDelegationsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryValidatorDelegationsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryValidatorDelegationsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegationResponse)
	(*x.list)[i] = concreteValue
}

func (x *_QueryValidatorDelegationsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegationResponse)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryValidatorDelegationsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(DelegationResponse)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryValidatorDelegationsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryValidatorDelegationsResponse_1_list) NewElement() protoreflect.Value {
	v := new(DelegationResponse)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryValidatorDelegationsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryValidatorDelegationsResponse                      protoreflect.MessageDescriptor
	fd_QueryValidatorDelegationsResponse_delegation_responses protoreflect.FieldDescriptor
	fd_QueryValidatorDelegationsResponse_pagination           protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryValidatorDelegationsResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryValidatorDelegationsResponse")
	fd_QueryValidatorDelegationsResponse_delegation_responses = md_QueryValidatorDelegationsResponse.Fields().ByName("delegation_responses")
	fd_QueryValidatorDelegationsResponse_pagination = md_QueryValidatorDelegationsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorDelegationsResponse)(nil)

type fastReflection_QueryValidatorDelegationsResponse QueryValidatorDelegationsResponse

func (x *QueryValidatorDelegationsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorDelegationsResponse)(x)
}

func (x *QueryValidatorDelegationsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorDelegationsResponse_messageType fastReflection_QueryValidatorDelegationsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorDelegationsResponse_messageType{}

type fastReflection_QueryValidatorDelegationsResponse_messageType struct{}

func (x fastReflection_QueryValidatorDelegationsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorDelegationsResponse)(nil)
}
func (x fastReflection_QueryValidatorDelegationsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorDelegationsResponse)
}
func (x fastReflection_QueryValidatorDelegationsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorDelegationsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorDelegationsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorDelegationsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorDelegationsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorDelegationsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorDelegationsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorDelegationsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorDelegationsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorDelegationsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorDelegationsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.DelegationResponses) != 0 {
		value := protoreflect.ValueOfList(&_QueryValidatorDelegationsResponse_1_list{list: &x.DelegationResponses})
		if !f(fd_QueryValidatorDelegationsResponse_delegation_responses, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryValidatorDelegationsResponse_pagination, value) {
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
func (x *fastReflection_QueryValidatorDelegationsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.delegation_responses":
		return len(x.DelegationResponses) != 0
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorDelegationsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.delegation_responses":
		x.DelegationResponses = nil
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorDelegationsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.delegation_responses":
		if len(x.DelegationResponses) == 0 {
			return protoreflect.ValueOfList(&_QueryValidatorDelegationsResponse_1_list{})
		}
		listValue := &_QueryValidatorDelegationsResponse_1_list{list: &x.DelegationResponses}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryValidatorDelegationsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.delegation_responses":
		lv := value.List()
		clv := lv.(*_QueryValidatorDelegationsResponse_1_list)
		x.DelegationResponses = *clv.list
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryValidatorDelegationsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.delegation_responses":
		if x.DelegationResponses == nil {
			x.DelegationResponses = []*DelegationResponse{}
		}
		value := &_QueryValidatorDelegationsResponse_1_list{list: &x.DelegationResponses}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorDelegationsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.delegation_responses":
		list := []*DelegationResponse{}
		return protoreflect.ValueOfList(&_QueryValidatorDelegationsResponse_1_list{list: &list})
	case "cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorDelegationsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryValidatorDelegationsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorDelegationsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorDelegationsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryValidatorDelegationsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorDelegationsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorDelegationsResponse)
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
		if len(x.DelegationResponses) > 0 {
			for _, e := range x.DelegationResponses {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryValidatorDelegationsResponse)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegationResponses) > 0 {
			for iNdEx := len(x.DelegationResponses) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.DelegationResponses[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*QueryValidatorDelegationsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorDelegationsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorDelegationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegationResponses", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegationResponses = append(x.DelegationResponses, &DelegationResponse{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DelegationResponses[len(x.DelegationResponses)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryValidatorUnbondingDelegationsRequest                protoreflect.MessageDescriptor
	fd_QueryValidatorUnbondingDelegationsRequest_validator_addr protoreflect.FieldDescriptor
	fd_QueryValidatorUnbondingDelegationsRequest_pagination     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryValidatorUnbondingDelegationsRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryValidatorUnbondingDelegationsRequest")
	fd_QueryValidatorUnbondingDelegationsRequest_validator_addr = md_QueryValidatorUnbondingDelegationsRequest.Fields().ByName("validator_addr")
	fd_QueryValidatorUnbondingDelegationsRequest_pagination = md_QueryValidatorUnbondingDelegationsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorUnbondingDelegationsRequest)(nil)

type fastReflection_QueryValidatorUnbondingDelegationsRequest QueryValidatorUnbondingDelegationsRequest

func (x *QueryValidatorUnbondingDelegationsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorUnbondingDelegationsRequest)(x)
}

func (x *QueryValidatorUnbondingDelegationsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorUnbondingDelegationsRequest_messageType fastReflection_QueryValidatorUnbondingDelegationsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorUnbondingDelegationsRequest_messageType{}

type fastReflection_QueryValidatorUnbondingDelegationsRequest_messageType struct{}

func (x fastReflection_QueryValidatorUnbondingDelegationsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorUnbondingDelegationsRequest)(nil)
}
func (x fastReflection_QueryValidatorUnbondingDelegationsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorUnbondingDelegationsRequest)
}
func (x fastReflection_QueryValidatorUnbondingDelegationsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorUnbondingDelegationsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorUnbondingDelegationsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorUnbondingDelegationsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorUnbondingDelegationsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorUnbondingDelegationsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddr != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddr)
		if !f(fd_QueryValidatorUnbondingDelegationsRequest_validator_addr, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryValidatorUnbondingDelegationsRequest_pagination, value) {
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
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.validator_addr":
		return x.ValidatorAddr != ""
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.validator_addr":
		x.ValidatorAddr = ""
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.validator_addr":
		value := x.ValidatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.validator_addr":
		x.ValidatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.validator_addr":
		panic(fmt.Errorf("field validator_addr of message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.validator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorUnbondingDelegationsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorUnbondingDelegationsRequest)
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
		l = len(x.ValidatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryValidatorUnbondingDelegationsRequest)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ValidatorAddr) > 0 {
			i -= len(x.ValidatorAddr)
			copy(dAtA[i:], x.ValidatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddr)))
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
		x := input.Message.Interface().(*QueryValidatorUnbondingDelegationsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorUnbondingDelegationsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorUnbondingDelegationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
				x.ValidatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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

var _ protoreflect.List = (*_QueryValidatorUnbondingDelegationsResponse_1_list)(nil)

type _QueryValidatorUnbondingDelegationsResponse_1_list struct {
	list *[]*UnbondingDelegation
}

func (x *_QueryValidatorUnbondingDelegationsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryValidatorUnbondingDelegationsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryValidatorUnbondingDelegationsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*UnbondingDelegation)
	(*x.list)[i] = concreteValue
}

func (x *_QueryValidatorUnbondingDelegationsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*UnbondingDelegation)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryValidatorUnbondingDelegationsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(UnbondingDelegation)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryValidatorUnbondingDelegationsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryValidatorUnbondingDelegationsResponse_1_list) NewElement() protoreflect.Value {
	v := new(UnbondingDelegation)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryValidatorUnbondingDelegationsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryValidatorUnbondingDelegationsResponse                     protoreflect.MessageDescriptor
	fd_QueryValidatorUnbondingDelegationsResponse_unbonding_responses protoreflect.FieldDescriptor
	fd_QueryValidatorUnbondingDelegationsResponse_pagination          protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryValidatorUnbondingDelegationsResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryValidatorUnbondingDelegationsResponse")
	fd_QueryValidatorUnbondingDelegationsResponse_unbonding_responses = md_QueryValidatorUnbondingDelegationsResponse.Fields().ByName("unbonding_responses")
	fd_QueryValidatorUnbondingDelegationsResponse_pagination = md_QueryValidatorUnbondingDelegationsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryValidatorUnbondingDelegationsResponse)(nil)

type fastReflection_QueryValidatorUnbondingDelegationsResponse QueryValidatorUnbondingDelegationsResponse

func (x *QueryValidatorUnbondingDelegationsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryValidatorUnbondingDelegationsResponse)(x)
}

func (x *QueryValidatorUnbondingDelegationsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryValidatorUnbondingDelegationsResponse_messageType fastReflection_QueryValidatorUnbondingDelegationsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryValidatorUnbondingDelegationsResponse_messageType{}

type fastReflection_QueryValidatorUnbondingDelegationsResponse_messageType struct{}

func (x fastReflection_QueryValidatorUnbondingDelegationsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryValidatorUnbondingDelegationsResponse)(nil)
}
func (x fastReflection_QueryValidatorUnbondingDelegationsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorUnbondingDelegationsResponse)
}
func (x fastReflection_QueryValidatorUnbondingDelegationsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorUnbondingDelegationsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryValidatorUnbondingDelegationsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryValidatorUnbondingDelegationsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryValidatorUnbondingDelegationsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryValidatorUnbondingDelegationsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.UnbondingResponses) != 0 {
		value := protoreflect.ValueOfList(&_QueryValidatorUnbondingDelegationsResponse_1_list{list: &x.UnbondingResponses})
		if !f(fd_QueryValidatorUnbondingDelegationsResponse_unbonding_responses, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryValidatorUnbondingDelegationsResponse_pagination, value) {
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
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.unbonding_responses":
		return len(x.UnbondingResponses) != 0
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.unbonding_responses":
		x.UnbondingResponses = nil
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.unbonding_responses":
		if len(x.UnbondingResponses) == 0 {
			return protoreflect.ValueOfList(&_QueryValidatorUnbondingDelegationsResponse_1_list{})
		}
		listValue := &_QueryValidatorUnbondingDelegationsResponse_1_list{list: &x.UnbondingResponses}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.unbonding_responses":
		lv := value.List()
		clv := lv.(*_QueryValidatorUnbondingDelegationsResponse_1_list)
		x.UnbondingResponses = *clv.list
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.unbonding_responses":
		if x.UnbondingResponses == nil {
			x.UnbondingResponses = []*UnbondingDelegation{}
		}
		value := &_QueryValidatorUnbondingDelegationsResponse_1_list{list: &x.UnbondingResponses}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.unbonding_responses":
		list := []*UnbondingDelegation{}
		return protoreflect.ValueOfList(&_QueryValidatorUnbondingDelegationsResponse_1_list{list: &list})
	case "cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryValidatorUnbondingDelegationsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryValidatorUnbondingDelegationsResponse)
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
		if len(x.UnbondingResponses) > 0 {
			for _, e := range x.UnbondingResponses {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryValidatorUnbondingDelegationsResponse)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.UnbondingResponses) > 0 {
			for iNdEx := len(x.UnbondingResponses) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.UnbondingResponses[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*QueryValidatorUnbondingDelegationsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorUnbondingDelegationsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryValidatorUnbondingDelegationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UnbondingResponses", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.UnbondingResponses = append(x.UnbondingResponses, &UnbondingDelegation{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.UnbondingResponses[len(x.UnbondingResponses)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryDelegationRequest                protoreflect.MessageDescriptor
	fd_QueryDelegationRequest_delegator_addr protoreflect.FieldDescriptor
	fd_QueryDelegationRequest_validator_addr protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegationRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegationRequest")
	fd_QueryDelegationRequest_delegator_addr = md_QueryDelegationRequest.Fields().ByName("delegator_addr")
	fd_QueryDelegationRequest_validator_addr = md_QueryDelegationRequest.Fields().ByName("validator_addr")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegationRequest)(nil)

type fastReflection_QueryDelegationRequest QueryDelegationRequest

func (x *QueryDelegationRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegationRequest)(x)
}

func (x *QueryDelegationRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegationRequest_messageType fastReflection_QueryDelegationRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegationRequest_messageType{}

type fastReflection_QueryDelegationRequest_messageType struct{}

func (x fastReflection_QueryDelegationRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegationRequest)(nil)
}
func (x fastReflection_QueryDelegationRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationRequest)
}
func (x fastReflection_QueryDelegationRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegationRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegationRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegationRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegationRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegationRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegationRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegationRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddr != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddr)
		if !f(fd_QueryDelegationRequest_delegator_addr, value) {
			return
		}
	}
	if x.ValidatorAddr != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddr)
		if !f(fd_QueryDelegationRequest_validator_addr, value) {
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
func (x *fastReflection_QueryDelegationRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationRequest.delegator_addr":
		return x.DelegatorAddr != ""
	case "cosmos.staking.v1beta1.QueryDelegationRequest.validator_addr":
		return x.ValidatorAddr != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationRequest.delegator_addr":
		x.DelegatorAddr = ""
	case "cosmos.staking.v1beta1.QueryDelegationRequest.validator_addr":
		x.ValidatorAddr = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegationRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationRequest.delegator_addr":
		value := x.DelegatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryDelegationRequest.validator_addr":
		value := x.ValidatorAddr
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegationRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationRequest.delegator_addr":
		x.DelegatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryDelegationRequest.validator_addr":
		x.ValidatorAddr = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegationRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationRequest.delegator_addr":
		panic(fmt.Errorf("field delegator_addr of message cosmos.staking.v1beta1.QueryDelegationRequest is not mutable"))
	case "cosmos.staking.v1beta1.QueryDelegationRequest.validator_addr":
		panic(fmt.Errorf("field validator_addr of message cosmos.staking.v1beta1.QueryDelegationRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegationRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationRequest.delegator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryDelegationRequest.validator_addr":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegationRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegationRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegationRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegationRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegationRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegationRequest)
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
		l = len(x.DelegatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddr)
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
		x := input.Message.Interface().(*QueryDelegationRequest)
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
		if len(x.ValidatorAddr) > 0 {
			i -= len(x.ValidatorAddr)
			copy(dAtA[i:], x.ValidatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddr)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddr) > 0 {
			i -= len(x.DelegatorAddr)
			copy(dAtA[i:], x.DelegatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddr)))
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
		x := input.Message.Interface().(*QueryDelegationRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddr", wireType)
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
				x.DelegatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
				x.ValidatorAddr = string(dAtA[iNdEx:postIndex])
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
	md_QueryDelegationResponse                     protoreflect.MessageDescriptor
	fd_QueryDelegationResponse_delegation_response protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegationResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegationResponse")
	fd_QueryDelegationResponse_delegation_response = md_QueryDelegationResponse.Fields().ByName("delegation_response")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegationResponse)(nil)

type fastReflection_QueryDelegationResponse QueryDelegationResponse

func (x *QueryDelegationResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegationResponse)(x)
}

func (x *QueryDelegationResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegationResponse_messageType fastReflection_QueryDelegationResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegationResponse_messageType{}

type fastReflection_QueryDelegationResponse_messageType struct{}

func (x fastReflection_QueryDelegationResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegationResponse)(nil)
}
func (x fastReflection_QueryDelegationResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationResponse)
}
func (x fastReflection_QueryDelegationResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegationResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegationResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegationResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegationResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegationResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegationResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegationResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegationResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegationResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegationResponse != nil {
		value := protoreflect.ValueOfMessage(x.DelegationResponse.ProtoReflect())
		if !f(fd_QueryDelegationResponse_delegation_response, value) {
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
func (x *fastReflection_QueryDelegationResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationResponse.delegation_response":
		return x.DelegationResponse != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationResponse.delegation_response":
		x.DelegationResponse = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegationResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationResponse.delegation_response":
		value := x.DelegationResponse
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegationResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationResponse.delegation_response":
		x.DelegationResponse = value.Message().Interface().(*DelegationResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegationResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationResponse.delegation_response":
		if x.DelegationResponse == nil {
			x.DelegationResponse = new(DelegationResponse)
		}
		return protoreflect.ValueOfMessage(x.DelegationResponse.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegationResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegationResponse.delegation_response":
		m := new(DelegationResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegationResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegationResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegationResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegationResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegationResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegationResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegationResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegationResponse)
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
		if x.DelegationResponse != nil {
			l = options.Size(x.DelegationResponse)
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
		x := input.Message.Interface().(*QueryDelegationResponse)
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
		if x.DelegationResponse != nil {
			encoded, err := options.Marshal(x.DelegationResponse)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*QueryDelegationResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegationResponse", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.DelegationResponse == nil {
					x.DelegationResponse = &DelegationResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DelegationResponse); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryUnbondingDelegationRequest                protoreflect.MessageDescriptor
	fd_QueryUnbondingDelegationRequest_delegator_addr protoreflect.FieldDescriptor
	fd_QueryUnbondingDelegationRequest_validator_addr protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryUnbondingDelegationRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryUnbondingDelegationRequest")
	fd_QueryUnbondingDelegationRequest_delegator_addr = md_QueryUnbondingDelegationRequest.Fields().ByName("delegator_addr")
	fd_QueryUnbondingDelegationRequest_validator_addr = md_QueryUnbondingDelegationRequest.Fields().ByName("validator_addr")
}

var _ protoreflect.Message = (*fastReflection_QueryUnbondingDelegationRequest)(nil)

type fastReflection_QueryUnbondingDelegationRequest QueryUnbondingDelegationRequest

func (x *QueryUnbondingDelegationRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryUnbondingDelegationRequest)(x)
}

func (x *QueryUnbondingDelegationRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryUnbondingDelegationRequest_messageType fastReflection_QueryUnbondingDelegationRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryUnbondingDelegationRequest_messageType{}

type fastReflection_QueryUnbondingDelegationRequest_messageType struct{}

func (x fastReflection_QueryUnbondingDelegationRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryUnbondingDelegationRequest)(nil)
}
func (x fastReflection_QueryUnbondingDelegationRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryUnbondingDelegationRequest)
}
func (x fastReflection_QueryUnbondingDelegationRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryUnbondingDelegationRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryUnbondingDelegationRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryUnbondingDelegationRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryUnbondingDelegationRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryUnbondingDelegationRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryUnbondingDelegationRequest) New() protoreflect.Message {
	return new(fastReflection_QueryUnbondingDelegationRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryUnbondingDelegationRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryUnbondingDelegationRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryUnbondingDelegationRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddr != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddr)
		if !f(fd_QueryUnbondingDelegationRequest_delegator_addr, value) {
			return
		}
	}
	if x.ValidatorAddr != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddr)
		if !f(fd_QueryUnbondingDelegationRequest_validator_addr, value) {
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
func (x *fastReflection_QueryUnbondingDelegationRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.delegator_addr":
		return x.DelegatorAddr != ""
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.validator_addr":
		return x.ValidatorAddr != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryUnbondingDelegationRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.delegator_addr":
		x.DelegatorAddr = ""
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.validator_addr":
		x.ValidatorAddr = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryUnbondingDelegationRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.delegator_addr":
		value := x.DelegatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.validator_addr":
		value := x.ValidatorAddr
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryUnbondingDelegationRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.delegator_addr":
		x.DelegatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.validator_addr":
		x.ValidatorAddr = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryUnbondingDelegationRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.delegator_addr":
		panic(fmt.Errorf("field delegator_addr of message cosmos.staking.v1beta1.QueryUnbondingDelegationRequest is not mutable"))
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.validator_addr":
		panic(fmt.Errorf("field validator_addr of message cosmos.staking.v1beta1.QueryUnbondingDelegationRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryUnbondingDelegationRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.delegator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationRequest.validator_addr":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryUnbondingDelegationRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryUnbondingDelegationRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryUnbondingDelegationRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryUnbondingDelegationRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryUnbondingDelegationRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryUnbondingDelegationRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryUnbondingDelegationRequest)
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
		l = len(x.DelegatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddr)
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
		x := input.Message.Interface().(*QueryUnbondingDelegationRequest)
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
		if len(x.ValidatorAddr) > 0 {
			i -= len(x.ValidatorAddr)
			copy(dAtA[i:], x.ValidatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddr)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddr) > 0 {
			i -= len(x.DelegatorAddr)
			copy(dAtA[i:], x.DelegatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddr)))
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
		x := input.Message.Interface().(*QueryUnbondingDelegationRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryUnbondingDelegationRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryUnbondingDelegationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddr", wireType)
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
				x.DelegatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
				x.ValidatorAddr = string(dAtA[iNdEx:postIndex])
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
	md_QueryUnbondingDelegationResponse        protoreflect.MessageDescriptor
	fd_QueryUnbondingDelegationResponse_unbond protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryUnbondingDelegationResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryUnbondingDelegationResponse")
	fd_QueryUnbondingDelegationResponse_unbond = md_QueryUnbondingDelegationResponse.Fields().ByName("unbond")
}

var _ protoreflect.Message = (*fastReflection_QueryUnbondingDelegationResponse)(nil)

type fastReflection_QueryUnbondingDelegationResponse QueryUnbondingDelegationResponse

func (x *QueryUnbondingDelegationResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryUnbondingDelegationResponse)(x)
}

func (x *QueryUnbondingDelegationResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryUnbondingDelegationResponse_messageType fastReflection_QueryUnbondingDelegationResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryUnbondingDelegationResponse_messageType{}

type fastReflection_QueryUnbondingDelegationResponse_messageType struct{}

func (x fastReflection_QueryUnbondingDelegationResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryUnbondingDelegationResponse)(nil)
}
func (x fastReflection_QueryUnbondingDelegationResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryUnbondingDelegationResponse)
}
func (x fastReflection_QueryUnbondingDelegationResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryUnbondingDelegationResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryUnbondingDelegationResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryUnbondingDelegationResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryUnbondingDelegationResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryUnbondingDelegationResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryUnbondingDelegationResponse) New() protoreflect.Message {
	return new(fastReflection_QueryUnbondingDelegationResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryUnbondingDelegationResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryUnbondingDelegationResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryUnbondingDelegationResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Unbond != nil {
		value := protoreflect.ValueOfMessage(x.Unbond.ProtoReflect())
		if !f(fd_QueryUnbondingDelegationResponse_unbond, value) {
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
func (x *fastReflection_QueryUnbondingDelegationResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationResponse.unbond":
		return x.Unbond != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryUnbondingDelegationResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationResponse.unbond":
		x.Unbond = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryUnbondingDelegationResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationResponse.unbond":
		value := x.Unbond
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryUnbondingDelegationResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationResponse.unbond":
		x.Unbond = value.Message().Interface().(*UnbondingDelegation)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryUnbondingDelegationResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationResponse.unbond":
		if x.Unbond == nil {
			x.Unbond = new(UnbondingDelegation)
		}
		return protoreflect.ValueOfMessage(x.Unbond.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryUnbondingDelegationResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryUnbondingDelegationResponse.unbond":
		m := new(UnbondingDelegation)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryUnbondingDelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryUnbondingDelegationResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryUnbondingDelegationResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryUnbondingDelegationResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryUnbondingDelegationResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryUnbondingDelegationResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryUnbondingDelegationResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryUnbondingDelegationResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryUnbondingDelegationResponse)
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
		if x.Unbond != nil {
			l = options.Size(x.Unbond)
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
		x := input.Message.Interface().(*QueryUnbondingDelegationResponse)
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
		if x.Unbond != nil {
			encoded, err := options.Marshal(x.Unbond)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*QueryUnbondingDelegationResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryUnbondingDelegationResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryUnbondingDelegationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Unbond", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Unbond == nil {
					x.Unbond = &UnbondingDelegation{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Unbond); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryDelegatorDelegationsRequest                protoreflect.MessageDescriptor
	fd_QueryDelegatorDelegationsRequest_delegator_addr protoreflect.FieldDescriptor
	fd_QueryDelegatorDelegationsRequest_pagination     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegatorDelegationsRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegatorDelegationsRequest")
	fd_QueryDelegatorDelegationsRequest_delegator_addr = md_QueryDelegatorDelegationsRequest.Fields().ByName("delegator_addr")
	fd_QueryDelegatorDelegationsRequest_pagination = md_QueryDelegatorDelegationsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorDelegationsRequest)(nil)

type fastReflection_QueryDelegatorDelegationsRequest QueryDelegatorDelegationsRequest

func (x *QueryDelegatorDelegationsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorDelegationsRequest)(x)
}

func (x *QueryDelegatorDelegationsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorDelegationsRequest_messageType fastReflection_QueryDelegatorDelegationsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorDelegationsRequest_messageType{}

type fastReflection_QueryDelegatorDelegationsRequest_messageType struct{}

func (x fastReflection_QueryDelegatorDelegationsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorDelegationsRequest)(nil)
}
func (x fastReflection_QueryDelegatorDelegationsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorDelegationsRequest)
}
func (x fastReflection_QueryDelegatorDelegationsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorDelegationsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorDelegationsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorDelegationsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorDelegationsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorDelegationsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorDelegationsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorDelegationsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorDelegationsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorDelegationsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorDelegationsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddr != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddr)
		if !f(fd_QueryDelegatorDelegationsRequest_delegator_addr, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryDelegatorDelegationsRequest_pagination, value) {
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
func (x *fastReflection_QueryDelegatorDelegationsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.delegator_addr":
		return x.DelegatorAddr != ""
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorDelegationsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.delegator_addr":
		x.DelegatorAddr = ""
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorDelegationsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.delegator_addr":
		value := x.DelegatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegatorDelegationsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.delegator_addr":
		x.DelegatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegatorDelegationsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.delegator_addr":
		panic(fmt.Errorf("field delegator_addr of message cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorDelegationsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.delegator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorDelegationsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorDelegationsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorDelegationsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegatorDelegationsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorDelegationsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorDelegationsRequest)
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
		l = len(x.DelegatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryDelegatorDelegationsRequest)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddr) > 0 {
			i -= len(x.DelegatorAddr)
			copy(dAtA[i:], x.DelegatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddr)))
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
		x := input.Message.Interface().(*QueryDelegatorDelegationsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorDelegationsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorDelegationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddr", wireType)
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
				x.DelegatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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

var _ protoreflect.List = (*_QueryDelegatorDelegationsResponse_1_list)(nil)

type _QueryDelegatorDelegationsResponse_1_list struct {
	list *[]*DelegationResponse
}

func (x *_QueryDelegatorDelegationsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryDelegatorDelegationsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryDelegatorDelegationsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegationResponse)
	(*x.list)[i] = concreteValue
}

func (x *_QueryDelegatorDelegationsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DelegationResponse)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryDelegatorDelegationsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(DelegationResponse)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegatorDelegationsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryDelegatorDelegationsResponse_1_list) NewElement() protoreflect.Value {
	v := new(DelegationResponse)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegatorDelegationsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryDelegatorDelegationsResponse                      protoreflect.MessageDescriptor
	fd_QueryDelegatorDelegationsResponse_delegation_responses protoreflect.FieldDescriptor
	fd_QueryDelegatorDelegationsResponse_pagination           protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegatorDelegationsResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegatorDelegationsResponse")
	fd_QueryDelegatorDelegationsResponse_delegation_responses = md_QueryDelegatorDelegationsResponse.Fields().ByName("delegation_responses")
	fd_QueryDelegatorDelegationsResponse_pagination = md_QueryDelegatorDelegationsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorDelegationsResponse)(nil)

type fastReflection_QueryDelegatorDelegationsResponse QueryDelegatorDelegationsResponse

func (x *QueryDelegatorDelegationsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorDelegationsResponse)(x)
}

func (x *QueryDelegatorDelegationsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorDelegationsResponse_messageType fastReflection_QueryDelegatorDelegationsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorDelegationsResponse_messageType{}

type fastReflection_QueryDelegatorDelegationsResponse_messageType struct{}

func (x fastReflection_QueryDelegatorDelegationsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorDelegationsResponse)(nil)
}
func (x fastReflection_QueryDelegatorDelegationsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorDelegationsResponse)
}
func (x fastReflection_QueryDelegatorDelegationsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorDelegationsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorDelegationsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorDelegationsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorDelegationsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorDelegationsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorDelegationsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorDelegationsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorDelegationsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorDelegationsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorDelegationsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.DelegationResponses) != 0 {
		value := protoreflect.ValueOfList(&_QueryDelegatorDelegationsResponse_1_list{list: &x.DelegationResponses})
		if !f(fd_QueryDelegatorDelegationsResponse_delegation_responses, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryDelegatorDelegationsResponse_pagination, value) {
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
func (x *fastReflection_QueryDelegatorDelegationsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.delegation_responses":
		return len(x.DelegationResponses) != 0
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorDelegationsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.delegation_responses":
		x.DelegationResponses = nil
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorDelegationsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.delegation_responses":
		if len(x.DelegationResponses) == 0 {
			return protoreflect.ValueOfList(&_QueryDelegatorDelegationsResponse_1_list{})
		}
		listValue := &_QueryDelegatorDelegationsResponse_1_list{list: &x.DelegationResponses}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegatorDelegationsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.delegation_responses":
		lv := value.List()
		clv := lv.(*_QueryDelegatorDelegationsResponse_1_list)
		x.DelegationResponses = *clv.list
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegatorDelegationsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.delegation_responses":
		if x.DelegationResponses == nil {
			x.DelegationResponses = []*DelegationResponse{}
		}
		value := &_QueryDelegatorDelegationsResponse_1_list{list: &x.DelegationResponses}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorDelegationsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.delegation_responses":
		list := []*DelegationResponse{}
		return protoreflect.ValueOfList(&_QueryDelegatorDelegationsResponse_1_list{list: &list})
	case "cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorDelegationsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorDelegationsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorDelegationsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegatorDelegationsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorDelegationsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorDelegationsResponse)
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
		if len(x.DelegationResponses) > 0 {
			for _, e := range x.DelegationResponses {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryDelegatorDelegationsResponse)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegationResponses) > 0 {
			for iNdEx := len(x.DelegationResponses) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.DelegationResponses[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*QueryDelegatorDelegationsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorDelegationsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorDelegationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegationResponses", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegationResponses = append(x.DelegationResponses, &DelegationResponse{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DelegationResponses[len(x.DelegationResponses)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryDelegatorUnbondingDelegationsRequest                protoreflect.MessageDescriptor
	fd_QueryDelegatorUnbondingDelegationsRequest_delegator_addr protoreflect.FieldDescriptor
	fd_QueryDelegatorUnbondingDelegationsRequest_pagination     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegatorUnbondingDelegationsRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegatorUnbondingDelegationsRequest")
	fd_QueryDelegatorUnbondingDelegationsRequest_delegator_addr = md_QueryDelegatorUnbondingDelegationsRequest.Fields().ByName("delegator_addr")
	fd_QueryDelegatorUnbondingDelegationsRequest_pagination = md_QueryDelegatorUnbondingDelegationsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorUnbondingDelegationsRequest)(nil)

type fastReflection_QueryDelegatorUnbondingDelegationsRequest QueryDelegatorUnbondingDelegationsRequest

func (x *QueryDelegatorUnbondingDelegationsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorUnbondingDelegationsRequest)(x)
}

func (x *QueryDelegatorUnbondingDelegationsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorUnbondingDelegationsRequest_messageType fastReflection_QueryDelegatorUnbondingDelegationsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorUnbondingDelegationsRequest_messageType{}

type fastReflection_QueryDelegatorUnbondingDelegationsRequest_messageType struct{}

func (x fastReflection_QueryDelegatorUnbondingDelegationsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorUnbondingDelegationsRequest)(nil)
}
func (x fastReflection_QueryDelegatorUnbondingDelegationsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorUnbondingDelegationsRequest)
}
func (x fastReflection_QueryDelegatorUnbondingDelegationsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorUnbondingDelegationsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorUnbondingDelegationsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorUnbondingDelegationsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorUnbondingDelegationsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorUnbondingDelegationsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddr != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddr)
		if !f(fd_QueryDelegatorUnbondingDelegationsRequest_delegator_addr, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryDelegatorUnbondingDelegationsRequest_pagination, value) {
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
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.delegator_addr":
		return x.DelegatorAddr != ""
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.delegator_addr":
		x.DelegatorAddr = ""
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.delegator_addr":
		value := x.DelegatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.delegator_addr":
		x.DelegatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.delegator_addr":
		panic(fmt.Errorf("field delegator_addr of message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.delegator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorUnbondingDelegationsRequest)
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
		l = len(x.DelegatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryDelegatorUnbondingDelegationsRequest)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddr) > 0 {
			i -= len(x.DelegatorAddr)
			copy(dAtA[i:], x.DelegatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddr)))
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
		x := input.Message.Interface().(*QueryDelegatorUnbondingDelegationsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorUnbondingDelegationsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorUnbondingDelegationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddr", wireType)
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
				x.DelegatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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

var _ protoreflect.List = (*_QueryDelegatorUnbondingDelegationsResponse_1_list)(nil)

type _QueryDelegatorUnbondingDelegationsResponse_1_list struct {
	list *[]*UnbondingDelegation
}

func (x *_QueryDelegatorUnbondingDelegationsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryDelegatorUnbondingDelegationsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryDelegatorUnbondingDelegationsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*UnbondingDelegation)
	(*x.list)[i] = concreteValue
}

func (x *_QueryDelegatorUnbondingDelegationsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*UnbondingDelegation)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryDelegatorUnbondingDelegationsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(UnbondingDelegation)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegatorUnbondingDelegationsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryDelegatorUnbondingDelegationsResponse_1_list) NewElement() protoreflect.Value {
	v := new(UnbondingDelegation)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegatorUnbondingDelegationsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryDelegatorUnbondingDelegationsResponse                     protoreflect.MessageDescriptor
	fd_QueryDelegatorUnbondingDelegationsResponse_unbonding_responses protoreflect.FieldDescriptor
	fd_QueryDelegatorUnbondingDelegationsResponse_pagination          protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegatorUnbondingDelegationsResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegatorUnbondingDelegationsResponse")
	fd_QueryDelegatorUnbondingDelegationsResponse_unbonding_responses = md_QueryDelegatorUnbondingDelegationsResponse.Fields().ByName("unbonding_responses")
	fd_QueryDelegatorUnbondingDelegationsResponse_pagination = md_QueryDelegatorUnbondingDelegationsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorUnbondingDelegationsResponse)(nil)

type fastReflection_QueryDelegatorUnbondingDelegationsResponse QueryDelegatorUnbondingDelegationsResponse

func (x *QueryDelegatorUnbondingDelegationsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorUnbondingDelegationsResponse)(x)
}

func (x *QueryDelegatorUnbondingDelegationsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorUnbondingDelegationsResponse_messageType fastReflection_QueryDelegatorUnbondingDelegationsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorUnbondingDelegationsResponse_messageType{}

type fastReflection_QueryDelegatorUnbondingDelegationsResponse_messageType struct{}

func (x fastReflection_QueryDelegatorUnbondingDelegationsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorUnbondingDelegationsResponse)(nil)
}
func (x fastReflection_QueryDelegatorUnbondingDelegationsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorUnbondingDelegationsResponse)
}
func (x fastReflection_QueryDelegatorUnbondingDelegationsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorUnbondingDelegationsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorUnbondingDelegationsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorUnbondingDelegationsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorUnbondingDelegationsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorUnbondingDelegationsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.UnbondingResponses) != 0 {
		value := protoreflect.ValueOfList(&_QueryDelegatorUnbondingDelegationsResponse_1_list{list: &x.UnbondingResponses})
		if !f(fd_QueryDelegatorUnbondingDelegationsResponse_unbonding_responses, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryDelegatorUnbondingDelegationsResponse_pagination, value) {
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
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.unbonding_responses":
		return len(x.UnbondingResponses) != 0
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.unbonding_responses":
		x.UnbondingResponses = nil
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.unbonding_responses":
		if len(x.UnbondingResponses) == 0 {
			return protoreflect.ValueOfList(&_QueryDelegatorUnbondingDelegationsResponse_1_list{})
		}
		listValue := &_QueryDelegatorUnbondingDelegationsResponse_1_list{list: &x.UnbondingResponses}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.unbonding_responses":
		lv := value.List()
		clv := lv.(*_QueryDelegatorUnbondingDelegationsResponse_1_list)
		x.UnbondingResponses = *clv.list
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.unbonding_responses":
		if x.UnbondingResponses == nil {
			x.UnbondingResponses = []*UnbondingDelegation{}
		}
		value := &_QueryDelegatorUnbondingDelegationsResponse_1_list{list: &x.UnbondingResponses}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.unbonding_responses":
		list := []*UnbondingDelegation{}
		return protoreflect.ValueOfList(&_QueryDelegatorUnbondingDelegationsResponse_1_list{list: &list})
	case "cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorUnbondingDelegationsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorUnbondingDelegationsResponse)
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
		if len(x.UnbondingResponses) > 0 {
			for _, e := range x.UnbondingResponses {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryDelegatorUnbondingDelegationsResponse)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.UnbondingResponses) > 0 {
			for iNdEx := len(x.UnbondingResponses) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.UnbondingResponses[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*QueryDelegatorUnbondingDelegationsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorUnbondingDelegationsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorUnbondingDelegationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UnbondingResponses", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.UnbondingResponses = append(x.UnbondingResponses, &UnbondingDelegation{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.UnbondingResponses[len(x.UnbondingResponses)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryRedelegationsRequest                    protoreflect.MessageDescriptor
	fd_QueryRedelegationsRequest_delegator_addr     protoreflect.FieldDescriptor
	fd_QueryRedelegationsRequest_src_validator_addr protoreflect.FieldDescriptor
	fd_QueryRedelegationsRequest_dst_validator_addr protoreflect.FieldDescriptor
	fd_QueryRedelegationsRequest_pagination         protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryRedelegationsRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryRedelegationsRequest")
	fd_QueryRedelegationsRequest_delegator_addr = md_QueryRedelegationsRequest.Fields().ByName("delegator_addr")
	fd_QueryRedelegationsRequest_src_validator_addr = md_QueryRedelegationsRequest.Fields().ByName("src_validator_addr")
	fd_QueryRedelegationsRequest_dst_validator_addr = md_QueryRedelegationsRequest.Fields().ByName("dst_validator_addr")
	fd_QueryRedelegationsRequest_pagination = md_QueryRedelegationsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryRedelegationsRequest)(nil)

type fastReflection_QueryRedelegationsRequest QueryRedelegationsRequest

func (x *QueryRedelegationsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryRedelegationsRequest)(x)
}

func (x *QueryRedelegationsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryRedelegationsRequest_messageType fastReflection_QueryRedelegationsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryRedelegationsRequest_messageType{}

type fastReflection_QueryRedelegationsRequest_messageType struct{}

func (x fastReflection_QueryRedelegationsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryRedelegationsRequest)(nil)
}
func (x fastReflection_QueryRedelegationsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryRedelegationsRequest)
}
func (x fastReflection_QueryRedelegationsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryRedelegationsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryRedelegationsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryRedelegationsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryRedelegationsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryRedelegationsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryRedelegationsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryRedelegationsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryRedelegationsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryRedelegationsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryRedelegationsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddr != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddr)
		if !f(fd_QueryRedelegationsRequest_delegator_addr, value) {
			return
		}
	}
	if x.SrcValidatorAddr != "" {
		value := protoreflect.ValueOfString(x.SrcValidatorAddr)
		if !f(fd_QueryRedelegationsRequest_src_validator_addr, value) {
			return
		}
	}
	if x.DstValidatorAddr != "" {
		value := protoreflect.ValueOfString(x.DstValidatorAddr)
		if !f(fd_QueryRedelegationsRequest_dst_validator_addr, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryRedelegationsRequest_pagination, value) {
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
func (x *fastReflection_QueryRedelegationsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.delegator_addr":
		return x.DelegatorAddr != ""
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.src_validator_addr":
		return x.SrcValidatorAddr != ""
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.dst_validator_addr":
		return x.DstValidatorAddr != ""
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryRedelegationsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.delegator_addr":
		x.DelegatorAddr = ""
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.src_validator_addr":
		x.SrcValidatorAddr = ""
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.dst_validator_addr":
		x.DstValidatorAddr = ""
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryRedelegationsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.delegator_addr":
		value := x.DelegatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.src_validator_addr":
		value := x.SrcValidatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.dst_validator_addr":
		value := x.DstValidatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryRedelegationsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.delegator_addr":
		x.DelegatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.src_validator_addr":
		x.SrcValidatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.dst_validator_addr":
		x.DstValidatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryRedelegationsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.delegator_addr":
		panic(fmt.Errorf("field delegator_addr of message cosmos.staking.v1beta1.QueryRedelegationsRequest is not mutable"))
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.src_validator_addr":
		panic(fmt.Errorf("field src_validator_addr of message cosmos.staking.v1beta1.QueryRedelegationsRequest is not mutable"))
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.dst_validator_addr":
		panic(fmt.Errorf("field dst_validator_addr of message cosmos.staking.v1beta1.QueryRedelegationsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryRedelegationsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.delegator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.src_validator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.dst_validator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryRedelegationsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryRedelegationsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryRedelegationsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryRedelegationsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryRedelegationsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryRedelegationsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryRedelegationsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryRedelegationsRequest)
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
		l = len(x.DelegatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.SrcValidatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DstValidatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryRedelegationsRequest)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.DstValidatorAddr) > 0 {
			i -= len(x.DstValidatorAddr)
			copy(dAtA[i:], x.DstValidatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DstValidatorAddr)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.SrcValidatorAddr) > 0 {
			i -= len(x.SrcValidatorAddr)
			copy(dAtA[i:], x.SrcValidatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SrcValidatorAddr)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddr) > 0 {
			i -= len(x.DelegatorAddr)
			copy(dAtA[i:], x.DelegatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddr)))
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
		x := input.Message.Interface().(*QueryRedelegationsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryRedelegationsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryRedelegationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddr", wireType)
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
				x.DelegatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SrcValidatorAddr", wireType)
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
				x.SrcValidatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DstValidatorAddr", wireType)
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
				x.DstValidatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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

var _ protoreflect.List = (*_QueryRedelegationsResponse_1_list)(nil)

type _QueryRedelegationsResponse_1_list struct {
	list *[]*RedelegationResponse
}

func (x *_QueryRedelegationsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryRedelegationsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryRedelegationsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*RedelegationResponse)
	(*x.list)[i] = concreteValue
}

func (x *_QueryRedelegationsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*RedelegationResponse)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryRedelegationsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(RedelegationResponse)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryRedelegationsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryRedelegationsResponse_1_list) NewElement() protoreflect.Value {
	v := new(RedelegationResponse)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryRedelegationsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryRedelegationsResponse                        protoreflect.MessageDescriptor
	fd_QueryRedelegationsResponse_redelegation_responses protoreflect.FieldDescriptor
	fd_QueryRedelegationsResponse_pagination             protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryRedelegationsResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryRedelegationsResponse")
	fd_QueryRedelegationsResponse_redelegation_responses = md_QueryRedelegationsResponse.Fields().ByName("redelegation_responses")
	fd_QueryRedelegationsResponse_pagination = md_QueryRedelegationsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryRedelegationsResponse)(nil)

type fastReflection_QueryRedelegationsResponse QueryRedelegationsResponse

func (x *QueryRedelegationsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryRedelegationsResponse)(x)
}

func (x *QueryRedelegationsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryRedelegationsResponse_messageType fastReflection_QueryRedelegationsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryRedelegationsResponse_messageType{}

type fastReflection_QueryRedelegationsResponse_messageType struct{}

func (x fastReflection_QueryRedelegationsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryRedelegationsResponse)(nil)
}
func (x fastReflection_QueryRedelegationsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryRedelegationsResponse)
}
func (x fastReflection_QueryRedelegationsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryRedelegationsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryRedelegationsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryRedelegationsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryRedelegationsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryRedelegationsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryRedelegationsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryRedelegationsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryRedelegationsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryRedelegationsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryRedelegationsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.RedelegationResponses) != 0 {
		value := protoreflect.ValueOfList(&_QueryRedelegationsResponse_1_list{list: &x.RedelegationResponses})
		if !f(fd_QueryRedelegationsResponse_redelegation_responses, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryRedelegationsResponse_pagination, value) {
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
func (x *fastReflection_QueryRedelegationsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.redelegation_responses":
		return len(x.RedelegationResponses) != 0
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryRedelegationsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.redelegation_responses":
		x.RedelegationResponses = nil
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryRedelegationsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.redelegation_responses":
		if len(x.RedelegationResponses) == 0 {
			return protoreflect.ValueOfList(&_QueryRedelegationsResponse_1_list{})
		}
		listValue := &_QueryRedelegationsResponse_1_list{list: &x.RedelegationResponses}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryRedelegationsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.redelegation_responses":
		lv := value.List()
		clv := lv.(*_QueryRedelegationsResponse_1_list)
		x.RedelegationResponses = *clv.list
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryRedelegationsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.redelegation_responses":
		if x.RedelegationResponses == nil {
			x.RedelegationResponses = []*RedelegationResponse{}
		}
		value := &_QueryRedelegationsResponse_1_list{list: &x.RedelegationResponses}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryRedelegationsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.redelegation_responses":
		list := []*RedelegationResponse{}
		return protoreflect.ValueOfList(&_QueryRedelegationsResponse_1_list{list: &list})
	case "cosmos.staking.v1beta1.QueryRedelegationsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryRedelegationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryRedelegationsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryRedelegationsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryRedelegationsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryRedelegationsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryRedelegationsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryRedelegationsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryRedelegationsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryRedelegationsResponse)
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
		if len(x.RedelegationResponses) > 0 {
			for _, e := range x.RedelegationResponses {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryRedelegationsResponse)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.RedelegationResponses) > 0 {
			for iNdEx := len(x.RedelegationResponses) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.RedelegationResponses[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*QueryRedelegationsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryRedelegationsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryRedelegationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RedelegationResponses", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RedelegationResponses = append(x.RedelegationResponses, &RedelegationResponse{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.RedelegationResponses[len(x.RedelegationResponses)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryDelegatorValidatorsRequest                protoreflect.MessageDescriptor
	fd_QueryDelegatorValidatorsRequest_delegator_addr protoreflect.FieldDescriptor
	fd_QueryDelegatorValidatorsRequest_pagination     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegatorValidatorsRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegatorValidatorsRequest")
	fd_QueryDelegatorValidatorsRequest_delegator_addr = md_QueryDelegatorValidatorsRequest.Fields().ByName("delegator_addr")
	fd_QueryDelegatorValidatorsRequest_pagination = md_QueryDelegatorValidatorsRequest.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorValidatorsRequest)(nil)

type fastReflection_QueryDelegatorValidatorsRequest QueryDelegatorValidatorsRequest

func (x *QueryDelegatorValidatorsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorsRequest)(x)
}

func (x *QueryDelegatorValidatorsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorValidatorsRequest_messageType fastReflection_QueryDelegatorValidatorsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorValidatorsRequest_messageType{}

type fastReflection_QueryDelegatorValidatorsRequest_messageType struct{}

func (x fastReflection_QueryDelegatorValidatorsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorsRequest)(nil)
}
func (x fastReflection_QueryDelegatorValidatorsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorsRequest)
}
func (x fastReflection_QueryDelegatorValidatorsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorValidatorsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorValidatorsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorValidatorsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddr != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddr)
		if !f(fd_QueryDelegatorValidatorsRequest_delegator_addr, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryDelegatorValidatorsRequest_pagination, value) {
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
func (x *fastReflection_QueryDelegatorValidatorsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.delegator_addr":
		return x.DelegatorAddr != ""
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.delegator_addr":
		x.DelegatorAddr = ""
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorValidatorsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.delegator_addr":
		value := x.DelegatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegatorValidatorsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.delegator_addr":
		x.DelegatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageRequest)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegatorValidatorsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageRequest)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.delegator_addr":
		panic(fmt.Errorf("field delegator_addr of message cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorValidatorsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.delegator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.pagination":
		m := new(v1beta1.PageRequest)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorValidatorsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorValidatorsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegatorValidatorsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorValidatorsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorValidatorsRequest)
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
		l = len(x.DelegatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryDelegatorValidatorsRequest)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddr) > 0 {
			i -= len(x.DelegatorAddr)
			copy(dAtA[i:], x.DelegatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddr)))
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
		x := input.Message.Interface().(*QueryDelegatorValidatorsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddr", wireType)
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
				x.DelegatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageRequest{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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

var _ protoreflect.List = (*_QueryDelegatorValidatorsResponse_1_list)(nil)

type _QueryDelegatorValidatorsResponse_1_list struct {
	list *[]*Validator
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	(*x.list)[i] = concreteValue
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	*x.list = append(*x.list, concreteValue)
}

func (x *_QueryDelegatorValidatorsResponse_1_list) AppendMutable() protoreflect.Value {
	v := new(Validator)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegatorValidatorsResponse_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_QueryDelegatorValidatorsResponse_1_list) NewElement() protoreflect.Value {
	v := new(Validator)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_QueryDelegatorValidatorsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_QueryDelegatorValidatorsResponse            protoreflect.MessageDescriptor
	fd_QueryDelegatorValidatorsResponse_validators protoreflect.FieldDescriptor
	fd_QueryDelegatorValidatorsResponse_pagination protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegatorValidatorsResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegatorValidatorsResponse")
	fd_QueryDelegatorValidatorsResponse_validators = md_QueryDelegatorValidatorsResponse.Fields().ByName("validators")
	fd_QueryDelegatorValidatorsResponse_pagination = md_QueryDelegatorValidatorsResponse.Fields().ByName("pagination")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorValidatorsResponse)(nil)

type fastReflection_QueryDelegatorValidatorsResponse QueryDelegatorValidatorsResponse

func (x *QueryDelegatorValidatorsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorsResponse)(x)
}

func (x *QueryDelegatorValidatorsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorValidatorsResponse_messageType fastReflection_QueryDelegatorValidatorsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorValidatorsResponse_messageType{}

type fastReflection_QueryDelegatorValidatorsResponse_messageType struct{}

func (x fastReflection_QueryDelegatorValidatorsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorsResponse)(nil)
}
func (x fastReflection_QueryDelegatorValidatorsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorsResponse)
}
func (x fastReflection_QueryDelegatorValidatorsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorValidatorsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorValidatorsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorValidatorsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Validators) != 0 {
		value := protoreflect.ValueOfList(&_QueryDelegatorValidatorsResponse_1_list{list: &x.Validators})
		if !f(fd_QueryDelegatorValidatorsResponse_validators, value) {
			return
		}
	}
	if x.Pagination != nil {
		value := protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
		if !f(fd_QueryDelegatorValidatorsResponse_pagination, value) {
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
func (x *fastReflection_QueryDelegatorValidatorsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.validators":
		return len(x.Validators) != 0
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.pagination":
		return x.Pagination != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.validators":
		x.Validators = nil
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.pagination":
		x.Pagination = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorValidatorsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.validators":
		if len(x.Validators) == 0 {
			return protoreflect.ValueOfList(&_QueryDelegatorValidatorsResponse_1_list{})
		}
		listValue := &_QueryDelegatorValidatorsResponse_1_list{list: &x.Validators}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.pagination":
		value := x.Pagination
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegatorValidatorsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.validators":
		lv := value.List()
		clv := lv.(*_QueryDelegatorValidatorsResponse_1_list)
		x.Validators = *clv.list
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.pagination":
		x.Pagination = value.Message().Interface().(*v1beta1.PageResponse)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegatorValidatorsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.validators":
		if x.Validators == nil {
			x.Validators = []*Validator{}
		}
		value := &_QueryDelegatorValidatorsResponse_1_list{list: &x.Validators}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.pagination":
		if x.Pagination == nil {
			x.Pagination = new(v1beta1.PageResponse)
		}
		return protoreflect.ValueOfMessage(x.Pagination.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorValidatorsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.validators":
		list := []*Validator{}
		return protoreflect.ValueOfList(&_QueryDelegatorValidatorsResponse_1_list{list: &list})
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.pagination":
		m := new(v1beta1.PageResponse)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorValidatorsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorValidatorsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegatorValidatorsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorValidatorsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorValidatorsResponse)
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
		if len(x.Validators) > 0 {
			for _, e := range x.Validators {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Pagination != nil {
			l = options.Size(x.Pagination)
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
		x := input.Message.Interface().(*QueryDelegatorValidatorsResponse)
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
		if x.Pagination != nil {
			encoded, err := options.Marshal(x.Pagination)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Validators) > 0 {
			for iNdEx := len(x.Validators) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Validators[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*QueryDelegatorValidatorsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Validators = append(x.Validators, &Validator{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validators[len(x.Validators)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pagination == nil {
					x.Pagination = &v1beta1.PageResponse{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pagination); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryDelegatorValidatorRequest                protoreflect.MessageDescriptor
	fd_QueryDelegatorValidatorRequest_delegator_addr protoreflect.FieldDescriptor
	fd_QueryDelegatorValidatorRequest_validator_addr protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegatorValidatorRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegatorValidatorRequest")
	fd_QueryDelegatorValidatorRequest_delegator_addr = md_QueryDelegatorValidatorRequest.Fields().ByName("delegator_addr")
	fd_QueryDelegatorValidatorRequest_validator_addr = md_QueryDelegatorValidatorRequest.Fields().ByName("validator_addr")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorValidatorRequest)(nil)

type fastReflection_QueryDelegatorValidatorRequest QueryDelegatorValidatorRequest

func (x *QueryDelegatorValidatorRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorRequest)(x)
}

func (x *QueryDelegatorValidatorRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorValidatorRequest_messageType fastReflection_QueryDelegatorValidatorRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorValidatorRequest_messageType{}

type fastReflection_QueryDelegatorValidatorRequest_messageType struct{}

func (x fastReflection_QueryDelegatorValidatorRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorRequest)(nil)
}
func (x fastReflection_QueryDelegatorValidatorRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorRequest)
}
func (x fastReflection_QueryDelegatorValidatorRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorValidatorRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorValidatorRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorValidatorRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorValidatorRequest) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorValidatorRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorValidatorRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorValidatorRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddr != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddr)
		if !f(fd_QueryDelegatorValidatorRequest_delegator_addr, value) {
			return
		}
	}
	if x.ValidatorAddr != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddr)
		if !f(fd_QueryDelegatorValidatorRequest_validator_addr, value) {
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
func (x *fastReflection_QueryDelegatorValidatorRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.delegator_addr":
		return x.DelegatorAddr != ""
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.validator_addr":
		return x.ValidatorAddr != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.delegator_addr":
		x.DelegatorAddr = ""
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.validator_addr":
		x.ValidatorAddr = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorValidatorRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.delegator_addr":
		value := x.DelegatorAddr
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.validator_addr":
		value := x.ValidatorAddr
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegatorValidatorRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.delegator_addr":
		x.DelegatorAddr = value.Interface().(string)
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.validator_addr":
		x.ValidatorAddr = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegatorValidatorRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.delegator_addr":
		panic(fmt.Errorf("field delegator_addr of message cosmos.staking.v1beta1.QueryDelegatorValidatorRequest is not mutable"))
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.validator_addr":
		panic(fmt.Errorf("field validator_addr of message cosmos.staking.v1beta1.QueryDelegatorValidatorRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorValidatorRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.delegator_addr":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorRequest.validator_addr":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorValidatorRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegatorValidatorRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorValidatorRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegatorValidatorRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorValidatorRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorValidatorRequest)
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
		l = len(x.DelegatorAddr)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddr)
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
		x := input.Message.Interface().(*QueryDelegatorValidatorRequest)
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
		if len(x.ValidatorAddr) > 0 {
			i -= len(x.ValidatorAddr)
			copy(dAtA[i:], x.ValidatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddr)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddr) > 0 {
			i -= len(x.DelegatorAddr)
			copy(dAtA[i:], x.DelegatorAddr)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddr)))
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
		x := input.Message.Interface().(*QueryDelegatorValidatorRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddr", wireType)
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
				x.DelegatorAddr = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
				x.ValidatorAddr = string(dAtA[iNdEx:postIndex])
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
	md_QueryDelegatorValidatorResponse           protoreflect.MessageDescriptor
	fd_QueryDelegatorValidatorResponse_validator protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryDelegatorValidatorResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryDelegatorValidatorResponse")
	fd_QueryDelegatorValidatorResponse_validator = md_QueryDelegatorValidatorResponse.Fields().ByName("validator")
}

var _ protoreflect.Message = (*fastReflection_QueryDelegatorValidatorResponse)(nil)

type fastReflection_QueryDelegatorValidatorResponse QueryDelegatorValidatorResponse

func (x *QueryDelegatorValidatorResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorResponse)(x)
}

func (x *QueryDelegatorValidatorResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryDelegatorValidatorResponse_messageType fastReflection_QueryDelegatorValidatorResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryDelegatorValidatorResponse_messageType{}

type fastReflection_QueryDelegatorValidatorResponse_messageType struct{}

func (x fastReflection_QueryDelegatorValidatorResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryDelegatorValidatorResponse)(nil)
}
func (x fastReflection_QueryDelegatorValidatorResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorResponse)
}
func (x fastReflection_QueryDelegatorValidatorResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryDelegatorValidatorResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryDelegatorValidatorResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryDelegatorValidatorResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryDelegatorValidatorResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryDelegatorValidatorResponse) New() protoreflect.Message {
	return new(fastReflection_QueryDelegatorValidatorResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryDelegatorValidatorResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryDelegatorValidatorResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryDelegatorValidatorResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Validator != nil {
		value := protoreflect.ValueOfMessage(x.Validator.ProtoReflect())
		if !f(fd_QueryDelegatorValidatorResponse_validator, value) {
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
func (x *fastReflection_QueryDelegatorValidatorResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorResponse.validator":
		return x.Validator != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorResponse.validator":
		x.Validator = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryDelegatorValidatorResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorResponse.validator":
		value := x.Validator
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryDelegatorValidatorResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorResponse.validator":
		x.Validator = value.Message().Interface().(*Validator)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryDelegatorValidatorResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorResponse.validator":
		if x.Validator == nil {
			x.Validator = new(Validator)
		}
		return protoreflect.ValueOfMessage(x.Validator.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryDelegatorValidatorResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryDelegatorValidatorResponse.validator":
		m := new(Validator)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryDelegatorValidatorResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryDelegatorValidatorResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryDelegatorValidatorResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryDelegatorValidatorResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryDelegatorValidatorResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryDelegatorValidatorResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryDelegatorValidatorResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryDelegatorValidatorResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryDelegatorValidatorResponse)
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
		if x.Validator != nil {
			l = options.Size(x.Validator)
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
		x := input.Message.Interface().(*QueryDelegatorValidatorResponse)
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
		if x.Validator != nil {
			encoded, err := options.Marshal(x.Validator)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*QueryDelegatorValidatorResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryDelegatorValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Validator == nil {
					x.Validator = &Validator{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validator); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryHistoricalInfoRequest        protoreflect.MessageDescriptor
	fd_QueryHistoricalInfoRequest_height protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryHistoricalInfoRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryHistoricalInfoRequest")
	fd_QueryHistoricalInfoRequest_height = md_QueryHistoricalInfoRequest.Fields().ByName("height")
}

var _ protoreflect.Message = (*fastReflection_QueryHistoricalInfoRequest)(nil)

type fastReflection_QueryHistoricalInfoRequest QueryHistoricalInfoRequest

func (x *QueryHistoricalInfoRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryHistoricalInfoRequest)(x)
}

func (x *QueryHistoricalInfoRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryHistoricalInfoRequest_messageType fastReflection_QueryHistoricalInfoRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryHistoricalInfoRequest_messageType{}

type fastReflection_QueryHistoricalInfoRequest_messageType struct{}

func (x fastReflection_QueryHistoricalInfoRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryHistoricalInfoRequest)(nil)
}
func (x fastReflection_QueryHistoricalInfoRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryHistoricalInfoRequest)
}
func (x fastReflection_QueryHistoricalInfoRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryHistoricalInfoRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryHistoricalInfoRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryHistoricalInfoRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryHistoricalInfoRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryHistoricalInfoRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryHistoricalInfoRequest) New() protoreflect.Message {
	return new(fastReflection_QueryHistoricalInfoRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryHistoricalInfoRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryHistoricalInfoRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryHistoricalInfoRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_QueryHistoricalInfoRequest_height, value) {
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
func (x *fastReflection_QueryHistoricalInfoRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoRequest.height":
		return x.Height != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryHistoricalInfoRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoRequest.height":
		x.Height = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryHistoricalInfoRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoRequest.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryHistoricalInfoRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoRequest.height":
		x.Height = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryHistoricalInfoRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoRequest.height":
		panic(fmt.Errorf("field height of message cosmos.staking.v1beta1.QueryHistoricalInfoRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryHistoricalInfoRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoRequest.height":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryHistoricalInfoRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryHistoricalInfoRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryHistoricalInfoRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryHistoricalInfoRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryHistoricalInfoRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryHistoricalInfoRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryHistoricalInfoRequest)
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
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
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
		x := input.Message.Interface().(*QueryHistoricalInfoRequest)
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
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*QueryHistoricalInfoRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryHistoricalInfoRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryHistoricalInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
	md_QueryHistoricalInfoResponse      protoreflect.MessageDescriptor
	fd_QueryHistoricalInfoResponse_hist protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryHistoricalInfoResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryHistoricalInfoResponse")
	fd_QueryHistoricalInfoResponse_hist = md_QueryHistoricalInfoResponse.Fields().ByName("hist")
}

var _ protoreflect.Message = (*fastReflection_QueryHistoricalInfoResponse)(nil)

type fastReflection_QueryHistoricalInfoResponse QueryHistoricalInfoResponse

func (x *QueryHistoricalInfoResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryHistoricalInfoResponse)(x)
}

func (x *QueryHistoricalInfoResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryHistoricalInfoResponse_messageType fastReflection_QueryHistoricalInfoResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryHistoricalInfoResponse_messageType{}

type fastReflection_QueryHistoricalInfoResponse_messageType struct{}

func (x fastReflection_QueryHistoricalInfoResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryHistoricalInfoResponse)(nil)
}
func (x fastReflection_QueryHistoricalInfoResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryHistoricalInfoResponse)
}
func (x fastReflection_QueryHistoricalInfoResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryHistoricalInfoResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryHistoricalInfoResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryHistoricalInfoResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryHistoricalInfoResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryHistoricalInfoResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryHistoricalInfoResponse) New() protoreflect.Message {
	return new(fastReflection_QueryHistoricalInfoResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryHistoricalInfoResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryHistoricalInfoResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryHistoricalInfoResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Hist != nil {
		value := protoreflect.ValueOfMessage(x.Hist.ProtoReflect())
		if !f(fd_QueryHistoricalInfoResponse_hist, value) {
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
func (x *fastReflection_QueryHistoricalInfoResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoResponse.hist":
		return x.Hist != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryHistoricalInfoResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoResponse.hist":
		x.Hist = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryHistoricalInfoResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoResponse.hist":
		value := x.Hist
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryHistoricalInfoResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoResponse.hist":
		x.Hist = value.Message().Interface().(*HistoricalInfo)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryHistoricalInfoResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoResponse.hist":
		if x.Hist == nil {
			x.Hist = new(HistoricalInfo)
		}
		return protoreflect.ValueOfMessage(x.Hist.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryHistoricalInfoResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryHistoricalInfoResponse.hist":
		m := new(HistoricalInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryHistoricalInfoResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryHistoricalInfoResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryHistoricalInfoResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryHistoricalInfoResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryHistoricalInfoResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryHistoricalInfoResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryHistoricalInfoResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryHistoricalInfoResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryHistoricalInfoResponse)
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
		if x.Hist != nil {
			l = options.Size(x.Hist)
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
		x := input.Message.Interface().(*QueryHistoricalInfoResponse)
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
		if x.Hist != nil {
			encoded, err := options.Marshal(x.Hist)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*QueryHistoricalInfoResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryHistoricalInfoResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryHistoricalInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hist", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Hist == nil {
					x.Hist = &HistoricalInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Hist); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryPoolRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryPoolRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryPoolRequest")
}

var _ protoreflect.Message = (*fastReflection_QueryPoolRequest)(nil)

type fastReflection_QueryPoolRequest QueryPoolRequest

func (x *QueryPoolRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryPoolRequest)(x)
}

func (x *QueryPoolRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryPoolRequest_messageType fastReflection_QueryPoolRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryPoolRequest_messageType{}

type fastReflection_QueryPoolRequest_messageType struct{}

func (x fastReflection_QueryPoolRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryPoolRequest)(nil)
}
func (x fastReflection_QueryPoolRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryPoolRequest)
}
func (x fastReflection_QueryPoolRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryPoolRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryPoolRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryPoolRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryPoolRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryPoolRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryPoolRequest) New() protoreflect.Message {
	return new(fastReflection_QueryPoolRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryPoolRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryPoolRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryPoolRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_QueryPoolRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryPoolRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryPoolRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryPoolRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryPoolRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryPoolRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryPoolRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryPoolRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryPoolRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryPoolRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryPoolRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryPoolRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryPoolRequest)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryPoolRequest)
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
		x := input.Message.Interface().(*QueryPoolRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryPoolRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryPoolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_QueryPoolResponse      protoreflect.MessageDescriptor
	fd_QueryPoolResponse_pool protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryPoolResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryPoolResponse")
	fd_QueryPoolResponse_pool = md_QueryPoolResponse.Fields().ByName("pool")
}

var _ protoreflect.Message = (*fastReflection_QueryPoolResponse)(nil)

type fastReflection_QueryPoolResponse QueryPoolResponse

func (x *QueryPoolResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryPoolResponse)(x)
}

func (x *QueryPoolResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryPoolResponse_messageType fastReflection_QueryPoolResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryPoolResponse_messageType{}

type fastReflection_QueryPoolResponse_messageType struct{}

func (x fastReflection_QueryPoolResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryPoolResponse)(nil)
}
func (x fastReflection_QueryPoolResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryPoolResponse)
}
func (x fastReflection_QueryPoolResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryPoolResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryPoolResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryPoolResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryPoolResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryPoolResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryPoolResponse) New() protoreflect.Message {
	return new(fastReflection_QueryPoolResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryPoolResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryPoolResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryPoolResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Pool != nil {
		value := protoreflect.ValueOfMessage(x.Pool.ProtoReflect())
		if !f(fd_QueryPoolResponse_pool, value) {
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
func (x *fastReflection_QueryPoolResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryPoolResponse.pool":
		return x.Pool != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryPoolResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryPoolResponse.pool":
		x.Pool = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryPoolResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryPoolResponse.pool":
		value := x.Pool
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryPoolResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryPoolResponse.pool":
		x.Pool = value.Message().Interface().(*Pool)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryPoolResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryPoolResponse.pool":
		if x.Pool == nil {
			x.Pool = new(Pool)
		}
		return protoreflect.ValueOfMessage(x.Pool.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryPoolResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryPoolResponse.pool":
		m := new(Pool)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryPoolResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryPoolResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryPoolResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryPoolResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryPoolResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryPoolResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryPoolResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryPoolResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryPoolResponse)
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
		if x.Pool != nil {
			l = options.Size(x.Pool)
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
		x := input.Message.Interface().(*QueryPoolResponse)
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
		if x.Pool != nil {
			encoded, err := options.Marshal(x.Pool)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*QueryPoolResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryPoolResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Pool == nil {
					x.Pool = &Pool{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pool); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
	md_QueryParamsRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryParamsRequest = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryParamsRequest")
}

var _ protoreflect.Message = (*fastReflection_QueryParamsRequest)(nil)

type fastReflection_QueryParamsRequest QueryParamsRequest

func (x *QueryParamsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryParamsRequest)(x)
}

func (x *QueryParamsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryParamsRequest_messageType fastReflection_QueryParamsRequest_messageType
var _ protoreflect.MessageType = fastReflection_QueryParamsRequest_messageType{}

type fastReflection_QueryParamsRequest_messageType struct{}

func (x fastReflection_QueryParamsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryParamsRequest)(nil)
}
func (x fastReflection_QueryParamsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryParamsRequest)
}
func (x fastReflection_QueryParamsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryParamsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryParamsRequest) Type() protoreflect.MessageType {
	return _fastReflection_QueryParamsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryParamsRequest) New() protoreflect.Message {
	return new(fastReflection_QueryParamsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryParamsRequest) Interface() protoreflect.ProtoMessage {
	return (*QueryParamsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryParamsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_QueryParamsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryParamsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryParamsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryParamsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryParamsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsRequest"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryParamsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryParamsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryParamsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryParamsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryParamsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryParamsRequest)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*QueryParamsRequest)
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
		x := input.Message.Interface().(*QueryParamsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_QueryParamsResponse        protoreflect.MessageDescriptor
	fd_QueryParamsResponse_params protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_query_proto_init()
	md_QueryParamsResponse = File_cosmos_staking_v1beta1_query_proto.Messages().ByName("QueryParamsResponse")
	fd_QueryParamsResponse_params = md_QueryParamsResponse.Fields().ByName("params")
}

var _ protoreflect.Message = (*fastReflection_QueryParamsResponse)(nil)

type fastReflection_QueryParamsResponse QueryParamsResponse

func (x *QueryParamsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_QueryParamsResponse)(x)
}

func (x *QueryParamsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[27]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_QueryParamsResponse_messageType fastReflection_QueryParamsResponse_messageType
var _ protoreflect.MessageType = fastReflection_QueryParamsResponse_messageType{}

type fastReflection_QueryParamsResponse_messageType struct{}

func (x fastReflection_QueryParamsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_QueryParamsResponse)(nil)
}
func (x fastReflection_QueryParamsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_QueryParamsResponse)
}
func (x fastReflection_QueryParamsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_QueryParamsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_QueryParamsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_QueryParamsResponse) Type() protoreflect.MessageType {
	return _fastReflection_QueryParamsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_QueryParamsResponse) New() protoreflect.Message {
	return new(fastReflection_QueryParamsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_QueryParamsResponse) Interface() protoreflect.ProtoMessage {
	return (*QueryParamsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_QueryParamsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_QueryParamsResponse_params, value) {
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
func (x *fastReflection_QueryParamsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryParamsResponse.params":
		return x.Params != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryParamsResponse.params":
		x.Params = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_QueryParamsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.QueryParamsResponse.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_QueryParamsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryParamsResponse.params":
		x.Params = value.Message().Interface().(*Params)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_QueryParamsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryParamsResponse.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_QueryParamsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.QueryParamsResponse.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.QueryParamsResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.QueryParamsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_QueryParamsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.QueryParamsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_QueryParamsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_QueryParamsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_QueryParamsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_QueryParamsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*QueryParamsResponse)
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
		if x.Params != nil {
			l = options.Size(x.Params)
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
		x := input.Message.Interface().(*QueryParamsResponse)
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
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*QueryParamsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
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
// source: cosmos/staking/v1beta1/query.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryValidatorsRequest is request type for Query/Validators RPC method.
type QueryValidatorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// status enables to query for validators matching a given status.
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryValidatorsRequest) Reset() {
	*x = QueryValidatorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorsRequest) ProtoMessage() {}

// Deprecated: Use QueryValidatorsRequest.ProtoReflect.Descriptor instead.
func (*QueryValidatorsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryValidatorsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *QueryValidatorsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryValidatorsResponse is response type for the Query/Validators RPC method
type QueryValidatorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validators contains all the queried validators.
	Validators []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryValidatorsResponse) Reset() {
	*x = QueryValidatorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorsResponse) ProtoMessage() {}

// Deprecated: Use QueryValidatorsResponse.ProtoReflect.Descriptor instead.
func (*QueryValidatorsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryValidatorsResponse) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *QueryValidatorsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryValidatorRequest is response type for the Query/Validator RPC method
type QueryValidatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
}

func (x *QueryValidatorRequest) Reset() {
	*x = QueryValidatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorRequest) ProtoMessage() {}

// Deprecated: Use QueryValidatorRequest.ProtoReflect.Descriptor instead.
func (*QueryValidatorRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryValidatorRequest) GetValidatorAddr() string {
	if x != nil {
		return x.ValidatorAddr
	}
	return ""
}

// QueryValidatorResponse is response type for the Query/Validator RPC method
type QueryValidatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator defines the the validator info.
	Validator *Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
}

func (x *QueryValidatorResponse) Reset() {
	*x = QueryValidatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorResponse) ProtoMessage() {}

// Deprecated: Use QueryValidatorResponse.ProtoReflect.Descriptor instead.
func (*QueryValidatorResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryValidatorResponse) GetValidator() *Validator {
	if x != nil {
		return x.Validator
	}
	return nil
}

// QueryValidatorDelegationsRequest is request type for the
// Query/ValidatorDelegations RPC method
type QueryValidatorDelegationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryValidatorDelegationsRequest) Reset() {
	*x = QueryValidatorDelegationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorDelegationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorDelegationsRequest) ProtoMessage() {}

// Deprecated: Use QueryValidatorDelegationsRequest.ProtoReflect.Descriptor instead.
func (*QueryValidatorDelegationsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{4}
}

func (x *QueryValidatorDelegationsRequest) GetValidatorAddr() string {
	if x != nil {
		return x.ValidatorAddr
	}
	return ""
}

func (x *QueryValidatorDelegationsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryValidatorDelegationsResponse is response type for the
// Query/ValidatorDelegations RPC method
type QueryValidatorDelegationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegationResponses []*DelegationResponse `protobuf:"bytes,1,rep,name=delegation_responses,json=delegationResponses,proto3" json:"delegation_responses,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryValidatorDelegationsResponse) Reset() {
	*x = QueryValidatorDelegationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorDelegationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorDelegationsResponse) ProtoMessage() {}

// Deprecated: Use QueryValidatorDelegationsResponse.ProtoReflect.Descriptor instead.
func (*QueryValidatorDelegationsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{5}
}

func (x *QueryValidatorDelegationsResponse) GetDelegationResponses() []*DelegationResponse {
	if x != nil {
		return x.DelegationResponses
	}
	return nil
}

func (x *QueryValidatorDelegationsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryValidatorUnbondingDelegationsRequest is required type for the
// Query/ValidatorUnbondingDelegations RPC method
type QueryValidatorUnbondingDelegationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryValidatorUnbondingDelegationsRequest) Reset() {
	*x = QueryValidatorUnbondingDelegationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorUnbondingDelegationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorUnbondingDelegationsRequest) ProtoMessage() {}

// Deprecated: Use QueryValidatorUnbondingDelegationsRequest.ProtoReflect.Descriptor instead.
func (*QueryValidatorUnbondingDelegationsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{6}
}

func (x *QueryValidatorUnbondingDelegationsRequest) GetValidatorAddr() string {
	if x != nil {
		return x.ValidatorAddr
	}
	return ""
}

func (x *QueryValidatorUnbondingDelegationsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryValidatorUnbondingDelegationsResponse is response type for the
// Query/ValidatorUnbondingDelegations RPC method.
type QueryValidatorUnbondingDelegationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnbondingResponses []*UnbondingDelegation `protobuf:"bytes,1,rep,name=unbonding_responses,json=unbondingResponses,proto3" json:"unbonding_responses,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryValidatorUnbondingDelegationsResponse) Reset() {
	*x = QueryValidatorUnbondingDelegationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidatorUnbondingDelegationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidatorUnbondingDelegationsResponse) ProtoMessage() {}

// Deprecated: Use QueryValidatorUnbondingDelegationsResponse.ProtoReflect.Descriptor instead.
func (*QueryValidatorUnbondingDelegationsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{7}
}

func (x *QueryValidatorUnbondingDelegationsResponse) GetUnbondingResponses() []*UnbondingDelegation {
	if x != nil {
		return x.UnbondingResponses
	}
	return nil
}

func (x *QueryValidatorUnbondingDelegationsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryDelegationRequest is request type for the Query/Delegation RPC method.
type QueryDelegationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_addr defines the delegator address to query for.
	DelegatorAddr string `protobuf:"bytes,1,opt,name=delegator_addr,json=delegatorAddr,proto3" json:"delegator_addr,omitempty"`
	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,2,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
}

func (x *QueryDelegationRequest) Reset() {
	*x = QueryDelegationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegationRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegationRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegationRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{8}
}

func (x *QueryDelegationRequest) GetDelegatorAddr() string {
	if x != nil {
		return x.DelegatorAddr
	}
	return ""
}

func (x *QueryDelegationRequest) GetValidatorAddr() string {
	if x != nil {
		return x.ValidatorAddr
	}
	return ""
}

// QueryDelegationResponse is response type for the Query/Delegation RPC method.
type QueryDelegationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegation_responses defines the delegation info of a delegation.
	DelegationResponse *DelegationResponse `protobuf:"bytes,1,opt,name=delegation_response,json=delegationResponse,proto3" json:"delegation_response,omitempty"`
}

func (x *QueryDelegationResponse) Reset() {
	*x = QueryDelegationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegationResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegationResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegationResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{9}
}

func (x *QueryDelegationResponse) GetDelegationResponse() *DelegationResponse {
	if x != nil {
		return x.DelegationResponse
	}
	return nil
}

// QueryUnbondingDelegationRequest is request type for the
// Query/UnbondingDelegation RPC method.
type QueryUnbondingDelegationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_addr defines the delegator address to query for.
	DelegatorAddr string `protobuf:"bytes,1,opt,name=delegator_addr,json=delegatorAddr,proto3" json:"delegator_addr,omitempty"`
	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,2,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
}

func (x *QueryUnbondingDelegationRequest) Reset() {
	*x = QueryUnbondingDelegationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUnbondingDelegationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUnbondingDelegationRequest) ProtoMessage() {}

// Deprecated: Use QueryUnbondingDelegationRequest.ProtoReflect.Descriptor instead.
func (*QueryUnbondingDelegationRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{10}
}

func (x *QueryUnbondingDelegationRequest) GetDelegatorAddr() string {
	if x != nil {
		return x.DelegatorAddr
	}
	return ""
}

func (x *QueryUnbondingDelegationRequest) GetValidatorAddr() string {
	if x != nil {
		return x.ValidatorAddr
	}
	return ""
}

// QueryDelegationResponse is response type for the Query/UnbondingDelegation
// RPC method.
type QueryUnbondingDelegationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unbond defines the unbonding information of a delegation.
	Unbond *UnbondingDelegation `protobuf:"bytes,1,opt,name=unbond,proto3" json:"unbond,omitempty"`
}

func (x *QueryUnbondingDelegationResponse) Reset() {
	*x = QueryUnbondingDelegationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUnbondingDelegationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUnbondingDelegationResponse) ProtoMessage() {}

// Deprecated: Use QueryUnbondingDelegationResponse.ProtoReflect.Descriptor instead.
func (*QueryUnbondingDelegationResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{11}
}

func (x *QueryUnbondingDelegationResponse) GetUnbond() *UnbondingDelegation {
	if x != nil {
		return x.Unbond
	}
	return nil
}

// QueryDelegatorDelegationsRequest is request type for the
// Query/DelegatorDelegations RPC method.
type QueryDelegatorDelegationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_addr defines the delegator address to query for.
	DelegatorAddr string `protobuf:"bytes,1,opt,name=delegator_addr,json=delegatorAddr,proto3" json:"delegator_addr,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryDelegatorDelegationsRequest) Reset() {
	*x = QueryDelegatorDelegationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorDelegationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorDelegationsRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegatorDelegationsRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegatorDelegationsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{12}
}

func (x *QueryDelegatorDelegationsRequest) GetDelegatorAddr() string {
	if x != nil {
		return x.DelegatorAddr
	}
	return ""
}

func (x *QueryDelegatorDelegationsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryDelegatorDelegationsResponse is response type for the
// Query/DelegatorDelegations RPC method.
type QueryDelegatorDelegationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegation_responses defines all the delegations' info of a delegator.
	DelegationResponses []*DelegationResponse `protobuf:"bytes,1,rep,name=delegation_responses,json=delegationResponses,proto3" json:"delegation_responses,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryDelegatorDelegationsResponse) Reset() {
	*x = QueryDelegatorDelegationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorDelegationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorDelegationsResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegatorDelegationsResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegatorDelegationsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{13}
}

func (x *QueryDelegatorDelegationsResponse) GetDelegationResponses() []*DelegationResponse {
	if x != nil {
		return x.DelegationResponses
	}
	return nil
}

func (x *QueryDelegatorDelegationsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryDelegatorUnbondingDelegationsRequest is request type for the
// Query/DelegatorUnbondingDelegations RPC method.
type QueryDelegatorUnbondingDelegationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_addr defines the delegator address to query for.
	DelegatorAddr string `protobuf:"bytes,1,opt,name=delegator_addr,json=delegatorAddr,proto3" json:"delegator_addr,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryDelegatorUnbondingDelegationsRequest) Reset() {
	*x = QueryDelegatorUnbondingDelegationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorUnbondingDelegationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorUnbondingDelegationsRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegatorUnbondingDelegationsRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegatorUnbondingDelegationsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{14}
}

func (x *QueryDelegatorUnbondingDelegationsRequest) GetDelegatorAddr() string {
	if x != nil {
		return x.DelegatorAddr
	}
	return ""
}

func (x *QueryDelegatorUnbondingDelegationsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryUnbondingDelegatorDelegationsResponse is response type for the
// Query/UnbondingDelegatorDelegations RPC method.
type QueryDelegatorUnbondingDelegationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnbondingResponses []*UnbondingDelegation `protobuf:"bytes,1,rep,name=unbonding_responses,json=unbondingResponses,proto3" json:"unbonding_responses,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryDelegatorUnbondingDelegationsResponse) Reset() {
	*x = QueryDelegatorUnbondingDelegationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorUnbondingDelegationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorUnbondingDelegationsResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegatorUnbondingDelegationsResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegatorUnbondingDelegationsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{15}
}

func (x *QueryDelegatorUnbondingDelegationsResponse) GetUnbondingResponses() []*UnbondingDelegation {
	if x != nil {
		return x.UnbondingResponses
	}
	return nil
}

func (x *QueryDelegatorUnbondingDelegationsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryRedelegationsRequest is request type for the Query/Redelegations RPC
// method.
type QueryRedelegationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_addr defines the delegator address to query for.
	DelegatorAddr string `protobuf:"bytes,1,opt,name=delegator_addr,json=delegatorAddr,proto3" json:"delegator_addr,omitempty"`
	// src_validator_addr defines the validator address to redelegate from.
	SrcValidatorAddr string `protobuf:"bytes,2,opt,name=src_validator_addr,json=srcValidatorAddr,proto3" json:"src_validator_addr,omitempty"`
	// dst_validator_addr defines the validator address to redelegate to.
	DstValidatorAddr string `protobuf:"bytes,3,opt,name=dst_validator_addr,json=dstValidatorAddr,proto3" json:"dst_validator_addr,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryRedelegationsRequest) Reset() {
	*x = QueryRedelegationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRedelegationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRedelegationsRequest) ProtoMessage() {}

// Deprecated: Use QueryRedelegationsRequest.ProtoReflect.Descriptor instead.
func (*QueryRedelegationsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{16}
}

func (x *QueryRedelegationsRequest) GetDelegatorAddr() string {
	if x != nil {
		return x.DelegatorAddr
	}
	return ""
}

func (x *QueryRedelegationsRequest) GetSrcValidatorAddr() string {
	if x != nil {
		return x.SrcValidatorAddr
	}
	return ""
}

func (x *QueryRedelegationsRequest) GetDstValidatorAddr() string {
	if x != nil {
		return x.DstValidatorAddr
	}
	return ""
}

func (x *QueryRedelegationsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryRedelegationsResponse is response type for the Query/Redelegations RPC
// method.
type QueryRedelegationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedelegationResponses []*RedelegationResponse `protobuf:"bytes,1,rep,name=redelegation_responses,json=redelegationResponses,proto3" json:"redelegation_responses,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryRedelegationsResponse) Reset() {
	*x = QueryRedelegationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRedelegationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRedelegationsResponse) ProtoMessage() {}

// Deprecated: Use QueryRedelegationsResponse.ProtoReflect.Descriptor instead.
func (*QueryRedelegationsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{17}
}

func (x *QueryRedelegationsResponse) GetRedelegationResponses() []*RedelegationResponse {
	if x != nil {
		return x.RedelegationResponses
	}
	return nil
}

func (x *QueryRedelegationsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryDelegatorValidatorsRequest is request type for the
// Query/DelegatorValidators RPC method.
type QueryDelegatorValidatorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_addr defines the delegator address to query for.
	DelegatorAddr string `protobuf:"bytes,1,opt,name=delegator_addr,json=delegatorAddr,proto3" json:"delegator_addr,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *v1beta1.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryDelegatorValidatorsRequest) Reset() {
	*x = QueryDelegatorValidatorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorValidatorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorValidatorsRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegatorValidatorsRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegatorValidatorsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{18}
}

func (x *QueryDelegatorValidatorsRequest) GetDelegatorAddr() string {
	if x != nil {
		return x.DelegatorAddr
	}
	return ""
}

func (x *QueryDelegatorValidatorsRequest) GetPagination() *v1beta1.PageRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryDelegatorValidatorsResponse is response type for the
// Query/DelegatorValidators RPC method.
type QueryDelegatorValidatorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validators defines the the validators' info of a delegator.
	Validators []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *v1beta1.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *QueryDelegatorValidatorsResponse) Reset() {
	*x = QueryDelegatorValidatorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorValidatorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorValidatorsResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegatorValidatorsResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegatorValidatorsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{19}
}

func (x *QueryDelegatorValidatorsResponse) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *QueryDelegatorValidatorsResponse) GetPagination() *v1beta1.PageResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// QueryDelegatorValidatorRequest is request type for the
// Query/DelegatorValidator RPC method.
type QueryDelegatorValidatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_addr defines the delegator address to query for.
	DelegatorAddr string `protobuf:"bytes,1,opt,name=delegator_addr,json=delegatorAddr,proto3" json:"delegator_addr,omitempty"`
	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,2,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
}

func (x *QueryDelegatorValidatorRequest) Reset() {
	*x = QueryDelegatorValidatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorValidatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorValidatorRequest) ProtoMessage() {}

// Deprecated: Use QueryDelegatorValidatorRequest.ProtoReflect.Descriptor instead.
func (*QueryDelegatorValidatorRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{20}
}

func (x *QueryDelegatorValidatorRequest) GetDelegatorAddr() string {
	if x != nil {
		return x.DelegatorAddr
	}
	return ""
}

func (x *QueryDelegatorValidatorRequest) GetValidatorAddr() string {
	if x != nil {
		return x.ValidatorAddr
	}
	return ""
}

// QueryDelegatorValidatorResponse response type for the
// Query/DelegatorValidator RPC method.
type QueryDelegatorValidatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// validator defines the the validator info.
	Validator *Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
}

func (x *QueryDelegatorValidatorResponse) Reset() {
	*x = QueryDelegatorValidatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDelegatorValidatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDelegatorValidatorResponse) ProtoMessage() {}

// Deprecated: Use QueryDelegatorValidatorResponse.ProtoReflect.Descriptor instead.
func (*QueryDelegatorValidatorResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{21}
}

func (x *QueryDelegatorValidatorResponse) GetValidator() *Validator {
	if x != nil {
		return x.Validator
	}
	return nil
}

// QueryHistoricalInfoRequest is request type for the Query/HistoricalInfo RPC
// method.
type QueryHistoricalInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// height defines at which height to query the historical info.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *QueryHistoricalInfoRequest) Reset() {
	*x = QueryHistoricalInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryHistoricalInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryHistoricalInfoRequest) ProtoMessage() {}

// Deprecated: Use QueryHistoricalInfoRequest.ProtoReflect.Descriptor instead.
func (*QueryHistoricalInfoRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{22}
}

func (x *QueryHistoricalInfoRequest) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

// QueryHistoricalInfoResponse is response type for the Query/HistoricalInfo RPC
// method.
type QueryHistoricalInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hist defines the historical info at the given height.
	Hist *HistoricalInfo `protobuf:"bytes,1,opt,name=hist,proto3" json:"hist,omitempty"`
}

func (x *QueryHistoricalInfoResponse) Reset() {
	*x = QueryHistoricalInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryHistoricalInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryHistoricalInfoResponse) ProtoMessage() {}

// Deprecated: Use QueryHistoricalInfoResponse.ProtoReflect.Descriptor instead.
func (*QueryHistoricalInfoResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{23}
}

func (x *QueryHistoricalInfoResponse) GetHist() *HistoricalInfo {
	if x != nil {
		return x.Hist
	}
	return nil
}

// QueryPoolRequest is request type for the Query/Pool RPC method.
type QueryPoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryPoolRequest) Reset() {
	*x = QueryPoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPoolRequest) ProtoMessage() {}

// Deprecated: Use QueryPoolRequest.ProtoReflect.Descriptor instead.
func (*QueryPoolRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{24}
}

// QueryPoolResponse is response type for the Query/Pool RPC method.
type QueryPoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pool defines the pool info.
	Pool *Pool `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (x *QueryPoolResponse) Reset() {
	*x = QueryPoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPoolResponse) ProtoMessage() {}

// Deprecated: Use QueryPoolResponse.ProtoReflect.Descriptor instead.
func (*QueryPoolResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{25}
}

func (x *QueryPoolResponse) GetPool() *Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryParamsRequest) Reset() {
	*x = QueryParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsRequest) ProtoMessage() {}

// Deprecated: Use QueryParamsRequest.ProtoReflect.Descriptor instead.
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{26}
}

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params holds all the parameters of this module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *QueryParamsResponse) Reset() {
	*x = QueryParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_query_proto_msgTypes[27]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsResponse) ProtoMessage() {}

// Deprecated: Use QueryParamsResponse.ProtoReflect.Descriptor instead.
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_query_proto_rawDescGZIP(), []int{27}
}

func (x *QueryParamsResponse) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_cosmos_staking_v1beta1_query_proto protoreflect.FileDescriptor

var file_cosmos_staking_v1beta1_query_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x2a, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a,
	0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xab, 0x01, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f,
	0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x22,
	0x5f, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0xab, 0x01, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe8,
	0x01, 0x0a, 0x21, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7a, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b,
	0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x13, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb4, 0x01, 0x0a, 0x29, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x62, 0x6f,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xd9, 0x01, 0x0a, 0x2a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x62, 0x0a, 0x13, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52,
	0x12, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa4, 0x01, 0x0a,
	0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x3f, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8,
	0xa0, 0x1f, 0x00, 0x22, 0x76, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b,
	0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x1f,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3f, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x3f, 0x0a, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x6d, 0x0a, 0x20, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x06, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x06, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x20, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3f, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0,
	0x1f, 0x00, 0x22, 0xd1, 0x01, 0x0a, 0x21, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x47, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x29, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0x88,
	0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xd9, 0x01, 0x0a, 0x2a, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x62,
	0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x12, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x02, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x46, 0x0a, 0x12, 0x73, 0x72, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18,
	0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x73, 0x72, 0x63, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x46, 0x0a, 0x12, 0x64, 0x73,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x10, 0x64, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00,
	0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xd0, 0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x16, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x15, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x47,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb4, 0x01, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x46, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xb4,
	0x01, 0x0a, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xac, 0x01, 0x0a, 0x1e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x3f, 0x0a, 0x0e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00,
	0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x68, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x34,
	0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x59, 0x0a, 0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x68, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x68, 0x69, 0x73, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x6f, 0x6f, 0x6c, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c,
	0x22, 0x14, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0xea, 0x15, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x99, 0x01, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0xa7, 0x01, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x2d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x12, 0xd4, 0x01, 0x0a, 0x14,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x41, 0x12, 0x3f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0xf9, 0x01, 0x0a, 0x1d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x62, 0x6f,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x55,
	0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x4b, 0x12, 0x49, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x2f, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xc7,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x52, 0x12, 0x50, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x12, 0xf7, 0x01, 0x0a, 0x13, 0x55, 0x6e, 0x62,
	0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55,
	0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x67, 0x12, 0x65, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x2f, 0x75, 0x6e,
	0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0xc9, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x12, 0xf9,
	0x01, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x62, 0x6f,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x41, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x55, 0x6e, 0x62, 0x6f, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x12,
	0x49, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x7d, 0x2f, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xc1, 0x01, 0x0a, 0x0d, 0x52,
	0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x49, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x12, 0x41, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f,
	0x7b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d,
	0x2f, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xd0,
	0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x40, 0x12, 0x3e, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0xde, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x51, 0x12, 0x4f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x7d, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x7d, 0x12, 0xb3, 0x01, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f,
	0x7b, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x7d, 0x12, 0x81, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x6f,
	0x6c, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x89, 0x01, 0x0a,
	0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0xea, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x53, 0x58, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x16, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x43, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_staking_v1beta1_query_proto_rawDescOnce sync.Once
	file_cosmos_staking_v1beta1_query_proto_rawDescData = file_cosmos_staking_v1beta1_query_proto_rawDesc
)

func file_cosmos_staking_v1beta1_query_proto_rawDescGZIP() []byte {
	file_cosmos_staking_v1beta1_query_proto_rawDescOnce.Do(func() {
		file_cosmos_staking_v1beta1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_staking_v1beta1_query_proto_rawDescData)
	})
	return file_cosmos_staking_v1beta1_query_proto_rawDescData
}

var file_cosmos_staking_v1beta1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 28)
var file_cosmos_staking_v1beta1_query_proto_goTypes = []interface{}{
	(*QueryValidatorsRequest)(nil),                     // 0: cosmos.staking.v1beta1.QueryValidatorsRequest
	(*QueryValidatorsResponse)(nil),                    // 1: cosmos.staking.v1beta1.QueryValidatorsResponse
	(*QueryValidatorRequest)(nil),                      // 2: cosmos.staking.v1beta1.QueryValidatorRequest
	(*QueryValidatorResponse)(nil),                     // 3: cosmos.staking.v1beta1.QueryValidatorResponse
	(*QueryValidatorDelegationsRequest)(nil),           // 4: cosmos.staking.v1beta1.QueryValidatorDelegationsRequest
	(*QueryValidatorDelegationsResponse)(nil),          // 5: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse
	(*QueryValidatorUnbondingDelegationsRequest)(nil),  // 6: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest
	(*QueryValidatorUnbondingDelegationsResponse)(nil), // 7: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse
	(*QueryDelegationRequest)(nil),                     // 8: cosmos.staking.v1beta1.QueryDelegationRequest
	(*QueryDelegationResponse)(nil),                    // 9: cosmos.staking.v1beta1.QueryDelegationResponse
	(*QueryUnbondingDelegationRequest)(nil),            // 10: cosmos.staking.v1beta1.QueryUnbondingDelegationRequest
	(*QueryUnbondingDelegationResponse)(nil),           // 11: cosmos.staking.v1beta1.QueryUnbondingDelegationResponse
	(*QueryDelegatorDelegationsRequest)(nil),           // 12: cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest
	(*QueryDelegatorDelegationsResponse)(nil),          // 13: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse
	(*QueryDelegatorUnbondingDelegationsRequest)(nil),  // 14: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest
	(*QueryDelegatorUnbondingDelegationsResponse)(nil), // 15: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse
	(*QueryRedelegationsRequest)(nil),                  // 16: cosmos.staking.v1beta1.QueryRedelegationsRequest
	(*QueryRedelegationsResponse)(nil),                 // 17: cosmos.staking.v1beta1.QueryRedelegationsResponse
	(*QueryDelegatorValidatorsRequest)(nil),            // 18: cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest
	(*QueryDelegatorValidatorsResponse)(nil),           // 19: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse
	(*QueryDelegatorValidatorRequest)(nil),             // 20: cosmos.staking.v1beta1.QueryDelegatorValidatorRequest
	(*QueryDelegatorValidatorResponse)(nil),            // 21: cosmos.staking.v1beta1.QueryDelegatorValidatorResponse
	(*QueryHistoricalInfoRequest)(nil),                 // 22: cosmos.staking.v1beta1.QueryHistoricalInfoRequest
	(*QueryHistoricalInfoResponse)(nil),                // 23: cosmos.staking.v1beta1.QueryHistoricalInfoResponse
	(*QueryPoolRequest)(nil),                           // 24: cosmos.staking.v1beta1.QueryPoolRequest
	(*QueryPoolResponse)(nil),                          // 25: cosmos.staking.v1beta1.QueryPoolResponse
	(*QueryParamsRequest)(nil),                         // 26: cosmos.staking.v1beta1.QueryParamsRequest
	(*QueryParamsResponse)(nil),                        // 27: cosmos.staking.v1beta1.QueryParamsResponse
	(*v1beta1.PageRequest)(nil),                        // 28: cosmos.base.query.v1beta1.PageRequest
	(*Validator)(nil),                                  // 29: cosmos.staking.v1beta1.Validator
	(*v1beta1.PageResponse)(nil),                       // 30: cosmos.base.query.v1beta1.PageResponse
	(*DelegationResponse)(nil),                         // 31: cosmos.staking.v1beta1.DelegationResponse
	(*UnbondingDelegation)(nil),                        // 32: cosmos.staking.v1beta1.UnbondingDelegation
	(*RedelegationResponse)(nil),                       // 33: cosmos.staking.v1beta1.RedelegationResponse
	(*HistoricalInfo)(nil),                             // 34: cosmos.staking.v1beta1.HistoricalInfo
	(*Pool)(nil),                                       // 35: cosmos.staking.v1beta1.Pool
	(*Params)(nil),                                     // 36: cosmos.staking.v1beta1.Params
}
var file_cosmos_staking_v1beta1_query_proto_depIdxs = []int32{
	28, // 0: cosmos.staking.v1beta1.QueryValidatorsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	29, // 1: cosmos.staking.v1beta1.QueryValidatorsResponse.validators:type_name -> cosmos.staking.v1beta1.Validator
	30, // 2: cosmos.staking.v1beta1.QueryValidatorsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	29, // 3: cosmos.staking.v1beta1.QueryValidatorResponse.validator:type_name -> cosmos.staking.v1beta1.Validator
	28, // 4: cosmos.staking.v1beta1.QueryValidatorDelegationsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	31, // 5: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.delegation_responses:type_name -> cosmos.staking.v1beta1.DelegationResponse
	30, // 6: cosmos.staking.v1beta1.QueryValidatorDelegationsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	28, // 7: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	32, // 8: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.unbonding_responses:type_name -> cosmos.staking.v1beta1.UnbondingDelegation
	30, // 9: cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	31, // 10: cosmos.staking.v1beta1.QueryDelegationResponse.delegation_response:type_name -> cosmos.staking.v1beta1.DelegationResponse
	32, // 11: cosmos.staking.v1beta1.QueryUnbondingDelegationResponse.unbond:type_name -> cosmos.staking.v1beta1.UnbondingDelegation
	28, // 12: cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	31, // 13: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.delegation_responses:type_name -> cosmos.staking.v1beta1.DelegationResponse
	30, // 14: cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	28, // 15: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	32, // 16: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.unbonding_responses:type_name -> cosmos.staking.v1beta1.UnbondingDelegation
	30, // 17: cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	28, // 18: cosmos.staking.v1beta1.QueryRedelegationsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	33, // 19: cosmos.staking.v1beta1.QueryRedelegationsResponse.redelegation_responses:type_name -> cosmos.staking.v1beta1.RedelegationResponse
	30, // 20: cosmos.staking.v1beta1.QueryRedelegationsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	28, // 21: cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest.pagination:type_name -> cosmos.base.query.v1beta1.PageRequest
	29, // 22: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.validators:type_name -> cosmos.staking.v1beta1.Validator
	30, // 23: cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse.pagination:type_name -> cosmos.base.query.v1beta1.PageResponse
	29, // 24: cosmos.staking.v1beta1.QueryDelegatorValidatorResponse.validator:type_name -> cosmos.staking.v1beta1.Validator
	34, // 25: cosmos.staking.v1beta1.QueryHistoricalInfoResponse.hist:type_name -> cosmos.staking.v1beta1.HistoricalInfo
	35, // 26: cosmos.staking.v1beta1.QueryPoolResponse.pool:type_name -> cosmos.staking.v1beta1.Pool
	36, // 27: cosmos.staking.v1beta1.QueryParamsResponse.params:type_name -> cosmos.staking.v1beta1.Params
	0,  // 28: cosmos.staking.v1beta1.Query.Validators:input_type -> cosmos.staking.v1beta1.QueryValidatorsRequest
	2,  // 29: cosmos.staking.v1beta1.Query.Validator:input_type -> cosmos.staking.v1beta1.QueryValidatorRequest
	4,  // 30: cosmos.staking.v1beta1.Query.ValidatorDelegations:input_type -> cosmos.staking.v1beta1.QueryValidatorDelegationsRequest
	6,  // 31: cosmos.staking.v1beta1.Query.ValidatorUnbondingDelegations:input_type -> cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest
	8,  // 32: cosmos.staking.v1beta1.Query.Delegation:input_type -> cosmos.staking.v1beta1.QueryDelegationRequest
	10, // 33: cosmos.staking.v1beta1.Query.UnbondingDelegation:input_type -> cosmos.staking.v1beta1.QueryUnbondingDelegationRequest
	12, // 34: cosmos.staking.v1beta1.Query.DelegatorDelegations:input_type -> cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest
	14, // 35: cosmos.staking.v1beta1.Query.DelegatorUnbondingDelegations:input_type -> cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest
	16, // 36: cosmos.staking.v1beta1.Query.Redelegations:input_type -> cosmos.staking.v1beta1.QueryRedelegationsRequest
	18, // 37: cosmos.staking.v1beta1.Query.DelegatorValidators:input_type -> cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest
	20, // 38: cosmos.staking.v1beta1.Query.DelegatorValidator:input_type -> cosmos.staking.v1beta1.QueryDelegatorValidatorRequest
	22, // 39: cosmos.staking.v1beta1.Query.HistoricalInfo:input_type -> cosmos.staking.v1beta1.QueryHistoricalInfoRequest
	24, // 40: cosmos.staking.v1beta1.Query.Pool:input_type -> cosmos.staking.v1beta1.QueryPoolRequest
	26, // 41: cosmos.staking.v1beta1.Query.Params:input_type -> cosmos.staking.v1beta1.QueryParamsRequest
	1,  // 42: cosmos.staking.v1beta1.Query.Validators:output_type -> cosmos.staking.v1beta1.QueryValidatorsResponse
	3,  // 43: cosmos.staking.v1beta1.Query.Validator:output_type -> cosmos.staking.v1beta1.QueryValidatorResponse
	5,  // 44: cosmos.staking.v1beta1.Query.ValidatorDelegations:output_type -> cosmos.staking.v1beta1.QueryValidatorDelegationsResponse
	7,  // 45: cosmos.staking.v1beta1.Query.ValidatorUnbondingDelegations:output_type -> cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse
	9,  // 46: cosmos.staking.v1beta1.Query.Delegation:output_type -> cosmos.staking.v1beta1.QueryDelegationResponse
	11, // 47: cosmos.staking.v1beta1.Query.UnbondingDelegation:output_type -> cosmos.staking.v1beta1.QueryUnbondingDelegationResponse
	13, // 48: cosmos.staking.v1beta1.Query.DelegatorDelegations:output_type -> cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse
	15, // 49: cosmos.staking.v1beta1.Query.DelegatorUnbondingDelegations:output_type -> cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse
	17, // 50: cosmos.staking.v1beta1.Query.Redelegations:output_type -> cosmos.staking.v1beta1.QueryRedelegationsResponse
	19, // 51: cosmos.staking.v1beta1.Query.DelegatorValidators:output_type -> cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse
	21, // 52: cosmos.staking.v1beta1.Query.DelegatorValidator:output_type -> cosmos.staking.v1beta1.QueryDelegatorValidatorResponse
	23, // 53: cosmos.staking.v1beta1.Query.HistoricalInfo:output_type -> cosmos.staking.v1beta1.QueryHistoricalInfoResponse
	25, // 54: cosmos.staking.v1beta1.Query.Pool:output_type -> cosmos.staking.v1beta1.QueryPoolResponse
	27, // 55: cosmos.staking.v1beta1.Query.Params:output_type -> cosmos.staking.v1beta1.QueryParamsResponse
	42, // [42:56] is the sub-list for method output_type
	28, // [28:42] is the sub-list for method input_type
	28, // [28:28] is the sub-list for extension type_name
	28, // [28:28] is the sub-list for extension extendee
	0,  // [0:28] is the sub-list for field type_name
}

func init() { file_cosmos_staking_v1beta1_query_proto_init() }
func file_cosmos_staking_v1beta1_query_proto_init() {
	if File_cosmos_staking_v1beta1_query_proto != nil {
		return
	}
	file_cosmos_staking_v1beta1_staking_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_staking_v1beta1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorsRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorsResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorDelegationsRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorDelegationsResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorUnbondingDelegationsRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidatorUnbondingDelegationsResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegationRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegationResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUnbondingDelegationRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUnbondingDelegationResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorDelegationsRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorDelegationsResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorUnbondingDelegationsRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorUnbondingDelegationsResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRedelegationsRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRedelegationsResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorValidatorsRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorValidatorsResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorValidatorRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDelegatorValidatorResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryHistoricalInfoRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryHistoricalInfoResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPoolRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPoolResponse); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsRequest); i {
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
		file_cosmos_staking_v1beta1_query_proto_msgTypes[27].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsResponse); i {
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
			RawDescriptor: file_cosmos_staking_v1beta1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   28,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_staking_v1beta1_query_proto_goTypes,
		DependencyIndexes: file_cosmos_staking_v1beta1_query_proto_depIdxs,
		MessageInfos:      file_cosmos_staking_v1beta1_query_proto_msgTypes,
	}.Build()
	File_cosmos_staking_v1beta1_query_proto = out.File
	file_cosmos_staking_v1beta1_query_proto_rawDesc = nil
	file_cosmos_staking_v1beta1_query_proto_goTypes = nil
	file_cosmos_staking_v1beta1_query_proto_depIdxs = nil
}
