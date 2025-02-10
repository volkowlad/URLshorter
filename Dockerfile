FROM golang:alpine
LABEL authors="MielPops"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd/app; go build -o /pivo .
EXPOSE 8080
CMD ["/pivo"]