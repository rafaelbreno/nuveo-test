FROM golang:1.18-alpine AS build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/app services/writer/main.go

FROM golang:1.18-alpine

COPY --from=build /go/src/app/bin/app .

EXPOSE 8070

CMD [ "./app", "-env=false" ]
