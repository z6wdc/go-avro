# go-avro

A demo project for showcasing [Apache Avro](https://avro.apache.org/) serialization in Go, using Clean Architecture and a Cobra-based CLI.

This project demonstrates:
- Encoding and decoding structured data with Avro
- Schema versioning (starting from v1)
- Separation of concerns with use cases and interfaces
- CLI interaction via `encode` and `decode` commands

## 📦 Tech Stack

- [Go](https://golang.org/) 1.22+
- [Apache Avro](https://avro.apache.org/) (via `goavro`)
- [Cobra CLI](https://github.com/spf13/cobra)
- Clean Architecture principles
- [GoMock](https://github.com/golang/mock) for unit testing
- GitHub Actions for CLI integration testing

## 🚀 Getting Started

### Build the CLI

```bash
go build -o go-avro
```

### Encode JSON to Avro binary

```bash
go-avro encode --input testdata/input.json --output testdata/output.avro
```

### Decode Avro binary to JSON

```bash
go-avro decode --input testdata/output.avro --output testdata/output.json
```

> Sample input file: `testdata/input.json`

```json
{
  "id": "cli-123",
  "userId": 99,
  "message": "Hello, Avro CLI!"
}
```

## 🧪 Running Tests

### Unit tests (with GoMock)

```bash
go test ./internal/... -v
```

### CLI integration test

Make sure the binary is built first:

```bash
go build -o go-avro
go test ./test -v
```

### Or run all tests:

```bash
go test ./... -v
```

## 🛠 Project Structure

```
cmd/                # Cobra CLI commands (encode, decode)
internal/
  entity/           # Notification domain model
  infra/avro/       # Avro codec (encoder/decoder)
  usecase/          # Use case layer
pkg/utils/          # Shared utilities (e.g. project root path)
test/               # Integration tests
testdata/           # JSON/Avro sample files
```

## 🧹 Development Tools

- Generate mocks (via go generate):

```bash
go generate ./...
```
