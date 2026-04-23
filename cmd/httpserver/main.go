package main

import (
	"gottp/internal/request"
	"gottp/internal/response"
	"gottp/internal/server"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port = 42069

func main() {
	s, err := server.Serve(port, func(w io.Writer, req *request.Request) *server.HandlerError {
		if req.RequestLine.RequestTarget == "/yourproblem" {
			return &server.HandlerError{
				StatusCode: response.StatusBadRequest,
				Message:    "your problem\n",
			}
		}
		if req.RequestLine.RequestTarget == "/myproblem" {
			return &server.HandlerError{
				StatusCode: response.StatusInternalServerError,
				Message:    "oopsie problem\n",
			}
		} else {
			w.Write([]byte("All good fr, fr\n"))
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer s.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
