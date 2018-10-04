package main

import (
	"bufio"
	"fmt"
	"os"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"

	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Requested path is %q", ctx.Path())
	}

	withCors := cors.NewCorsHandler(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type"},
	})

	s := &fasthttp.Server{
		Handler: withCors.CorsMiddleware(requestHandler),
	}

	go s.ListenAndServe("127.0.0.1:8080")
	fmt.Println("Server started! Press ENTER to exit.")

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fmt.Println("Bye!")
		os.Exit(0)
	}
}
