FROM golang:1.19

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o count-registries

CMD [ "./count-registries" ]
