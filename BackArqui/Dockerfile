FROM golang

RUN mkdir -p /go/src/backapp

WORKDIR /go/src/backapp

COPY . /go/src/backapp

ENV GOPATH /go/src/backapp

#Installar las dependencias
# go get install xx
RUN go get github.com/lib/pq
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/fatih/structs

EXPOSE 8080

CMD ["go", "run","main.go"]
