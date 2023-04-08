FROM golang:1.20

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /mistrzownia-converter-stats-api ./cmd/main.go

CMD [ "/mistrzownia-converter-stats-api" ]