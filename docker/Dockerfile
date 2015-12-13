FROM busybox:ubuntu-14.04

ENV GODSON_VERSION 0.0.3

ADD https://github.com/tsaikd/godson/releases/download/${GODSON_VERSION}/godson-Linux-x86_64 /usr/local/bin/godson

RUN chmod +x /usr/local/bin/godson

CMD ["godson"]
