FROM golang:1.20.2-alpine3.17 


WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download

RUN go mod tidy

# RUN go get github.com/canthefason/go-watcher/cmd/watcher

# RUN go install github.com/canthefason/go-watcher/cmd/watcher

# RUN install github.com/githubnemo/CompileDaemon

# RUN go build -o /go/src/app/main .

#Run the command to start the watcher
# CMD ["watcher", "-run", "main.go"]

# ENTRYPOINT /go/src/app 

# ENTRYPOINT CompileDaemon -build="go build -o main ." -command="./main"

# EXPOSE 5000

# CMD ["/go/src/app/main"]