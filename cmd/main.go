package main

import (
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
		return nil
	})
}
