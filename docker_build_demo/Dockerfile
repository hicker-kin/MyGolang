## release: 1.03G alpine: 539MB
FROM golang:1.19-alpine

WORKDIR /app/
COPY . .

RUN go build main.go

EXPOSE 8081
ENTRYPOINT ["./main"]


#FROM golang:alpine
#WORKDIR $GOPATH/src/gin_docker
#ADD . ./
#ENV GO111MODULE=on
#ENV GOPROXY="https://goproxy.io"
#RUN go build -o gin_docker .
#EXPOSE 8080
#ENTRYPOINT  ["./gin_docker"]