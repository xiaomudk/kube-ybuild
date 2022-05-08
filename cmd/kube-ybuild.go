package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/xiaomudk/kube-ybuild/internal/server"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	if err := server.NewHTTPServer().Run(); err != nil {
		panic(err)
	}
}
