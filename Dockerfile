FROM golang:1.13.0

RUN apt-get update && apt-get install -y

WORKDIR /app/backup-scheduler

COPY . .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o binary

CMD ["/app/backup-scheduler/binary"] 