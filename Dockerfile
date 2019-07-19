FROM golang:1.12 as builder 
LABEL maintainer="Sang Li <sang.lx@teko.vn>"
WORKDIR /app

#### Cache Vendor ... 
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/opa-server .


######## Start a new stage from scratch #######
FROM alpine:latest
WORKDIR /opa-server-app
COPY rbac.rego .
COPY --from=builder /usr/local/bin/opa-server .


EXPOSE 3000

CMD [ "./opa-server" ]
