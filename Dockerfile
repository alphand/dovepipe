FROM alpine:latest

MAINTAINER Niko Darmawan <Darmawan.niko@gmail.com>

WORKDIR "/opt"

ADD .docker_build/dovepipe /opt/bin/dovepipe

CMD ["/opt/bin/dovepipe"]
