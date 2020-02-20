FROM golang:alpine as base
WORKDIR /app
COPY . .
RUN go build 

EXPOSE 8080
CMD ["./external-service"] 