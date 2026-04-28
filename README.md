# TCP to HTTP Server in Go or as I like to call it gottp :)

A minimal HTTP server built directly on top of TCP sockets, without using any high-level HTTP libraries.

This project focuses on understanding how web servers actually work under the hood by manually handling connections, parsing requests, and constructing HTTP responses.

---

## Features:

* Raw TCP server using Go's `net` package
* Manual HTTP/1.1 request parsing
* Custom response builder
* Basic routing
* Lightweight and dependency-free

---

## Concepts Covered:
Since modern frameworks abstract so much networking and http...
It was a great experience going through all these concepts and having an even deeper understanding of:

* TCP connections and sockets
* Byte stream handling
* HTTP protocol structure
* Request/response lifecycle
* Headers and status codes
* Stateless communication model

---

## Tech Stack

* **Language:** Go
* **Networking:** `net` package
* **Protocol:** TCP / HTTP 1.1

---

## How It Works

1. Starts a TCP listener on a specified port
2. Accepts incoming client connections
3. Reads raw byte streams from the connection
4. Parses the HTTP request manually
5. Builds a valid HTTP response
6. Writes the response back to the client

---

## Running the Project

```bash
go run main.go
```
