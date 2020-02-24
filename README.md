# Number server

This is an application that opens a socket and restricts input to at most 5 concurrent clients. Clients will connect to the Application and write any number of 9 digit numbers, and then close the connection. The Application must write a de-duplicated list of these numbers to a log file in no particular order.

## Prepare (build)
clone this project to folder number-server into your $GOPATH and build
1. `git clone git@github.com:brushknight/number-server.git ~/go/src/number-server`
2. `cd ./number-server`
3. `go build number-server.go`

## Tests
1. `go test ./...`

## Run
1. `./number-server` 

### Additional flags
`-h` - HELP for flags

`-max-clients` - maximum of concurrent clients (default 5)

`-interface` - an interface to startup (default 0.0.0.0:4000)

`-log-file` - path to a log file (default ./numbers.log)

`-env` - an environment (default prd, allowed prd, dev)

## Code structure

![code_structure](resources/images/code_structure.png)

## Application design

![application_design](resources/images/application_design.png)


## Tester
@TBD (add link to tester repository)