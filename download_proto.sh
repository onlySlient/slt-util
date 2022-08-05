#/bin/bash

set -e

para="-x http://x.x.x.x:7890 --retry 1 --retry-delay 1 --retry-connrefused -m 1 --connect-timeout 1"

fileList=("any" "api" "descriptor" "duration" "empty" "field_mask" "source_context" "struct" "timestamp" "type" "wrappers")

echo api
mkdir -p proto/include/google/api/
curl $para -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > proto/include/google/api/annotations.proto
curl $para -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto > proto/include/google/api/field_behavior.proto
curl $para -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > proto/include/google/api/http.proto
curl $para -sSL https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto > proto/include/google/api/httpbody.proto

echo protobuf
mkdir -p proto/include/google/protobuf/
for file in ${fileList[*]}
do
    echo $file
    curl $para -sSL https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/$file.proto > proto/include/google/protobuf/$file.proto
done

echo options
mkdir -p proto/include/protoc-gen-openapiv2/options/
curl $para -sSL https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-openapiv2/options/annotations.proto > proto/include/protoc-gen-openapiv2/options/annotations.proto
curl $para -sSL https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-openapiv2/options/openapiv2.proto > proto/include/protoc-gen-openapiv2/options/openapiv2.proto

echo validate
mkdir -p proto/include/validate/
curl $para -sSL https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/master/validate/validate.proto > proto/include/validate/validate.proto