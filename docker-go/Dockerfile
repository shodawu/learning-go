FROM golang:1.20

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./cfg/* ./cfg/

COPY . .
RUN go build -v -o /usr/local/bin/app ./

CMD ["app"]