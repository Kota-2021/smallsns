.PHONY: run test lint sqlc help

# アプリケーションの実行
run:
	go run main.go

# テストの実行
test:
	go test -v ./...

# 静的解析と整形を一括実行
lint:
	gofmt -l -w .
	go vet ./...
	golangci-lint run

# sqlc によるコード生成
sqlc:
	sqlc generate

# ヘルプコマンド (make とだけ打った時に説明を表示)
help:
	@echo "Usage:"
	@echo "  make run    - アプリケーションを起動"
	@echo "  make test   - テストを実行"
	@echo "  make lint   - コードの整形とリンターの実行"
	@echo "  make sqlc   - SQLからGoのコードを生成"