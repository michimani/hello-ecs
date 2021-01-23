FROM golang:1.15.5-alpine3.12 as build

ADD . .
RUN go build -o /main

FROM golang:1.15.5-alpine3.12

COPY --from=build /main /main
ENTRYPOINT [ "/main" ]