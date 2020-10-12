FROM golang:1.13
RUN go get -u github.com/gorilla/mux
RUN mkdir /app
ENV ROUTE :8081
ADD . /app
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]
