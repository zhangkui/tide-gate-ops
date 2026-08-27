package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/zhangkui/tide-gate-ops/internal/handler"
	"gitlab.com/zhangkui/tide-gate-ops/internal/service"
	"gitlab.com/zhangkui/tide-gate-ops/internal/store"
)

func main() {
	path := os.Getenv("TIDE_GATE_OPS_DB")
	if path == "" {
		path = "data/tide-gate-ops.db"
	}
	repository, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer repository.Close()
	app := service.NewLab(repository)
	defer app.Close()
	addr := os.Getenv("TIDE_GATE_OPS_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("tide-gate-ops listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler.New(app)))
}
