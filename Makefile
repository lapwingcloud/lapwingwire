.PHONY: dev
dev:
	go mod tidy -C controller
	go generate -C controller ./
	npm install -C controller-ui
	docker compose up -d --remove-orphans
