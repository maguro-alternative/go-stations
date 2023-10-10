package main

import (
	"log"
	"os"
	"time"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/TechBowl-japan/go-stations/db"
	"github.com/TechBowl-japan/go-stations/handler/router"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func realMain() error {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln("main: failed to load .env, err =", err)
		return err
	}
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln("main: failed to set time zone, err =", err)
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		log.Fatalln("main: failed to set up sqlite3, err =", err)
		return err
	}
	defer todoDB.Close()

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(todoDB)

	// TODO: サーバーをlistenする
	// NOTE: ポート番号は上記のport変数を使用すること
	// NOTE: エラーが発生した場合はlog.Fatalでログを出力すること
	log.Fatal(http.ListenAndServe(port, mux))

	return nil
}
