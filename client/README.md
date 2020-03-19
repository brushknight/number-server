# Number server tester
This application connects to TCP [server](https://github.com/brushknight/number-server) and spam it with random numbers, ololo, terminate message

## Actual server

This is a [server](https://github.com/brushknight/number-server) that opens a socket and restricts input to at most 5 concurrent clients. Clients will connect to the Application and write any number of 9 digit numbers, and then close the connection. The Application must write a de-duplicated list of these numbers to a log file in no particular order.

## Prepare (build)
1. `go build  -o ./number_server_client main.go`

## Run
1. `./number_server_client` 

### Additional flags
`-h` - HELP for flags

`-mode` - mode to test (default 1, allowed: 1 - load, 2 - random errors, 3 terminate)

`-numbers` - number of numbers to send (default 10M)

`-interface` - an interface to connect (default 0.0.0.0:4000)

### Flag usage

`./number_server_client -mode=3` - will set max concurrent clients to 3
