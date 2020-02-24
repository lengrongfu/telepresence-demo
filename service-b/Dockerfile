FROM golang:1.7.3 as build
WORKDIR /go/src/github.com/lengrongfu/telepresence-demo/service-b
COPY . .
#RUN cd /test/multi-stage-demo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# ARG workspace
# WORKDIR /root/
# COPY --from=build /Users/lengrongfu/docker/workspace/multi-stage-demo/main .
# CMD ["./main"]

FROM alpine:latest

COPY --from=build /go/src/github.com/lengrongfu/telepresence-demo/service-b/main \
                 /go/src/github.com/lengrongfu/telepresence-demo/service-b/config.json \
                /root/
EXPOSE 8000

ENTRYPOINT [ "/root/main","-conf=/root/config.json" ]
