protoc --go_out=./pb --go_opt=paths=source_relative ^
 --go-grpc_out=./pb --go-grpc_opt=paths=source_relative ^
 --custom_out=./pb --custom_opt=paths=source_relative ^
  2pc.proto