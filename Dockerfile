FROM alpine:3.8

COPY bin/* /usr/local/bin/

RUN /usr/local/bin/api -log.alsoStdout=true -log.name=peanuts