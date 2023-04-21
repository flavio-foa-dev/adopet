FROM golang:1.20 as build
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=Linux  GOARCH=amd64 go buil -o server

FROM scratch
COPY --from=build /app/server /server
ENTRYPOINT [ "/server" ]

