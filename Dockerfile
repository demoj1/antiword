FROM golang:alpine

ADD ./src /go/src/antiword

RUN go install antiword

#ENTRYPOINT /go/bin/antiword

EXPOSE 8005
