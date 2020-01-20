SHA     := $(shell git rev-parse --short HEAD)
VERSION := $(shell cat VERSION)-$(SHA)

OUTDIRNAME := bin
OUTDIR := ./$(OUTDIRNAME)

all: generate $(OUTDIR)/thing

$(OUTDIR)/thing: VERSION
	go build -o $@ -ldflags "-X main.version=$(VERSION)" cmd/thing/main.go 

generate: generate/api generate/render

generate/api:
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package thingiverse -generate client,types api/openapi-spec/thingiverse.yaml  > pkg/api/thingiverse/api.gen.go

generate/render:
	go get github.com/wlbr/templify
	go run github.com/wlbr/templify -p render -o pkg/render/item_md.gen.go pkg/render/item.md.tmpl
	go run github.com/wlbr/templify -p render -o pkg/render/list_md.gen.go pkg/render/list.md.tmpl
	go run github.com/wlbr/templify -p render -o pkg/render/item_txt.gen.go pkg/render/item.txt.tmpl
	go run github.com/wlbr/templify -p render -o pkg/render/list_txt.gen.go pkg/render/list.txt.tmpl
