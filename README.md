# growlog-server
**TODO**

## Installation

Get our latest code.

```bash
go install github.com/bartmika/growlog-server@latest
```

## Usage

To start the server, run the following command in your **terminal**:

```bash
export GROWLOG_DATABASE_URL="postgres://golang:123passwordd@localhost:5432/growlog_db"
$GOBIN/growlog-server serve
```

That's it! If everything works, you should see a message saying `gRPC server is running.`.

## Sub-Commands Reference

### ``serve``

**Details:**

```text
Run the gRPC server to allow other services to access the application

Usage:
  growlog-server serve [flags]

Flags:
  -h, --help                           help for serve
  -p, --port int                       The port to run this server on (default 50051)
```

**Example:**

```bash
$GOBIN/tstorage-server serve -p=50051
```

## Contributing
### Development
If you'd like to setup the project for development. Here are the installation steps:

1. Go to your development folder.

    ```bash
    cd ~/go/src/github.com/bartmika
    ```

2. Clone the repository.

    ```bash
    git clone https://github.com/bartmika/growlog-server.git
    cd growlog-server
    ```

3. Install the package dependencies

    ```bash
    go mod tidy
    ```

4. In your **terminal**, make sure we export our path (if you haven’t done this before) by writing the following:

    ```bash
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

5. Run the following to generate our new gRPC interface. Please note in your development, if you make any changes to the gRPC service definition then you'll need to rerun the following:

    ```bash
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/growlog.proto
    ```

6. You are now ready to start the server and begin contributing!

    ```bash
    go run main.go serve
    ```

### Quality Assurance

Found a bug? Need Help? Please create an [issue](https://github.com/bartmika/growlog-server/issues).


## License

[**ISC License**](LICENSE) © Bartlomiej Mika
