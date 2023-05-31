.PHONY: dev
dev:
	docker compose up -d
	go mod tidy
	go generate ./controller/ent
	air --build.cmd "go build -o tmp/controller controller/main.go" --build.bin "tmp/controller" --build.full_bin "tmp/controller | jq"
