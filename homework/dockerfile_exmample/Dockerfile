FROM golang:1.17-alpine
COPY ./httpserver  /go/src/httpserver
WORKDIR /go/src/httpserver/
RUN go build -o /bin/httpserver
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .
ENTRYPOINT ["/bin/httpserver"]
EXPOSE 80
