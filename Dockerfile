FROM golang:1.21 AS BUILDER

WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
# COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
 go build -o meuprimeirocrudgo . 

FROM golang:1.21-alpine AS RUNNER

RUN adduser -D gabriel

COPY --from=BUILDER /app/meuprimeirocrudgo /app/meuprimeirocrudgo

RUN chown -R gabriel:gabriel /app
RUN chmod +X /app/meuprimeirocrudgo

EXPOSE 8080

USER gabriel

CMD [ "/app/meuprimeirocrudgo" ]