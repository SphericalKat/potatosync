FROM scratch

COPY ./bin/potatosync /

ENTRYPOINT ["/potatosync"]
