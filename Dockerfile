FROM golang:alpine as builder
RUN mkdir -p /go/src/github.com/sacquatella/slides-remark-server

WORKDIR /go/src/github.com/sacquatella/slides-remark-server
#WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache git
RUN go get -u github.com/kardianos/govendor
RUN go get -u github.com/jteeuwen/go-bindata/...
RUN go get github.com/rakyll/statik
RUN govendor sync
RUN statik -f -src=./public
RUN go-bindata -o templates.go templates
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o slides-server .

FROM scratch
LABEL org.label-schema.name="slides-server"
LABEL org.label-schema.description="Slide as Code Server (markdown)"
LABEL org.label-schema.vendor="s.acquatella@gmail.com"

COPY --from=builder /go/src/github.com/sacquatella/slides-remark-server/slides-server /app/
WORKDIR /app
ENV GIN_MODE="release"
CMD ["./slides-server"]