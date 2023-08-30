# go-connect-greet
[Connect for Go.](https://connectrpc.com/docs/go/getting-started/)

## Install protoc.
- Compile the proto file and create the generated descriptor set (`.pb file).
- [Get Started Cloud Run for gRPC](https://cloud.google.com/api-gateway/docs/get-started-cloud-run-grpc?hl=ja)
- [protobuf](https://github.com/protocolbuffers/protobuf/releases)
```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip
unzip protoc-24.2-linux-x86_64.zip 
mv bin/protoc /usr/local/bin/

protoc --proto_path=greet/v1 \
    --descriptor_set_out=./descriptor.pb \
    greet.proto
```

## Create `api_config.yaml`
```yaml
# The configuration schema is defined by the service.proto file.
# https://github.com/googleapis/googleapis/blob/master/google/api/service.proto

type: google.api.Service
config_version: 3
# In the name field, a service named *.apigateway.PROJECT_ID.cloud.goog where PROJECT_ID is the name of your Google Cloud project ID.
name: "*.apigateway.shinonome-375705.cloud.goog"
title: API Gateway + Cloud Run gRPC
apis:
  # package + service id
  - name: greet.v1.GreetService
usage:
  rules:
  # ListShelves methods can be called without an API Key.
  # package + service id + rpc
  - selector: greet.v1.GreetService.Greet
    allow_unregistered_calls: true
backend:
  rules:
    # Note: To use gRPC with TLS on Cloud Run, the address field must have the scheme grpcs:// instead of https://.
    - selector: "*"
      address: grpcs://go-connect-greet-zugntxuibq-an.a.run.app
```

## Environments.
```bash
export PROJECT_ID="shinonome-375705"
export APIGATEWAY_API="shinonome-grpc"
export APIGATEWAY_CONFIG_ID="grpc-config"
export GATEWAY_ID="shinonome-gw"
export GATEWAY_REGION="asia-northeast1"
```
## Create api-configs.
```bash
gcloud api-gateway api-configs create $APIGATEWAY_CONFIG_ID \
--api=$APIGATEWAY_API --project=$PROJECT_ID \
--grpc-files=descriptor.pb,api_config.yaml
```

```bash
gcloud api-gateway api-configs describe $APIGATEWAY_CONFIG_ID \
  --api=$APIGATEWAY_API --project=$PROJECT_ID
```

## Create API Gateway.
```bash
gcloud api-gateway gateways create $GATEWAY_ID \
  --api=$APIGATEWAY_API --api-config=$APIGATEWAY_CONFIG_ID \
  --location=$GATEWAY_REGION --project=$PROJECT_ID
```

## Sending a request to the API.
```bash
curl \
    --header "Content-Type: application/json" \
    --data '{"name": "Jane"}' \
https://shinonome-gw-8g4o26at.an.gateway.dev/greet.v1.GreetService/Greet
```
