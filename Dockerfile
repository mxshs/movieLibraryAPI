FROM golang:1.22.1-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build

CMD ["./movieLibrary"]
