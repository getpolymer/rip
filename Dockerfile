FROM golang:1.12-alpine as build

WORKDIR /app
COPY . /app

RUN go mod download
RUN go mod verify

RUN go build -o rip .

FROM alpine

COPY --from=build /app/rip /

ENTRYPOINT [ "./rip" ]
