FROM golang:1.15 AS build-env
WORKDIR /go/src/hello/
COPY ./apps/golang/main.go /go/src/hello/
RUN ["go","build","-tags","netgo"]

FROM scratch
COPY --from=build-env /go/src/hello/hello hello
EXPOSE 8080
ENTRYPOINT ["./hello"]
