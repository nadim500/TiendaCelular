FROM golang

RUN mkdir -p /go/src/backapp

WORKDIR /go/src/backapp

COPY . /go/src/backapp

#Installar las dependencias
# go get install xx
RUN go get github.com/lib/pq

EXPOSE 8080

CMD ["go", "run","main.go"]
