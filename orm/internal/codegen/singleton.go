package codegen

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/dynamicpb"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
)

type singletonGen struct {
	fileGen
	msg      *protogen.Message
	table    *ormv1alpha1.SingletonDescriptor
	ormTable ormtable.Table
}

func newSingletonGen(fileGen fileGen, msg *protogen.Message, table *ormv1alpha1.SingletonDescriptor) *singletonGen {
	s := &singletonGen{fileGen: fileGen, msg: msg, table: table}
	ot, err := ormtable.Build(ormtable.Options{
		MessageType:         dynamicpb.NewMessageType(msg.Desc),
		SingletonDescriptor: table,
	})
	if err != nil {
		panic(err)
	}
	s.ormTable = ot
	return s
}

func (s singletonGen) gen() {
	s.genInterface()
	s.genStruct()
	s.genInterfaceGuard()
	s.genMethods()
	s.genConstructor()
}

func (s singletonGen) genInterface() {
	s.P("// singleton store")
	s.P("type ", s.messageStoreInterfaceName(s.msg), " interface {")
	s.P("Get(ctx ", contextPkg.Ident("Context"), ") (*", s.msg.GoIdent.GoName, ", error)")
	s.P("Save(ctx ", contextPkg.Ident("Context"), ", ", s.param(s.msg.GoIdent.GoName), "*", s.msg.GoIdent.GoName, ") error")
	s.P("}")
	s.P()
}

func (s singletonGen) genStruct() {
	s.P("type ", s.messageStoreReceiverName(s.msg), " struct {")
	s.P("table ", tablePkg.Ident("Table"))
	s.P("}")
	s.P()
}

func (s singletonGen) genInterfaceGuard() {
	s.P("var _ ", s.messageStoreInterfaceName(s.msg), " = ", s.messageStoreReceiverName(s.msg), "{}")
}

func (s singletonGen) genMethods() {
	receiver := fmt.Sprintf("func (x %s) ", s.messageStoreReceiverName(s.msg))
	varName := s.param(s.msg.GoIdent.GoName)
	// Get
	s.P(receiver, "Get(ctx ", contextPkg.Ident("Context"), ") (*", s.msg.GoIdent.GoName, ", error) {")
	s.P("var ", varName, " ", s.msg.GoIdent.GoName)
	s.P("found, err := x.table.Get(ctx, &", varName, ")")
	s.P("if !found {")
	s.P("return nil, err")
	s.P("}")
	s.P("return &", varName, ", err")
	s.P("}")
	s.P()

	// Save
	s.P(receiver, "Save(ctx ", contextPkg.Ident("Context"), ", ", varName, " *", s.msg.GoIdent.GoName, ") error {")
	s.P("return x.table.Save(ctx, ", varName, ")")
	s.P("}")
	s.P()
}

func (s singletonGen) genConstructor() {
	iface := s.messageStoreInterfaceName(s.msg)
	s.P("func New", iface, "(db ", ormdbPkg.Ident("ModuleDB"), ") (", iface, ", error) {")
	s.P("table := db.GetTable(&", s.msg.GoIdent.GoName, "{})")
	s.P("if table == nil {")
	s.P("return nil, ", ormErrPkg.Ident("TableNotFound.Wrap"), "(string((&", s.msg.GoIdent.GoName, "{}).ProtoReflect().Descriptor().FullName()))")
	s.P("}")
	s.P("return &", s.messageStoreReceiverName(s.msg), "{table}, nil")
	s.P("}")
}
