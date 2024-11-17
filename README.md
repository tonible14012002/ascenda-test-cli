# Hotel Data Merger CLI Tool

This is a CLI tool built in Go for fetching, cleaning, merging, and organizing hotel data from multiple suppliers. The tool allows users to query hotel information based on specific hotel and destination IDs and outputs the results in a structured JSON format.

## Features

- Fetches raw hotel data from three suppliers.
- Cleans and merges data to produce a consolidated and complete dataset.
- Filters results based on provided `hotel_ids` and `destination_ids`.
- Returns hotel data in a structured JSON format.

## Development
### Prerequisites
- `Go 1.22.2` or later
- Bash

### Installation

1. Clone the repo

```bash
git clone git@github.com:tonible14012002/ascenda-test-cli.git
cd ascenda-test-cli
```
2. Install dependencies

```bash
make setup
go mod download # if you dont have make command
```

3. Ensure the runner script is executable:

```bash
chmod +x runner
```

## Usage

- `hotel_ids`: A comma-separated list of hotel IDs to filter by. Use none if no hotel ID filtering is needed.
- `destination_ids`: A comma-separated list of destination IDs to filter by. Use none if no destination ID filtering is needed.

```bash
./runner <hotel_ids> <destination_ids>
```
## Technical Document

1. **Project structure**
- `cmd/`: contains main.go files to run the program
- `bin/`: contains build files
- `core`: handling and define business logic and domains, soul of the app
   - `domain/`: define entities
   - `port/`: define interface to enforce services to follow so it can access repository package
   - `services/`: handling business logic 
- `internal/`: contains external, replacable and third-party services like apis, cache, database,...
- `logger/`: setup debug logging
- `utils/`: contains most reusable functions
- `.golangci.yaml`: specify linters to apply for source code
- `go.mod`: contain and manage dependencies, specifies go version
- `go.sum`: contains checksum of dependencies listed in `go.mod`, ensure those not be tempered
- `Makefile`: contain command lines for setup and running the application
