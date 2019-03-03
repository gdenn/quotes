# Use golang base image
FROM golang:latest

# Set working directory
WORKDIR /go/src/github.com/gdenn/quotes

# Add local code to working directory
ADD . .

# Compule go service
RUN go build -o main

# Set executable
ENTRYPOINT ./main

# Expose 8080 port
EXPOSE 8080