.PHONY: openapi
aether-models:
	oapi-codegen -generate types -o aether_types.gen.go -package main aether-2.0.0-openapi3.yaml
#	oapi-codegen -generate chi-server -o aether_server.gen.go -package main aether-2.0.0-openapi3.yaml
	oapi-codegen -generate server -o aether_server.gen.go -package main aether-2.0.0-openapi3.yaml

build:
	rm -rf  build/_output;mkdir -p  build/_output
	go build -o build/_output/roc-api ./cmd/roc-api

image:
	docker build -t roc-api .

