FROM golang:1.20.7-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  user

WORKDIR $GOPATH/src/golangginrestservice/app/

COPY . .

RUN git config --global http.sslVerify "false"

RUN GOSUMDB=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5
RUN $GOPATH/bin/swagger generate spec -o ./cmd/api/swaggerui/swagger.json --scan-models

# Build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -p 1 -o /golangginrestservice main.go

# Run Stage
#FROM scratch
FROM alpine:3.17.0

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=base /golangginrestservice .
COPY --from=base go/src/golangginrestservice/app/ .

USER user:user

EXPOSE 9001

CMD ["./golangginrestservice"]
