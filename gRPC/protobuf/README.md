#How to define a protocol message?
```
    Name of the message: UpperCammelCase
    Name of the field: lower_snake_case
    Data types:
        - string, bool, bytes
        - float, double
        - int32, int64, uint32, uint64, sint32, sint64 etc...
    Data type can be user defined enums
```

##Project Folder structure
```
    protobuf/
        pb/
        proto/
            example.proto
```

#How To Generate Go Code And GRPC Code in pb/ Folder?

>Run make gen
Manually run below command
```go
    protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto
```