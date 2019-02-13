### Build application
1. `go build debate-api`

*Note*: Make sure project is in $GOPATH

### Endpoints
Api runs on `localhost:8080`

GET `/opinion` with query param `url`

Example: `localhost:8080/opinion?url=https://www.debate.org/opinions/should-drug-users-be-put-in-prison`

### Run tests
1. `go test debate-api...`

#### Improvements if I had more time
1. More tests (testing successes, errors, etc)
1. Split code into components and test individually mocking out external components
