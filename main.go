package main

import (
	"log"
	"login-demo/internal/routers"
	"net/http"
	"time"
)

/**
* @Author: super
* @Date: 2020-09-29 10:34
* @Description:
**/

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Start to listening the incoming requests on http address :%d", 8080)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("start listen server err: %v", err)
	}
}