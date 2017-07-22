FROM golang:1.8
ADD main.go main.go
RUN go build --ldflags '-s -w -linkmode external -extldflags -static' -o /server

FROM scratch
COPY --from=0 /server /server
EXPOSE 8080
CMD ["/server"]
