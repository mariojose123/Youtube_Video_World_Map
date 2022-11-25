FROM golang:latest
WORKDIR /app



COPY . .

RUN go build -o server server.go

EXPOSE 8080


CMD [ "/app/server" ]
