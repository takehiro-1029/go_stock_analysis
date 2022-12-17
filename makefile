# メタ情報

export GO111MODULE=on
# 開発に必要な依存をインストールする
## Setup for Development env
.PHONY: setup-tools
setup-tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.37.0
	go install github.com/Songmu/make2help/cmd/make2help@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/cespare/reflex@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/moznion/gonstructor/cmd/gonstructor@latest
	go install github.com/volatiletech/sqlboiler/v4@v4.6.0
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.6.0
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.5.0
	go install github.com/k0kubun/sqldef/cmd/mysqldef@v0.8.8


## run server with hot-reload
dev:
	reflex -r '\.go' -s go run main.go


## Generate types to create api handler
.PHONY: gen-handler
gen-handler:
	oapi-codegen -package handler -generate "types" -templates apitemplates api.yaml > server/api/handler/types.gen.go
	oapi-codegen -package handler -generate "chi-server" -templates apitemplates api.yaml > server/api/handler/handler_wrapper.gen.go

.PHONY: gen-dao
gen-dao:
	sqlboiler mysql

.PHONY: build
build:
	go build main.go


