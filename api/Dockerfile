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
EXPOSE 8080
WORKDIR /

# Install wget
RUN apt-get update && apt-get install -y wget && rm -rf /var/lib/apt/lists/*

COPY --from=build-env /server /
# CMD ["/server"]
ENTRYPOINT ["/server"]
