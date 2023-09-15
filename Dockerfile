FROM golang:1.21.0-bookworm

RUN mkdir app
WORKDIR /app

COPY . .

CMD [ "go", "run", "sample.go" ]
