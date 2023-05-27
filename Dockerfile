FROM golang:1.18-alpine
ENV GO111MODULE=on

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add git

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# COPY *.go ./

RUN go build -o /main .
EXPOSE 3000
CMD ["/main"]