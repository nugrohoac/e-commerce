FROM golang:1.17-alpine as build

WORKDIR /app

COPY . .

# download dependencies
RUN go mod vendor

# build binary
RUN go build -o e-commerse main.go

#
EXPOSE 9000

CMD [ "./e-commerse rest" ]