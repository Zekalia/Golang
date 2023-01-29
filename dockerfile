FROM golang:1.19.4-alpine as dev

WORKDIR /work

COPY . .

RUN go build -o api-product

CMD ./api-product