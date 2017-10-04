FROM golang

RUN mkdir /api

WORKDIR /api

CMD ["./heu-api"]