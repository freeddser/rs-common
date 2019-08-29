package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type Middleware func(http.HandlerFunc, ...interface{}) http.HandlerFunc

//1 2 3--> 3 2 1
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func RecoverHandler() Middleware {
	middleware := func(next http.HandlerFunc, args ...interface{}) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			e, ok := recover().(error)
			if ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
				// logging
				log.Println("WARN: panic fired in %v.panic - %v", next, e)
				log.Println(string(debug.Stack()))
			}
			next(w, r)
		}
		return handler
	}
	return middleware
}

func Timer() Middleware {
	middleware := func(next http.HandlerFunc, args ...interface{}) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			defer func(begin time.Time) {
				fmt.Println(r.Method, r.RequestURI, r.Proto, "----->", time.Since(begin))
			}(time.Now())
			next(w, r)
		}
		return handler
	}
	return middleware
}

func LogerClient() Middleware {
	middleware := func(next http.HandlerFunc, args ...interface{}) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			//fmt.Println("TLS:", r.TLS)
			//fmt.Println("server:", r.Host, "client:", r.RemoteAddr)
			//fmt.Println(r.Method, r.RequestURI, r.Proto)
			//for key, value := range r.Header {
			//	fmt.Println(key, ":", value)
			//}
			//if r.Method == "POST" {
			//	r.ParseForm()
			//	fmt.Println("POST")
			//	fmt.Println(r.PostForm)
			//
			//}
			next(w, r)
		}
		return handler
	}
	return middleware
}
