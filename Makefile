GO = go

.PHONY: build/cli
build/cli:
	$(GO) build -o bin/ayb .

.PHONY: build/http
build/http:
	$(GO) build -o bin/ayb ./pkg/handler/ayb

.PHONY: generate
generate:
	cd pkg/witticism/ && statik -m -src=. -include="*.json"
