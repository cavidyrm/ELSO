FROM golang:1.21
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o healthchain .
CMD ["/app/healthchain", "--algorithm", "elso"] # Default to ELSO