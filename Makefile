.PHONY: fmt lint test testacc build install generate

default: fmt lint install generate

build:
	@go build -v ./...

lint:
	@golangci-lint run

fmt:
	@gofmt -s -w -e .

test:
	@go test -v -cover -timeout=120s -parallel=10 ./...

testacc:
	TF_ACC=1 go test -v -cover -timeout 120m ./...

publish:
	cd dist && \
		zip -r -j helixer-zitadel_$(version:v%=%)_linux_amd64.zip terraform-provider-zitadel_$(version:v%=%)_linux_amd64/* && \
		zip -r -j helixer-zitadel_$(version:v%=%)_linux_arm64.zip terraform-provider-zitadel_$(version:v%=%)_linux_arm64/* && \
		zip -r -j helixer-zitadel_$(version:v%=%)_darwin_amd64.zip terraform-provider-zitadel_$(version:v%=%)_darwin_amd64/* && \
		zip -r -j helixer-zitadel_$(version:v%=%)_darwin_arm64.zip terraform-provider-zitadel_$(version:v%=%)_darwin_arm64/* && \
		depot provider helixer zitadel $(version:v%=%) 6.0 -f -k $(key) -r https://depot.services.helixer.io
