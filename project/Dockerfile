
FROM golang:1.21.0 as builder
WORKDIR /usr/src/app
COPY . .
RUN go mod download

#EXPOSE 4000

CMD ["go", "run", "/usr/src/app/cmd/epl", "."]


