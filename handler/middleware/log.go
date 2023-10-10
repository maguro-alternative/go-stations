package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TechBowl-japan/go-stations/model"
)

func GetLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("start")
		h.ServeHTTP(w, r)
		OS, err := GetOS(r.Context())
		if err != nil {
			log.Println(err)
			return
		}
		res := model.AccessLog{
			TimeStamp: start,
			Latency:   time.Since(start).Milliseconds(),
			Path:      r.URL.Path,
			OS:        OS,
		}

		bytes, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(bytes))
		log.Println("end")
		log.Printf("Process took %s\n", time.Since(start))
	})
}
