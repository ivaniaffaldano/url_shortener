FROM golang:latest 

RUN mkdir /go/src/url_shortener
WORKDIR /go/src/url_shortener

RUN go get github.com/joho/godotenv
RUN go get github.com/go-chi/chi
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/swaggo/http-swagger
RUN go get github.com/alecthomas/template
RUN go get github.com/go-chi/cors

ADD . /go/src/url_shortener

RUN go build

CMD go run url_shortener

EXPOSE 8080