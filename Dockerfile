FROM alpine:3.8

EXPOSE 8080

ENV REMOTE=https://www.google.com:443/ IGNORE=true

RUN apk add --no-cache ca-certificates
COPY striptls /

CMD /striptls -l=:8080 -r=$REMOTE -i=$IGNORE
