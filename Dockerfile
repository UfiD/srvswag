FROM golang:1.24.2-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN apk add --no-cache make

RUN go build main.go

FROM golang:1.24.2-alpine AS runner

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=build /build/main ./main

CMD ["/app/main" ]

