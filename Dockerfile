from golang:1.15.3-alpine3.12

WORKDIR /go/pkg/mod/quantime.ai/scheduling
COPY . /go/pkg/mod/quantime.ai/scheduling/
EXPOSE 8000

RUN go mod vendor
RUN env GOOS=linux go build -o out/q lib/pq/priorityqueue.go
CMD out/qt
