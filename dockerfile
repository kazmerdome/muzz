FROM golang:alpine3.20 as dev
WORKDIR /
RUN apk add --update make
EXPOSE 4444
COPY . /
ENV CGO_ENABLED 0
RUN make build-gateway

FROM golang:alpine3.20 as prod
RUN apk --no-cache add ca-certificates
WORKDIR /run/
COPY --from=dev /build/gateway .
CMD ["./gateway"]
