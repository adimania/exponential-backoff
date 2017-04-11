FROM alpine

RUN apk update && apk add ca-certificates # since we will call a HTTPS URL.
ADD https://github.com/adimania/exponential-backoff/releases/download/0.1/exponential /usr/local/bin/exponential
RUN chmod +x /usr/local/bin/exponential
CMD /usr/local/bin/exponential