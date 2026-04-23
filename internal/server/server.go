package server

import (
	"bytes"
	"fmt"
	"gottp/internal/request"
	"gottp/internal/response"
	"io"
	"net"
)

type HandlerError struct {
	StatusCode response.StatusCode
	Message    string
}

type Handler func(w io.Writer, req *request.Request) *HandlerError

type Server struct {
	closed  bool
	handler Handler
}

func runConn(s *Server, conn io.ReadWriteCloser) {
	defer conn.Close()

	headers := response.GetDefaultHeaders(0)
	r, err := request.RequestFromReader(conn)
	if err != nil {
		response.WriteStatusLine(conn, response.StatusBadRequest)
		response.WriteHeaders(conn, headers)
		return
	}

	writer := bytes.NewBuffer([]byte{})
	handlerError := s.handler(writer, r)

	if handlerError != nil {
		response.WriteStatusLine(conn, handlerError.StatusCode)
		response.WriteHeaders(conn, headers)
		conn.Write([]byte(handlerError.Message))
		return
	}

	body := writer.Bytes()
	headers.Replace("Content-length", fmt.Sprintf("%d", len(body)))

	response.WriteStatusLine(conn, response.StatusOk)
	response.WriteHeaders(conn, headers)
	conn.Write(body)
}

func runServer(s *Server, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if s.closed {
			return
		}
		if err != nil {
			return
		}
		go runConn(s, conn)
	}
}

func Serve(port uint16, handler Handler) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &Server{
		closed:  false,
		handler: handler,
	}
	go runServer(server, listener)

	return server, nil
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}
