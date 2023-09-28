package middleware

import "net/http"

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: ここに実装をする
		defer func() {
			// panicが発生したら500を返す
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// 通常の処理を実行
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}