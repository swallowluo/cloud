FROM golang:1.18.1 AS build
WORKDIR /home/ubuntu/cncamp/homework/c2/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver httpserver.go


FROM busybox
COPY --from=build /home/ubuntu/cncamp/homework/c2/httpserver /bin/httpserver
ENV VERSION=my.httpserver.ver.1.0.0
ENTRYPOINT ["/bin/httpserver"]


