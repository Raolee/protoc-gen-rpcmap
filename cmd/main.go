package main

import (
	"google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
	"protoc_gen_rpcmap"
)

func main() {
	var flags protogen.Options
	flags.Run(func(gen *protogen.Plugin) error {
		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}
			err := protoc_gen_rpcmap.GenerateFile(gen, file)
			if err != nil {
				return err
			}
		}

		// 이걸 달아 줘야... proto version 3를 지원한다.
		// google.golang.org/protobuf 코드를 더 까보고 알아냄
		gen.SupportedFeatures = internal_gengo.SupportedFeatures
		return nil
	})
}
