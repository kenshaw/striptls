FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY striptls /
ENV REMOTE=https://www.google.com:443/
EXPOSE 8080
ENTRYPOINT /striptls -r=$REMOTE -l=:8080
