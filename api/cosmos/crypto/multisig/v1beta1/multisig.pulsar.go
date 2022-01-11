package multisigv1beta1

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

var _ protoreflect.List = (*_MultiSignature_1_list)(nil)

type _MultiSignature_1_list struct {
	list *[][]byte
}

func (x *_MultiSignature_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MultiSignature_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfBytes((*x.list)[i])
}

func (x *_MultiSignature_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MultiSignature_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MultiSignature_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MultiSignature at list field Signatures as it is not of Message kind"))
}

func (x *_MultiSignature_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MultiSignature_1_list) NewElement() protoreflect.Value {
	var v []byte
	return protoreflect.ValueOfBytes(v)
}

func (x *_MultiSignature_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MultiSignature            protoreflect.MessageDescriptor
	fd_MultiSignature_signatures protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_multisig_v1beta1_multisig_proto_init()
	md_MultiSignature = File_cosmos_crypto_multisig_v1beta1_multisig_proto.Messages().ByName("MultiSignature")
	fd_MultiSignature_signatures = md_MultiSignature.Fields().ByName("signatures")
}

var _ protoreflect.Message = (*fastReflection_MultiSignature)(nil)

type fastReflection_MultiSignature MultiSignature

func (x *MultiSignature) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MultiSignature)(x)
}

func (x *MultiSignature) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_multisig_v1beta1_multisig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MultiSignature_messageType fastReflection_MultiSignature_messageType
var _ protoreflect.MessageType = fastReflection_MultiSignature_messageType{}

type fastReflection_MultiSignature_messageType struct{}

func (x fastReflection_MultiSignature_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MultiSignature)(nil)
}
func (x fastReflection_MultiSignature_messageType) New() protoreflect.Message {
	return new(fastReflection_MultiSignature)
}
func (x fastReflection_MultiSignature_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiSignature
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MultiSignature) Descriptor() protoreflect.MessageDescriptor {
	return md_MultiSignature
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MultiSignature) Type() protoreflect.MessageType {
	return _fastReflection_MultiSignature_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MultiSignature) New() protoreflect.Message {
	return new(fastReflection_MultiSignature)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MultiSignature) Interface() protoreflect.ProtoMessage {
	return (*MultiSignature)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MultiSignature) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Signatures) != 0 {
		value := protoreflect.ValueOfList(&_MultiSignature_1_list{list: &x.Signatures})
		if !f(fd_MultiSignature_signatures, value) {
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
func (x *fastReflection_MultiSignature) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.MultiSignature.signatures":
		return len(x.Signatures) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.MultiSignature"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.MultiSignature does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiSignature) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.MultiSignature.signatures":
		x.Signatures = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.MultiSignature"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.MultiSignature does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MultiSignature) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.multisig.v1beta1.MultiSignature.signatures":
		if len(x.Signatures) == 0 {
			return protoreflect.ValueOfList(&_MultiSignature_1_list{})
		}
		listValue := &_MultiSignature_1_list{list: &x.Signatures}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.MultiSignature"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.MultiSignature does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MultiSignature) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.MultiSignature.signatures":
		lv := value.List()
		clv := lv.(*_MultiSignature_1_list)
		x.Signatures = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.MultiSignature"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.MultiSignature does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MultiSignature) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.MultiSignature.signatures":
		if x.Signatures == nil {
			x.Signatures = [][]byte{}
		}
		value := &_MultiSignature_1_list{list: &x.Signatures}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.MultiSignature"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.MultiSignature does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MultiSignature) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.MultiSignature.signatures":
		list := [][]byte{}
		return protoreflect.ValueOfList(&_MultiSignature_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.MultiSignature"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.MultiSignature does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MultiSignature) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.multisig.v1beta1.MultiSignature", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MultiSignature) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MultiSignature) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MultiSignature) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MultiSignature) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MultiSignature)
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
		if len(x.Signatures) > 0 {
			for _, b := range x.Signatures {
				l = len(b)
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
		x := input.Message.Interface().(*MultiSignature)
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
		if len(x.Signatures) > 0 {
			for iNdEx := len(x.Signatures) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Signatures[iNdEx])
				copy(dAtA[i:], x.Signatures[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signatures[iNdEx])))
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
		x := input.Message.Interface().(*MultiSignature)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiSignature: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MultiSignature: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
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
				x.Signatures = append(x.Signatures, make([]byte, postIndex-iNdEx))
				copy(x.Signatures[len(x.Signatures)-1], dAtA[iNdEx:postIndex])
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
	md_CompactBitArray                   protoreflect.MessageDescriptor
	fd_CompactBitArray_extra_bits_stored protoreflect.FieldDescriptor
	fd_CompactBitArray_elems             protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_multisig_v1beta1_multisig_proto_init()
	md_CompactBitArray = File_cosmos_crypto_multisig_v1beta1_multisig_proto.Messages().ByName("CompactBitArray")
	fd_CompactBitArray_extra_bits_stored = md_CompactBitArray.Fields().ByName("extra_bits_stored")
	fd_CompactBitArray_elems = md_CompactBitArray.Fields().ByName("elems")
}

var _ protoreflect.Message = (*fastReflection_CompactBitArray)(nil)

type fastReflection_CompactBitArray CompactBitArray

func (x *CompactBitArray) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CompactBitArray)(x)
}

func (x *CompactBitArray) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_multisig_v1beta1_multisig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CompactBitArray_messageType fastReflection_CompactBitArray_messageType
var _ protoreflect.MessageType = fastReflection_CompactBitArray_messageType{}

type fastReflection_CompactBitArray_messageType struct{}

func (x fastReflection_CompactBitArray_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CompactBitArray)(nil)
}
func (x fastReflection_CompactBitArray_messageType) New() protoreflect.Message {
	return new(fastReflection_CompactBitArray)
}
func (x fastReflection_CompactBitArray_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CompactBitArray
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CompactBitArray) Descriptor() protoreflect.MessageDescriptor {
	return md_CompactBitArray
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CompactBitArray) Type() protoreflect.MessageType {
	return _fastReflection_CompactBitArray_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CompactBitArray) New() protoreflect.Message {
	return new(fastReflection_CompactBitArray)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CompactBitArray) Interface() protoreflect.ProtoMessage {
	return (*CompactBitArray)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CompactBitArray) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ExtraBitsStored != uint32(0) {
		value := protoreflect.ValueOfUint32(x.ExtraBitsStored)
		if !f(fd_CompactBitArray_extra_bits_stored, value) {
			return
		}
	}
	if len(x.Elems) != 0 {
		value := protoreflect.ValueOfBytes(x.Elems)
		if !f(fd_CompactBitArray_elems, value) {
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
func (x *fastReflection_CompactBitArray) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.extra_bits_stored":
		return x.ExtraBitsStored != uint32(0)
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.elems":
		return len(x.Elems) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.CompactBitArray"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.CompactBitArray does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CompactBitArray) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.extra_bits_stored":
		x.ExtraBitsStored = uint32(0)
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.elems":
		x.Elems = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.CompactBitArray"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.CompactBitArray does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CompactBitArray) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.extra_bits_stored":
		value := x.ExtraBitsStored
		return protoreflect.ValueOfUint32(value)
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.elems":
		value := x.Elems
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.CompactBitArray"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.CompactBitArray does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_CompactBitArray) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.extra_bits_stored":
		x.ExtraBitsStored = uint32(value.Uint())
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.elems":
		x.Elems = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.CompactBitArray"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.CompactBitArray does not contain field %s", fd.FullName()))
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
func (x *fastReflection_CompactBitArray) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.extra_bits_stored":
		panic(fmt.Errorf("field extra_bits_stored of message cosmos.crypto.multisig.v1beta1.CompactBitArray is not mutable"))
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.elems":
		panic(fmt.Errorf("field elems of message cosmos.crypto.multisig.v1beta1.CompactBitArray is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.CompactBitArray"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.CompactBitArray does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CompactBitArray) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.extra_bits_stored":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.crypto.multisig.v1beta1.CompactBitArray.elems":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.multisig.v1beta1.CompactBitArray"))
		}
		panic(fmt.Errorf("message cosmos.crypto.multisig.v1beta1.CompactBitArray does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CompactBitArray) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.multisig.v1beta1.CompactBitArray", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CompactBitArray) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CompactBitArray) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_CompactBitArray) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CompactBitArray) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CompactBitArray)
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
		if x.ExtraBitsStored != 0 {
			n += 1 + runtime.Sov(uint64(x.ExtraBitsStored))
		}
		l = len(x.Elems)
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
		x := input.Message.Interface().(*CompactBitArray)
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
		if len(x.Elems) > 0 {
			i -= len(x.Elems)
			copy(dAtA[i:], x.Elems)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Elems)))
			i--
			dAtA[i] = 0x12
		}
		if x.ExtraBitsStored != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ExtraBitsStored))
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
		x := input.Message.Interface().(*CompactBitArray)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CompactBitArray: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CompactBitArray: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ExtraBitsStored", wireType)
				}
				x.ExtraBitsStored = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ExtraBitsStored |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Elems", wireType)
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
				x.Elems = append(x.Elems[:0], dAtA[iNdEx:postIndex]...)
				if x.Elems == nil {
					x.Elems = []byte{}
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
// source: cosmos/crypto/multisig/v1beta1/multisig.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MultiSignature wraps the signatures from a multisig.LegacyAminoPubKey.
// See cosmos.tx.v1betata1.ModeInfo.Multi for how to specify which signers
// signed and with which modes.
type MultiSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signatures [][]byte `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *MultiSignature) Reset() {
	*x = MultiSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_multisig_v1beta1_multisig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiSignature) ProtoMessage() {}

// Deprecated: Use MultiSignature.ProtoReflect.Descriptor instead.
func (*MultiSignature) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescGZIP(), []int{0}
}

func (x *MultiSignature) GetSignatures() [][]byte {
	if x != nil {
		return x.Signatures
	}
	return nil
}

// CompactBitArray is an implementation of a space efficient bit array.
// This is used to ensure that the encoded data takes up a minimal amount of
// space after proto encoding.
// This is not thread safe, and is not intended for concurrent usage.
type CompactBitArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtraBitsStored uint32 `protobuf:"varint,1,opt,name=extra_bits_stored,json=extraBitsStored,proto3" json:"extra_bits_stored,omitempty"`
	Elems           []byte `protobuf:"bytes,2,opt,name=elems,proto3" json:"elems,omitempty"`
}

func (x *CompactBitArray) Reset() {
	*x = CompactBitArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_multisig_v1beta1_multisig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactBitArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactBitArray) ProtoMessage() {}

// Deprecated: Use CompactBitArray.ProtoReflect.Descriptor instead.
func (*CompactBitArray) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescGZIP(), []int{1}
}

func (x *CompactBitArray) GetExtraBitsStored() uint32 {
	if x != nil {
		return x.ExtraBitsStored
	}
	return 0
}

func (x *CompactBitArray) GetElems() []byte {
	if x != nil {
		return x.Elems
	}
	return nil
}

var File_cosmos_crypto_multisig_v1beta1_multisig_proto protoreflect.FileDescriptor

var file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x3a, 0x04, 0xd0, 0xa1, 0x1f, 0x01, 0x22, 0x59, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x42, 0x69, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x42, 0x69, 0x74, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6c, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x6c, 0x65,
	0x6d, 0x73, 0x3a, 0x04, 0x98, 0xa0, 0x1f, 0x00, 0x42, 0x9f, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42,
	0x0d, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x4d, 0xaa, 0x02, 0x1e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67,
	0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1e, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69,
	0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2a, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73,
	0x69, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a,
	0x3a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x3a, 0x3a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69,
	0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescOnce sync.Once
	file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescData = file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDesc
)

func file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescGZIP() []byte {
	file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescOnce.Do(func() {
		file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescData)
	})
	return file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDescData
}

var file_cosmos_crypto_multisig_v1beta1_multisig_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_crypto_multisig_v1beta1_multisig_proto_goTypes = []interface{}{
	(*MultiSignature)(nil),  // 0: cosmos.crypto.multisig.v1beta1.MultiSignature
	(*CompactBitArray)(nil), // 1: cosmos.crypto.multisig.v1beta1.CompactBitArray
}
var file_cosmos_crypto_multisig_v1beta1_multisig_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cosmos_crypto_multisig_v1beta1_multisig_proto_init() }
func file_cosmos_crypto_multisig_v1beta1_multisig_proto_init() {
	if File_cosmos_crypto_multisig_v1beta1_multisig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_crypto_multisig_v1beta1_multisig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiSignature); i {
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
		file_cosmos_crypto_multisig_v1beta1_multisig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactBitArray); i {
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
			RawDescriptor: file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_crypto_multisig_v1beta1_multisig_proto_goTypes,
		DependencyIndexes: file_cosmos_crypto_multisig_v1beta1_multisig_proto_depIdxs,
		MessageInfos:      file_cosmos_crypto_multisig_v1beta1_multisig_proto_msgTypes,
	}.Build()
	File_cosmos_crypto_multisig_v1beta1_multisig_proto = out.File
	file_cosmos_crypto_multisig_v1beta1_multisig_proto_rawDesc = nil
	file_cosmos_crypto_multisig_v1beta1_multisig_proto_goTypes = nil
	file_cosmos_crypto_multisig_v1beta1_multisig_proto_depIdxs = nil
}
