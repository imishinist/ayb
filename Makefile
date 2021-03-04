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

.PHONY: deploy
deploy:
	@gomplate -f app.yaml.tmpl | tee app.yaml && \
	gcloud app deploy || rm app.yaml
