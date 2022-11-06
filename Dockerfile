FROM golang:latest
LABEL name devstack

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD [ "./main" ]