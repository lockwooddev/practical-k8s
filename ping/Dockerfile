FROM golang:1.15-alpine
WORKDIR /go/src/github.com/lockwooddev/practical-k8s/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /ping ping/main.go

FROM alpine:latest
COPY --from=0 /ping .