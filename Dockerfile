FROM golang:1.23.2

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o taskify

EXPOSE 8080

CMD ["./taskify"]
