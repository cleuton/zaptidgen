# Estágio de Build da app Go
FROM golang AS build-stage
WORKDIR /app
COPY ./ ./
## Instala o protocol buffer
RUN apt-get update
RUN apt install -y protobuf-compiler golang-goprotobuf-dev 
RUN GO111MODULE=on \
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
## Baixa as dependências
RUN go mod download && go mod verify
## Compila os protocolos

RUN cd protos && protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    idgen.proto 
RUN cp protos/*.go gen

## Compila o programa
RUN go build -v -o build/zaptidgen ./zaptidgen.go

# Estágio de execução da aplicação: 
FROM alpine:3.18.5
WORKDIR /app
COPY --from=build-stage /app/build/zaptidgen .
RUN apk add libc6-compat
EXPOSE 8888
ENTRYPOINT [ "./zaptidgen" ]

