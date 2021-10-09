package server

import (
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

func Startserver(r *request.Result) {

	fmt.Printf("Starting server at port 8080\n")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        newHandler(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.Handle("/"+r.Dtype, s.Handler)
	log.Fatal(s.ListenAndServe())

}
