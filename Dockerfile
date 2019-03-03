FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY ./quotes.json .
RUN go build -o main .
CMD ["/app/main"]