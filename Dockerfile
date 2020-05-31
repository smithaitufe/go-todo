FROM golang:latest

RUN mkdir /go-todo 

WORKDIR /go-todo

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/todo/*.go

EXPOSE 5000

CMD ["./main"]