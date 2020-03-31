FROM alpine:latest
RUN apk --update add ca-certificates
ARG CERT_FILE=/etc/ssl/certs/ca-certificates.crt

FROM alpine:latest
COPY --from=0 $CERT_FILE $CERT_FILE
ADD /tezos-tests /
ENTRYPOINT ["/tezos-tests"]
