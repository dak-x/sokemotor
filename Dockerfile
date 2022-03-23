FROM golang:1.16-buster

RUN go get github.com/cespare/reflex

COPY reflex.conf /


