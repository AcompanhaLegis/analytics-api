FROM golang:1.21

WORKDIR /etc/acompanhalegis

RUN mkdir -p /etc/acompanhalegis/certs

COPY . .

RUN go get .