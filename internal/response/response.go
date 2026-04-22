package response

import (
	"fmt"
	"gottp/internal/headers"
	"io"
)

type Response struct {
}

type StatusCode int

const (
	StatusOk                  StatusCode = 200
	StatusBadRequest          StatusCode = 400
	StatusInternalServerError            = 500
)

func GetDefaultHeaders(contentLen int) *headers.Headers {
	h := headers.NewHeaders()
	h.Set("Content-Length", fmt.Sprintf("%d", contentLen))
	h.Set("Content-Type", "text/plain")
	h.Set("Connection", "closed")

	return h
}

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	statusLine := []byte{}

	switch statusCode {
	case StatusOk:
		statusLine = []byte("HTTP/1.1 200 OK")
	case StatusBadRequest:
		statusLine = []byte("HTTP/1.1 400 Bad Request")
	case StatusInternalServerError:
		statusLine = []byte("HTTP/1.1 500 Internal Server Error")
	default:
		return fmt.Errorf("unrecognized error code")
	}

	_, err := w.Write(statusLine)
	return err
}
