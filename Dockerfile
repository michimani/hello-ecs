FROM golang:1.15.5-alpine3.12

COPY . /app
RUN cd /app \
&& go build

ENTRYPOINT [ "/app/hello-ecs" ]