FROM golang:1.20

LABEL maintainer="cIzEgnSFXZERhT"
RUN groupadd -r -g 1600 cIzEgnSFXZERhT
RUN useradd -r -g 1600 -u 1500 cIzEgnSFXZERhT

RUN chsh -s /usr/sbin/nologin root

WORKDIR /app
COPY --chown=1500:1600 . ./
RUN chown -R 1500:1600 /app

RUN go mod download

RUN go build -o /mistrzownia-converter-stats-api ./cmd/main.go

USER mistrzownia-converter-stats-api-user

CMD [ "/mistrzownia-converter-stats-api" ]
