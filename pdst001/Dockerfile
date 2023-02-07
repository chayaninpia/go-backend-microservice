# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

RUN apk update && apk add --no-cache build-base git
ADD . /app
WORKDIR /app

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /pdst001

EXPOSE 8001

FROM alpine

COPY conf /app/conf
WORKDIR /app
COPY --from=build /pdst001 /app/pdst001 
RUN apk update && apk add nano 

CMD [ "/app/pdst001" ]
