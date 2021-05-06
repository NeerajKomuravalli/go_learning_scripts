# Command to run proto buf
protoc -I models/proto/ models/proto/node.proto --go_out=plugins=grpc:./models/proto/
