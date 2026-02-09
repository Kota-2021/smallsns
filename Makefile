.PHONY: run test lint sqlc help

# デフォルトのターゲットを help に設定
.DEFAULT_GOAL := help

# アプリケーションの実行
run:
	go run ./backend/cmd/api/main.go

# テストの実行
test:
	go test -v ./backend/...

# 静的解析と整形を一括実行
lint:
	gofmt -l -w backend/
	go vet ./backend/...
	-golangci-lint run ./backend/...

# sqlc によるコード生成
sqlc:
	sqlc generate -f backend/sqlc.yaml

# ヘルプコマンド (make とだけ打った時に説明を表示)
help:
	@echo "Available commands:"
	@echo "  make run    - Start the application"
	@echo "  make test   - Run tests"
	@echo "  make lint   - Formatting and Linting"
	@echo "  make sqlc   - Generate Go code from SQL"