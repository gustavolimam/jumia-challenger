.DEFAULT_GOAL := run

run: 
	go mod tidy && go run cmd/main.go

build_react_app:
	cd web/frontend && npm run build

mock_repository:
	cd internal/repository && mockery --all --output=mocks

mock_service:
	cd internal/service && mockery --all --output=mocks

test:
	go test ./...