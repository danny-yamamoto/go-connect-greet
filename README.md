# go-connect-greet
Connect for Go.

```bash
export PROJECT_ID="shinonome-375705"
export APIGATEWAY_API="shinonome-grpc"
export APIGATEWAY_CONFIG_ID="grpc-config"
gcloud api-gateway api-configs create grpc-config \
--api=shinonome-grpc --project=shinonome-375705 \
--grpc-files=greet.proto,api_config.yaml
```

```bash
gcloud api-gateway api-configs describe CONFIG_ID \
  --api=$APIGATEWAY_API --project=$PROJECT_ID
```

```bash
wget https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip
unzip protoc-24.2-linux-x86_64.zip 
mv bin/protoc /usr/local/bin/

protoc --proto_path=greet/v1 \
    --descriptor_set_out=./descriptor.pb \
    greet.proto
```
