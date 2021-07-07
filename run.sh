#!/bin/bash

#for i in `find api -name *.proto|xargs`;do
#    protoc --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $i
#done

for i in `find server -name *.proto|xargs`;do
    protoc --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $i
done

cd server/data && ent generate ./ent/schema && cd ../..
cd server && mkdir -p bin && go build -o bin/ ./... && ./bin/server
