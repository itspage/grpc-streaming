# grpc-streaming

Example Golang implementation of GRPC streaming.

## Run

`go get github.com/itspage/grpc-streaming`

Start server:

`go run server/server.go`

Run the client:

`go run client/client.go`

```
2017/02/11 09:45:25 Greeting: Hello (0) Dave
2017/02/11 09:45:25 Greeting: Hello (1) Dave
2017/02/11 09:45:25 Greeting: Hello (2) Dave
2017/02/11 09:45:25 Greeting: Hello (3) Dave
2017/02/11 09:45:25 Greeting: Hello (4) Dave
2017/02/11 09:45:25 Greeting: Hello (5) Dave
2017/02/11 09:45:25 Greeting: Hello (6) Dave
2017/02/11 09:45:25 Greeting: Hello (7) Dave
2017/02/11 09:45:25 Greeting: Hello (8) Dave
2017/02/11 09:45:25 Greeting: Hello (9) Dave
```
