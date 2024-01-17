protoc --go_out=./pb --go_opt=paths=source_relative ^
 --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ^
 --model-converter_out=./pb --model-converter_opt=paths=source_relative ^
  2pc.proto