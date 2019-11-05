# Stage 1: Build binary
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/ATechnoHazard/potatosync

COPY . .

RUN go get -v -t -d

RUN GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -a -installsuffix cgo -o /go/bin/potatosync

# Stage 2: Copy binary to scratch 
FROM scratch

COPY --from=builder /go/bin/potatosync /go/bin/potatosync

ENTRYPOINT ["/go/bin/potatosync"]
