FROM golang:1.21.4
ENV GO111MODULE=on
ENV TZ=Asia/Jakarta
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD ["go","run","/app/main.go"]