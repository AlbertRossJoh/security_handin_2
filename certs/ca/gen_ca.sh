#!/bin/sh
# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 \
  -days 365 \
  -nodes -keyout ./certs/ca/ca-key.pem \
  -out ./certs/ca/ca-cert.pem \
  -subj "/C=DK/L=Copenhagen/O=ITU/OU=Education/CN=*.itu.dk/emailAddress=alrj@itu.dk" >/dev/null 2>&1
