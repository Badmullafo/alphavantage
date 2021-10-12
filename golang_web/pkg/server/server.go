package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Badmullafo/alphavantage/golang_web/pkg/request"
)

type resultHandler struct {
	mu  sync.Mutex // guards n
	c   int
	res request.Result
}

var doOnce sync.Once

type Handlers interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (h *resultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	h.c++
	h.mu.Unlock()
	fmt.Fprintf(w, "The %s is %.2f\n", h.res.Dtype, h.res.Value)
	fmt.Fprintf(w, "Stock symbol is %s\n", h.res.Symbol)
	fmt.Fprintf(w, "The count is %d\n", h.c)
}

func newHandler(r *request.Result) Handlers {
	return &resultHandler{res: *r}
}

func Startserver(ctx context.Context, rChan <-chan *request.Result) {

	r := <-rChan

	log.Printf("Starting server at port 8080\n")

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        newHandler(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	doOnce.Do(func() {
		http.Handle("/"+r.Dtype, srv.Handler)
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()

	log.Printf("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown error: %v\n", err)
	} else {
		log.Printf("gracefully stopped\n")
	}

}
