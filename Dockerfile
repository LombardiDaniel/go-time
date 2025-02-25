# #### - DEV - ####
# FROM golang:1.22.3 AS dev

# WORKDIR /cmd

# COPY cmd/go.mod go.mod
# COPY cmd/go.sum go.sum
# RUN go mod download

# COPY cmd/ ./

# CMD ["go", "run", "."]

# #### - TESTS - ####
# FROM golang:1.22.3 AS tester

# WORKDIR /cmd

# COPY cmd/go.mod go.mod
# COPY cmd/go.sum go.sum
# RUN go mod download

# COPY cmd/ ./

# CMD ["go", "test", "-v", "./..."]


#### - BUILDER - ####
FROM golang:1.22.3 AS builder

WORKDIR /usr

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN go build -o /bin/main main.go


#### - SERVER - ####
FROM alpine:3.19.1 as server

RUN apk add --no-cache \
    gcompat=1.1.0-r4 \
    libstdc++=13.2.1_git20231014-r0

WORKDIR /srv

COPY --from=builder /bin/main ./main

RUN adduser --system --no-create-home nonroot
USER nonroot

ENV TERM=xterm

ENTRYPOINT [ "./main" ]
