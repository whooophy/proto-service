
## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

install dependencies

```bash
  brew install protobuff
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

add proto file into proto folder in root project
```bash
  touch proto/{filename}.proto
```

generate proto file

```bash
  go run main.go generate:grpc {filename}
```

generate multiple proto file
```bash
go run main.go generate:grpc {filename1} {filename2} {filename3} etc...
```

Generated file will placed in pb folder in root dir