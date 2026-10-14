FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app .

FROM build AS test
RUN go test ./...

FROM gcr.io/distroless/base
COPY --from=build /app /app
ENTRYPOINT ["/app"]
