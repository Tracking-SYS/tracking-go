#!/bin/bash

set -ex

ROOT=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

export $(xargs < ${ROOT}/VERSION)

curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > proto/google/api/annotations.proto
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > proto/google/api/http.proto

mkdir -p proto/validate
curl https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/v${PROTOC_GEN_GO_VALIDATE_VERSION}/validate/validate.proto > proto/validate/validate.proto

mkdir -p  proto/protoc-gen-openapiv2/options
curl https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/v${GRPC_GATEWAY_VERSION}/protoc-gen-openapiv2/options/annotations.proto > proto/protoc-gen-openapiv2/options/annotations.proto
curl https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/v${GRPC_GATEWAY_VERSION}/protoc-gen-openapiv2/options/openapiv2.proto > proto/protoc-gen-openapiv2/options/openapiv2.proto