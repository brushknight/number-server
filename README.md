# Number server

This project contains TCP server that process 9 digit numbers from clients, find all uniques and store them into a log file.

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

## Tester
@TBD (add link to tester repository)