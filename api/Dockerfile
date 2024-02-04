# Compile stage
# bullseye is needed, otherwise glibc error on start
FROM golang:1.21-bullseye AS build-env

ADD ../.. /app
WORKDIR /app

RUN go build -o /server
RUN chmod a+x /server
# RUN ls -la /server

# Final stage
FROM  debian:buster
EXPOSE 80
WORKDIR /

COPY --from=build-env /server /
# CMD ["/server"]
ENTRYPOINT ["/server"]
