package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartHTTPServer(port string, handler http.Handler) {
	errs := make(chan error, 2)

	go func() {
		log.Printf("Listening and serving HTTP on %v", port)
		errs <- http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	}()

	go func() {
		stopChan := make(chan os.Signal)
		signal.Notify(stopChan, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-stopChan)
	}()

	log.Fatalf("Terminated. Cause: %v", <-errs)
}
