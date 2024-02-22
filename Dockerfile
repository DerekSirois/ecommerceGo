FROM alpine:latest

RUN mkdir /app

COPY ecommerceGo /app
COPY public /public

CMD ["/app/ecommerceGo"]