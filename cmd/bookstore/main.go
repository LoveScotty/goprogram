package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LoveScotty/goprogram/internal/store/factory"
	_ "github.com/LoveScotty/goprogram/internal/store/mem"
	"github.com/LoveScotty/goprogram/server/bookstore"
)

func main() {
	s, err := factory.New("bookstore")
	if err != nil {
		panic(err)
	}
	srv := bookstore.NewServer(":8888", s)

	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Println("bookstore server start failed:", err)
		return
	}
	log.Println("bookstore server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// 优雅退出
	select {
	case err = <-errChan:
		log.Println("bookstore server run failed:", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("bookstore program exit error:", err)
		return
	}
	log.Println("bookstore program exit ok")
}
