FROM golang:1.23
LABEL authors="MielPops"
ENV GOPATH=/

COPY ./ ./

# install pql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh /bin/sh

# build app
RUN go mod download
RUN go build -o main ./cmd/app/main.go

CMD ["./main"]