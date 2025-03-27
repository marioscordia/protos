install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen-sso-code:
	protoc -I proto -I vendor.protogen \
  		proto/sso/sso.proto \
  		--go_out=./gen/go --go_opt=paths=source_relative \
  		--go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative \
  		--grpc-gateway_out=./gen/go --grpc-gateway_opt=paths=source_relative


gen-notes-code:
	protoc -I proto -I vendor.protogen \
	  	proto/notes/notes.proto \
  		--go_out=./gen/go --go_opt=paths=source_relative \
  		--go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative \
  		--grpc-gateway_out=./gen/go --grpc-gateway_opt=paths=source_relative


gen-updates-code:
	protoc -I proto -I vendor.protogen \
  		proto/updates/updates.proto \
  		--go_out=./gen/go --go_opt=paths=source_relative \
  		--go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative \
  		--grpc-gateway_out=./gen/go --grpc-gateway_opt=paths=source_relative

	
vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi