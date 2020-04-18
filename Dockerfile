FROM golang:1.13.4-alpine as backend

# Install dependencies
WORKDIR /go/src/github.com/abatilo/contract-ai
COPY ./go.mod ./go.sum ./
RUN go mod download

# Build artifacts
WORKDIR /go/src/github.com/abatilo/contract-ai
COPY ./cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/contract-ai cmd/contract-ai.go

FROM scratch
# SSL Certs
COPY --from=backend /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=backend /go/bin/contract-ai /go/bin/contract-ai

# Run the hello binary.
ENTRYPOINT ["/go/bin/contract-ai", "api"]
