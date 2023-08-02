FROM golang:alpine
WORKDIR /app
ENV TELEGRAM_APITOKEN=YOUR_APITOKEN
ENV OPENAI_APITOKEN=YOUR_APITOKEN
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go
EXPOSE 8080
CMD ["./main"]
