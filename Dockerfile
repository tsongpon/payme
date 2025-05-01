FROM golang:1.23.0-bullseye

COPY . .
RUN go build

EXPOSE 1323

ENTRYPOINT ["./payme"]