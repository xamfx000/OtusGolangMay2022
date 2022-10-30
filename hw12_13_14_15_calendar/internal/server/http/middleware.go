package internalhttp

import (
	"fmt"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter{ResponseWriter: w}
		start := time.Now()
		next.ServeHTTP(&sw, r)
		duration := time.Since(start)
		if sw.status == 0 {
			sw.status = 200
		}
		fmt.Printf("%s [%s] %s %s HTTP/%d.%d %d %dµs \"%s\"\n",
			r.RemoteAddr,
			time.Now().Format(time.RFC3339),
			r.Method,
			r.RequestURI,
			r.ProtoMajor,
			r.ProtoMinor,
			sw.status,
			duration,
			r.UserAgent(),
		)
	})
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}
