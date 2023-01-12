FROM golang:1.18 as builder

COPY . /entimport
WORKDIR /entimport/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzOpenImport.a && \
    clang -fsanitize=fuzzer fuzzOpenImport.a -o fuzzOpenImport.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /entimport/mayhem/fuzzOpenImport.libfuzzer /