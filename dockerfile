FROM golang:1.21-alpine3.18 AS builder

RUN apk add --no-cache git gcc musl-dev
# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# RUN GO111MODULE=on CGO_ENABLED=1
RUN go env -w GO111MODULE=on
RUN go env -w CGO_ENABLED=1

RUN go build -o interview-rbh .


FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/interview-rbh .
EXPOSE 8080
CMD ["./interview-rbh"]
