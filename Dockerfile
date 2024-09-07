FROM golang:1.22-alpine as build
LABEL authors="eduard<eduard@vesh95.ru>"

WORKDIR /opt/todo

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls -al

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

FROM debian:latest
COPY --from=build /opt/todo/bin/main /opt/todo
EXPOSE 8080
CMD /opt/todo