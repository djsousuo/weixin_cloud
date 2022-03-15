FROM centos:7.4.1708

WORKDIR /app

ADD c main

CMD ["/app/main"]
