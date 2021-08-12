ARG builder_image=golang:alpine
ARG base_image=alpine

FROM ${builder_image} AS builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ../check ./cmd/check
RUN go build -o ../in ./cmd/in
RUN go build -o ../out ./cmd/out

FROM ${base_image}

COPY --from=builder /check /opt/resource/check
COPY --from=builder /in /opt/resource/in
COPY --from=builder /out /opt/resource/out
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
