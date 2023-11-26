package cmd

import "bookkeeping/internal/router"

func RunServer() {
	r := router.New()
	r.Run()
}
