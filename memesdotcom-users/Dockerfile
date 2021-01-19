FROM golang:1.14.3-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .
EXPOSE 8085

CMD ["go", "run", "main.go"]