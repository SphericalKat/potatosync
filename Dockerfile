FROM golang

WORKDIR /go/src/github.com/ATechnoHazard/potatonotes-api

COPY . .

ENTRYPOINT ["go", "run", "."]