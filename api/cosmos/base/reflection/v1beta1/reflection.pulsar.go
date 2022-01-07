package reflectionv1beta1

import (
	context "context"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ListAllInterfacesRequest protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_base_reflection_v1beta1_reflection_proto_init()
	md_ListAllInterfacesRequest = File_cosmos_base_reflection_v1beta1_reflection_proto.Messages().ByName("ListAllInterfacesRequest")
}

var _ protoreflect.Message = (*fastReflection_ListAllInterfacesRequest)(nil)

type fastReflection_ListAllInterfacesRequest ListAllInterfacesRequest

func (x *ListAllInterfacesRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ListAllInterfacesRequest)(x)
}

func (x *ListAllInterfacesRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ListAllInterfacesRequest_messageType fastReflection_ListAllInterfacesRequest_messageType
var _ protoreflect.MessageType = fastReflection_ListAllInterfacesRequest_messageType{}

type fastReflection_ListAllInterfacesRequest_messageType struct{}

func (x fastReflection_ListAllInterfacesRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ListAllInterfacesRequest)(nil)
}
func (x fastReflection_ListAllInterfacesRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_ListAllInterfacesRequest)
}
func (x fastReflection_ListAllInterfacesRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ListAllInterfacesRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ListAllInterfacesRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_ListAllInterfacesRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ListAllInterfacesRequest) Type() protoreflect.MessageType {
	return _fastReflection_ListAllInterfacesRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ListAllInterfacesRequest) New() protoreflect.Message {
	return new(fastReflection_ListAllInterfacesRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ListAllInterfacesRequest) Interface() protoreflect.ProtoMessage {
	return (*ListAllInterfacesRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ListAllInterfacesRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_ListAllInterfacesRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ListAllInterfacesRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ListAllInterfacesRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ListAllInterfacesRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ListAllInterfacesRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ListAllInterfacesRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ListAllInterfacesRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v1beta1.ListAllInterfacesRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ListAllInterfacesRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ListAllInterfacesRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ListAllInterfacesRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ListAllInterfacesRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ListAllInterfacesRequest)
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
		x := input.Message.Interface().(*ListAllInterfacesRequest)
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
		x := input.Message.Interface().(*ListAllInterfacesRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ListAllInterfacesRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ListAllInterfacesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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

var _ protoreflect.List = (*_ListAllInterfacesResponse_1_list)(nil)

type _ListAllInterfacesResponse_1_list struct {
	list *[]string
}

func (x *_ListAllInterfacesResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ListAllInterfacesResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_ListAllInterfacesResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_ListAllInterfacesResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_ListAllInterfacesResponse_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message ListAllInterfacesResponse at list field InterfaceNames as it is not of Message kind"))
}

func (x *_ListAllInterfacesResponse_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_ListAllInterfacesResponse_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_ListAllInterfacesResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ListAllInterfacesResponse                 protoreflect.MessageDescriptor
	fd_ListAllInterfacesResponse_interface_names protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v1beta1_reflection_proto_init()
	md_ListAllInterfacesResponse = File_cosmos_base_reflection_v1beta1_reflection_proto.Messages().ByName("ListAllInterfacesResponse")
	fd_ListAllInterfacesResponse_interface_names = md_ListAllInterfacesResponse.Fields().ByName("interface_names")
}

var _ protoreflect.Message = (*fastReflection_ListAllInterfacesResponse)(nil)

type fastReflection_ListAllInterfacesResponse ListAllInterfacesResponse

func (x *ListAllInterfacesResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ListAllInterfacesResponse)(x)
}

func (x *ListAllInterfacesResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ListAllInterfacesResponse_messageType fastReflection_ListAllInterfacesResponse_messageType
var _ protoreflect.MessageType = fastReflection_ListAllInterfacesResponse_messageType{}

type fastReflection_ListAllInterfacesResponse_messageType struct{}

func (x fastReflection_ListAllInterfacesResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ListAllInterfacesResponse)(nil)
}
func (x fastReflection_ListAllInterfacesResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_ListAllInterfacesResponse)
}
func (x fastReflection_ListAllInterfacesResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ListAllInterfacesResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ListAllInterfacesResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_ListAllInterfacesResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ListAllInterfacesResponse) Type() protoreflect.MessageType {
	return _fastReflection_ListAllInterfacesResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ListAllInterfacesResponse) New() protoreflect.Message {
	return new(fastReflection_ListAllInterfacesResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ListAllInterfacesResponse) Interface() protoreflect.ProtoMessage {
	return (*ListAllInterfacesResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ListAllInterfacesResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.InterfaceNames) != 0 {
		value := protoreflect.ValueOfList(&_ListAllInterfacesResponse_1_list{list: &x.InterfaceNames})
		if !f(fd_ListAllInterfacesResponse_interface_names, value) {
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
func (x *fastReflection_ListAllInterfacesResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListAllInterfacesResponse.interface_names":
		return len(x.InterfaceNames) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ListAllInterfacesResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListAllInterfacesResponse.interface_names":
		x.InterfaceNames = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ListAllInterfacesResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v1beta1.ListAllInterfacesResponse.interface_names":
		if len(x.InterfaceNames) == 0 {
			return protoreflect.ValueOfList(&_ListAllInterfacesResponse_1_list{})
		}
		listValue := &_ListAllInterfacesResponse_1_list{list: &x.InterfaceNames}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ListAllInterfacesResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListAllInterfacesResponse.interface_names":
		lv := value.List()
		clv := lv.(*_ListAllInterfacesResponse_1_list)
		x.InterfaceNames = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ListAllInterfacesResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListAllInterfacesResponse.interface_names":
		if x.InterfaceNames == nil {
			x.InterfaceNames = []string{}
		}
		value := &_ListAllInterfacesResponse_1_list{list: &x.InterfaceNames}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ListAllInterfacesResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListAllInterfacesResponse.interface_names":
		list := []string{}
		return protoreflect.ValueOfList(&_ListAllInterfacesResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListAllInterfacesResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListAllInterfacesResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ListAllInterfacesResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v1beta1.ListAllInterfacesResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ListAllInterfacesResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ListAllInterfacesResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ListAllInterfacesResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ListAllInterfacesResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ListAllInterfacesResponse)
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
		if len(x.InterfaceNames) > 0 {
			for _, s := range x.InterfaceNames {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*ListAllInterfacesResponse)
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
		if len(x.InterfaceNames) > 0 {
			for iNdEx := len(x.InterfaceNames) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.InterfaceNames[iNdEx])
				copy(dAtA[i:], x.InterfaceNames[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.InterfaceNames[iNdEx])))
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
		x := input.Message.Interface().(*ListAllInterfacesResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ListAllInterfacesResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ListAllInterfacesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InterfaceNames", wireType)
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
				x.InterfaceNames = append(x.InterfaceNames, string(dAtA[iNdEx:postIndex]))
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
	md_ListImplementationsRequest                protoreflect.MessageDescriptor
	fd_ListImplementationsRequest_interface_name protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v1beta1_reflection_proto_init()
	md_ListImplementationsRequest = File_cosmos_base_reflection_v1beta1_reflection_proto.Messages().ByName("ListImplementationsRequest")
	fd_ListImplementationsRequest_interface_name = md_ListImplementationsRequest.Fields().ByName("interface_name")
}

var _ protoreflect.Message = (*fastReflection_ListImplementationsRequest)(nil)

type fastReflection_ListImplementationsRequest ListImplementationsRequest

func (x *ListImplementationsRequest) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ListImplementationsRequest)(x)
}

func (x *ListImplementationsRequest) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ListImplementationsRequest_messageType fastReflection_ListImplementationsRequest_messageType
var _ protoreflect.MessageType = fastReflection_ListImplementationsRequest_messageType{}

type fastReflection_ListImplementationsRequest_messageType struct{}

func (x fastReflection_ListImplementationsRequest_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ListImplementationsRequest)(nil)
}
func (x fastReflection_ListImplementationsRequest_messageType) New() protoreflect.Message {
	return new(fastReflection_ListImplementationsRequest)
}
func (x fastReflection_ListImplementationsRequest_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ListImplementationsRequest
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ListImplementationsRequest) Descriptor() protoreflect.MessageDescriptor {
	return md_ListImplementationsRequest
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ListImplementationsRequest) Type() protoreflect.MessageType {
	return _fastReflection_ListImplementationsRequest_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ListImplementationsRequest) New() protoreflect.Message {
	return new(fastReflection_ListImplementationsRequest)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ListImplementationsRequest) Interface() protoreflect.ProtoMessage {
	return (*ListImplementationsRequest)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ListImplementationsRequest) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.InterfaceName != "" {
		value := protoreflect.ValueOfString(x.InterfaceName)
		if !f(fd_ListImplementationsRequest_interface_name, value) {
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
func (x *fastReflection_ListImplementationsRequest) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsRequest.interface_name":
		return x.InterfaceName != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsRequest does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ListImplementationsRequest) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsRequest.interface_name":
		x.InterfaceName = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsRequest does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ListImplementationsRequest) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsRequest.interface_name":
		value := x.InterfaceName
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsRequest does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ListImplementationsRequest) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsRequest.interface_name":
		x.InterfaceName = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsRequest does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ListImplementationsRequest) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsRequest.interface_name":
		panic(fmt.Errorf("field interface_name of message cosmos.base.reflection.v1beta1.ListImplementationsRequest is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsRequest does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ListImplementationsRequest) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsRequest.interface_name":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsRequest"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsRequest does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ListImplementationsRequest) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v1beta1.ListImplementationsRequest", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ListImplementationsRequest) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ListImplementationsRequest) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ListImplementationsRequest) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ListImplementationsRequest) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ListImplementationsRequest)
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
		l = len(x.InterfaceName)
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
		x := input.Message.Interface().(*ListImplementationsRequest)
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
		if len(x.InterfaceName) > 0 {
			i -= len(x.InterfaceName)
			copy(dAtA[i:], x.InterfaceName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.InterfaceName)))
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
		x := input.Message.Interface().(*ListImplementationsRequest)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ListImplementationsRequest: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ListImplementationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InterfaceName", wireType)
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
				x.InterfaceName = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_ListImplementationsResponse_1_list)(nil)

type _ListImplementationsResponse_1_list struct {
	list *[]string
}

func (x *_ListImplementationsResponse_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ListImplementationsResponse_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_ListImplementationsResponse_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_ListImplementationsResponse_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_ListImplementationsResponse_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message ListImplementationsResponse at list field ImplementationMessageNames as it is not of Message kind"))
}

func (x *_ListImplementationsResponse_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_ListImplementationsResponse_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_ListImplementationsResponse_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ListImplementationsResponse                              protoreflect.MessageDescriptor
	fd_ListImplementationsResponse_implementation_message_names protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_reflection_v1beta1_reflection_proto_init()
	md_ListImplementationsResponse = File_cosmos_base_reflection_v1beta1_reflection_proto.Messages().ByName("ListImplementationsResponse")
	fd_ListImplementationsResponse_implementation_message_names = md_ListImplementationsResponse.Fields().ByName("implementation_message_names")
}

var _ protoreflect.Message = (*fastReflection_ListImplementationsResponse)(nil)

type fastReflection_ListImplementationsResponse ListImplementationsResponse

func (x *ListImplementationsResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ListImplementationsResponse)(x)
}

func (x *ListImplementationsResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ListImplementationsResponse_messageType fastReflection_ListImplementationsResponse_messageType
var _ protoreflect.MessageType = fastReflection_ListImplementationsResponse_messageType{}

type fastReflection_ListImplementationsResponse_messageType struct{}

func (x fastReflection_ListImplementationsResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ListImplementationsResponse)(nil)
}
func (x fastReflection_ListImplementationsResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_ListImplementationsResponse)
}
func (x fastReflection_ListImplementationsResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ListImplementationsResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ListImplementationsResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_ListImplementationsResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ListImplementationsResponse) Type() protoreflect.MessageType {
	return _fastReflection_ListImplementationsResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ListImplementationsResponse) New() protoreflect.Message {
	return new(fastReflection_ListImplementationsResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ListImplementationsResponse) Interface() protoreflect.ProtoMessage {
	return (*ListImplementationsResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ListImplementationsResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.ImplementationMessageNames) != 0 {
		value := protoreflect.ValueOfList(&_ListImplementationsResponse_1_list{list: &x.ImplementationMessageNames})
		if !f(fd_ListImplementationsResponse_implementation_message_names, value) {
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
func (x *fastReflection_ListImplementationsResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsResponse.implementation_message_names":
		return len(x.ImplementationMessageNames) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ListImplementationsResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsResponse.implementation_message_names":
		x.ImplementationMessageNames = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ListImplementationsResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsResponse.implementation_message_names":
		if len(x.ImplementationMessageNames) == 0 {
			return protoreflect.ValueOfList(&_ListImplementationsResponse_1_list{})
		}
		listValue := &_ListImplementationsResponse_1_list{list: &x.ImplementationMessageNames}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ListImplementationsResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsResponse.implementation_message_names":
		lv := value.List()
		clv := lv.(*_ListImplementationsResponse_1_list)
		x.ImplementationMessageNames = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ListImplementationsResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsResponse.implementation_message_names":
		if x.ImplementationMessageNames == nil {
			x.ImplementationMessageNames = []string{}
		}
		value := &_ListImplementationsResponse_1_list{list: &x.ImplementationMessageNames}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ListImplementationsResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.reflection.v1beta1.ListImplementationsResponse.implementation_message_names":
		list := []string{}
		return protoreflect.ValueOfList(&_ListImplementationsResponse_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.reflection.v1beta1.ListImplementationsResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.reflection.v1beta1.ListImplementationsResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ListImplementationsResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.reflection.v1beta1.ListImplementationsResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ListImplementationsResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ListImplementationsResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ListImplementationsResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ListImplementationsResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ListImplementationsResponse)
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
		if len(x.ImplementationMessageNames) > 0 {
			for _, s := range x.ImplementationMessageNames {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*ListImplementationsResponse)
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
		if len(x.ImplementationMessageNames) > 0 {
			for iNdEx := len(x.ImplementationMessageNames) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.ImplementationMessageNames[iNdEx])
				copy(dAtA[i:], x.ImplementationMessageNames[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ImplementationMessageNames[iNdEx])))
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
		x := input.Message.Interface().(*ListImplementationsResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ListImplementationsResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ListImplementationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ImplementationMessageNames", wireType)
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
				x.ImplementationMessageNames = append(x.ImplementationMessageNames, string(dAtA[iNdEx:postIndex]))
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReflectionServiceClient is the client API for ReflectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReflectionServiceClient interface {
	// ListAllInterfaces lists all the interfaces registered in the interface
	// registry.
	ListAllInterfaces(ctx context.Context, in *ListAllInterfacesRequest, opts ...grpc.CallOption) (*ListAllInterfacesResponse, error)
	// ListImplementations list all the concrete types that implement a given
	// interface.
	ListImplementations(ctx context.Context, in *ListImplementationsRequest, opts ...grpc.CallOption) (*ListImplementationsResponse, error)
}

type reflectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReflectionServiceClient(cc grpc.ClientConnInterface) ReflectionServiceClient {
	return &reflectionServiceClient{cc}
}

func (c *reflectionServiceClient) ListAllInterfaces(ctx context.Context, in *ListAllInterfacesRequest, opts ...grpc.CallOption) (*ListAllInterfacesResponse, error) {
	out := new(ListAllInterfacesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v1beta1.ReflectionService/ListAllInterfaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reflectionServiceClient) ListImplementations(ctx context.Context, in *ListImplementationsRequest, opts ...grpc.CallOption) (*ListImplementationsResponse, error) {
	out := new(ListImplementationsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.reflection.v1beta1.ReflectionService/ListImplementations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReflectionServiceServer is the server API for ReflectionService service.
// All implementations must embed UnimplementedReflectionServiceServer
// for forward compatibility
type ReflectionServiceServer interface {
	// ListAllInterfaces lists all the interfaces registered in the interface
	// registry.
	ListAllInterfaces(context.Context, *ListAllInterfacesRequest) (*ListAllInterfacesResponse, error)
	// ListImplementations list all the concrete types that implement a given
	// interface.
	ListImplementations(context.Context, *ListImplementationsRequest) (*ListImplementationsResponse, error)
	mustEmbedUnimplementedReflectionServiceServer()
}

// UnimplementedReflectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReflectionServiceServer struct {
}

func (UnimplementedReflectionServiceServer) ListAllInterfaces(context.Context, *ListAllInterfacesRequest) (*ListAllInterfacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllInterfaces not implemented")
}
func (UnimplementedReflectionServiceServer) ListImplementations(context.Context, *ListImplementationsRequest) (*ListImplementationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImplementations not implemented")
}
func (UnimplementedReflectionServiceServer) mustEmbedUnimplementedReflectionServiceServer() {}

// UnsafeReflectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReflectionServiceServer will
// result in compilation errors.
type UnsafeReflectionServiceServer interface {
	mustEmbedUnimplementedReflectionServiceServer()
}

func RegisterReflectionServiceServer(s grpc.ServiceRegistrar, srv ReflectionServiceServer) {
	s.RegisterService(&ReflectionService_ServiceDesc, srv)
}

func _ReflectionService_ListAllInterfaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllInterfacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).ListAllInterfaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v1beta1.ReflectionService/ListAllInterfaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).ListAllInterfaces(ctx, req.(*ListAllInterfacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReflectionService_ListImplementations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImplementationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReflectionServiceServer).ListImplementations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.reflection.v1beta1.ReflectionService/ListImplementations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReflectionServiceServer).ListImplementations(ctx, req.(*ListImplementationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReflectionService_ServiceDesc is the grpc.ServiceDesc for ReflectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReflectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.reflection.v1beta1.ReflectionService",
	HandlerType: (*ReflectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllInterfaces",
			Handler:    _ReflectionService_ListAllInterfaces_Handler,
		},
		{
			MethodName: "ListImplementations",
			Handler:    _ReflectionService_ListImplementations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/reflection/v1beta1/reflection.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/base/reflection/v1beta1/reflection.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListAllInterfacesRequest is the request type of the ListAllInterfaces RPC.
type ListAllInterfacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAllInterfacesRequest) Reset() {
	*x = ListAllInterfacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllInterfacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllInterfacesRequest) ProtoMessage() {}

// Deprecated: Use ListAllInterfacesRequest.ProtoReflect.Descriptor instead.
func (*ListAllInterfacesRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescGZIP(), []int{0}
}

// ListAllInterfacesResponse is the response type of the ListAllInterfaces RPC.
type ListAllInterfacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// interface_names is an array of all the registered interfaces.
	InterfaceNames []string `protobuf:"bytes,1,rep,name=interface_names,json=interfaceNames,proto3" json:"interface_names,omitempty"`
}

func (x *ListAllInterfacesResponse) Reset() {
	*x = ListAllInterfacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllInterfacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllInterfacesResponse) ProtoMessage() {}

// Deprecated: Use ListAllInterfacesResponse.ProtoReflect.Descriptor instead.
func (*ListAllInterfacesResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescGZIP(), []int{1}
}

func (x *ListAllInterfacesResponse) GetInterfaceNames() []string {
	if x != nil {
		return x.InterfaceNames
	}
	return nil
}

// ListImplementationsRequest is the request type of the ListImplementations
// RPC.
type ListImplementationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// interface_name defines the interface to query the implementations for.
	InterfaceName string `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
}

func (x *ListImplementationsRequest) Reset() {
	*x = ListImplementationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImplementationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImplementationsRequest) ProtoMessage() {}

// Deprecated: Use ListImplementationsRequest.ProtoReflect.Descriptor instead.
func (*ListImplementationsRequest) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescGZIP(), []int{2}
}

func (x *ListImplementationsRequest) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

// ListImplementationsResponse is the response type of the ListImplementations
// RPC.
type ListImplementationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImplementationMessageNames []string `protobuf:"bytes,1,rep,name=implementation_message_names,json=implementationMessageNames,proto3" json:"implementation_message_names,omitempty"`
}

func (x *ListImplementationsResponse) Reset() {
	*x = ListImplementationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImplementationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImplementationsResponse) ProtoMessage() {}

// Deprecated: Use ListImplementationsResponse.ProtoReflect.Descriptor instead.
func (*ListImplementationsResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescGZIP(), []int{3}
}

func (x *ListImplementationsResponse) GetImplementationMessageNames() []string {
	if x != nil {
		return x.ImplementationMessageNames
	}
	return nil
}

var File_cosmos_base_reflection_v1beta1_reflection_proto protoreflect.FileDescriptor

var file_cosmos_base_reflection_v1beta1_reflection_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x19, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0x43, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x1c, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1a, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0xb8, 0x03, 0x0a, 0x11, 0x52, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbc, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x12, 0x2a, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0xe3, 0x01, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72,
	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x4d, 0x12, 0x4b, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0xa3, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f, 0x52, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x42, 0x52, 0xaa, 0x02, 0x1e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c,
	0x42, 0x61, 0x73, 0x65, 0x5c, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2a, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42,
	0x61, 0x73, 0x65, 0x3a, 0x3a, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescOnce sync.Once
	file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescData = file_cosmos_base_reflection_v1beta1_reflection_proto_rawDesc
)

func file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescGZIP() []byte {
	file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescOnce.Do(func() {
		file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescData)
	})
	return file_cosmos_base_reflection_v1beta1_reflection_proto_rawDescData
}

var file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_base_reflection_v1beta1_reflection_proto_goTypes = []interface{}{
	(*ListAllInterfacesRequest)(nil),    // 0: cosmos.base.reflection.v1beta1.ListAllInterfacesRequest
	(*ListAllInterfacesResponse)(nil),   // 1: cosmos.base.reflection.v1beta1.ListAllInterfacesResponse
	(*ListImplementationsRequest)(nil),  // 2: cosmos.base.reflection.v1beta1.ListImplementationsRequest
	(*ListImplementationsResponse)(nil), // 3: cosmos.base.reflection.v1beta1.ListImplementationsResponse
}
var file_cosmos_base_reflection_v1beta1_reflection_proto_depIdxs = []int32{
	0, // 0: cosmos.base.reflection.v1beta1.ReflectionService.ListAllInterfaces:input_type -> cosmos.base.reflection.v1beta1.ListAllInterfacesRequest
	2, // 1: cosmos.base.reflection.v1beta1.ReflectionService.ListImplementations:input_type -> cosmos.base.reflection.v1beta1.ListImplementationsRequest
	1, // 2: cosmos.base.reflection.v1beta1.ReflectionService.ListAllInterfaces:output_type -> cosmos.base.reflection.v1beta1.ListAllInterfacesResponse
	3, // 3: cosmos.base.reflection.v1beta1.ReflectionService.ListImplementations:output_type -> cosmos.base.reflection.v1beta1.ListImplementationsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_base_reflection_v1beta1_reflection_proto_init() }
func file_cosmos_base_reflection_v1beta1_reflection_proto_init() {
	if File_cosmos_base_reflection_v1beta1_reflection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllInterfacesRequest); i {
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
		file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllInterfacesResponse); i {
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
		file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImplementationsRequest); i {
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
		file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImplementationsResponse); i {
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
			RawDescriptor: file_cosmos_base_reflection_v1beta1_reflection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_base_reflection_v1beta1_reflection_proto_goTypes,
		DependencyIndexes: file_cosmos_base_reflection_v1beta1_reflection_proto_depIdxs,
		MessageInfos:      file_cosmos_base_reflection_v1beta1_reflection_proto_msgTypes,
	}.Build()
	File_cosmos_base_reflection_v1beta1_reflection_proto = out.File
	file_cosmos_base_reflection_v1beta1_reflection_proto_rawDesc = nil
	file_cosmos_base_reflection_v1beta1_reflection_proto_goTypes = nil
	file_cosmos_base_reflection_v1beta1_reflection_proto_depIdxs = nil
}
