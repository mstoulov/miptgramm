FROM golang:1.21-alpine AS build
RUN apk --no-cache add gcc g++ make git

WORKDIR /go/src/backend

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV ARCH=amd64
ENV CGO_CFLAGS="-D_LARGEFILE64_SOURCE"

COPY . .
RUN go mod download

RUN go build -o ./bin/backend ./cmd/app/main.go

FROM alpine AS runtime
WORKDIR /

COPY --from=build /go/src/backend/bin/backend /go/bin/

EXPOSE 8080
ENTRYPOINT /go/bin/backend

