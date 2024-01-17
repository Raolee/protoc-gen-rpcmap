package model_converter

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

var E_BsonName = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50001,
	Name:          "myoptions.bson_name",
	Tag:           "bytes,50001,opt,name=bson_name",
	Filename:      "path/to/your_proto_file.proto",
}
