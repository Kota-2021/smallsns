package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// ... ハンドラの設定 ...

	// Graceful Shutdownへ対応させる(runを使用せずにListenAndServeを使用する)
	// 1. HTTPサーバーを構造体として定義（ListenAndServe を goroutine で動かすため）
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// 2. サーバー起動にListenAndServeを使ってgoroutine で実行
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 3. 終了シグナルを待機するチャネルを作成
	// SIGINT (Ctrl+C) と SIGTERM (ECSからの停止命令) を受け取る
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // ここでシグナルが来るまでブロック
	log.Println("Shutdown Server ...")

	// 3-1. HTTPサーバーの停止（猶予時間を設ける）
	// ECS 側の stopTimeout（デフォルト 30 秒）より短く設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// 3-2. WebSocket Hub の停止
	// hub.Stop() などを呼び出して、内部の goroutine を安全に終了させる

	// 3-3. DB 接続のクローズ
	// 処理中のクエリが完了してから閉じること

	log.Println("Server exited properly")
}
