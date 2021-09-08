FROM golang:alpine as dev

EXPOSE 8080
WORKDIR go/src/github.com/nooderg/pipiSpot/
COPY ../go.mod .
COPY ../go.sum .
RUN go mod download -x
COPY .. .

CMD ["go", "run", "cmd/main.go"]