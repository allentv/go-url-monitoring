package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// Run will execute the server process with the provided config
// Reference: https://github.com/gorilla/mux#graceful-shutdown
func Run(srvConfig *Config) {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%s", srvConfig.Domain, srvConfig.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: srvConfig.WriteTimeout,
		ReadTimeout:  srvConfig.ReadTimeout,
		IdleTimeout:  srvConfig.IdleTimeout,
		Handler:      srvConfig.Handler, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), srvConfig.DefaultTimeout)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
