FROM golang:1.21

ARG GOOGLE_APPLICATION_CREDENTIALS_ENCODED

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN echo "$GOOGLE_APPLICATION_CREDENTIALS_ENCODED" > firebase-admin.base64
RUN base64 -d firebase-admin.base64 > firebase-admin.json
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/firebase-admin.json

RUN CGO_ENABLED=0 GOOS=linux go build -o /mandarin

RUN useradd --shell /bin/bash go
RUN chown -R go:go .
USER go

CMD ["/mandarin"]
