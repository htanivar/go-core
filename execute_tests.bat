go test ./... -v

@REM execute all the tests without gettig it from cache
@REM go test -count=1 ./...

@REM go test ./... -coverprofile=coverage.out
@REM go tool cover -html=coverage.out

@REM go test -cover ./functions
@REM golangci-lint run
