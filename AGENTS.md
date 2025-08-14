## Build, Lint, and Test

- **Build:** `go build -o roll -v ./...`
- **Test:** `go test -v ./...`
- **Run a single test:** `go test -v -run ^TestSumOfRolls$ learn-go/dice-roll-go`
- **Lint:** `golangci-lint run`
- **Tidy:** `go mod tidy`

## Code Style

- **Formatting:** Use `gofmt` for code formatting.
- **Imports:** Group imports into standard library and third-party library imports.
- **Types:** Use `big.Int` for dice roll calculations to prevent overflow.
- **Naming:** Use `camelCase` for variables, constants, functions, and structs.
- **Error Handling:** Check for errors and print to `os.Stderr`, then exit with status 1.
- **Comments:** Add comments to explain the purpose of functions and complex logic.
- **Dependencies:** Use `go mod` for dependency management.
- **Structure:** The project is structured with a `main` package containing the core logic.
- **Concurrency:** The project does not currently use concurrency.
- **Testing:** Write unit tests for new functionality and place them in `_test.go` files.
