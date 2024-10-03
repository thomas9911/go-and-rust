default: 
    @just --list

build: rust-build go-build

run *args:
    ./out/main {{args}}

rust-build:
    cargo build --release
    @cp target/release/libgo_and_rust.a out/
    # @cp target/release/go_and_rust.dll out/
    @just generate-headers

build-run *args: build
    @just run {{args}}

go-build:
    # CGO_ENABLED=1 go build -ldflags="-r {{justfile_directory()}}\out" main.go
    CGO_ENABLED=1 go build -ldflags "-s -w" -o out/main main.go

generate-headers:
    cargo run --bin generate-headers --features headers
