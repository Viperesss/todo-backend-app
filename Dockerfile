FROM golang:1.27-bookworm

WORKDIR /todoapp

COPY . .

RUN go mod tidy
RUN go build -o /todoapp/exe main.go

CMD ["/todoapp/exe"]