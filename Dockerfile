FROM golang:1.16 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o hackme

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/hackme /hackme

CMD ["/hackme"]