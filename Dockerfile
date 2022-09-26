FROM golang:1.19-alpine

WORKDIR /

COPY go.mod /
COPY go.sum /
RUN go mod download

COPY . .

RUN go build -o game

EXPOSE 8080

CMD [ "/game" ]
