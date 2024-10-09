#!/bin/sh

rm certs/*.pem
rm -rf shared_volume/*
rm certs/ca/ca-cert.pem
rm certs/ca/ca-cert.srl
rm certs/ca/ca-key.pem

chmod +x certs/ca/gen_ca.sh
./certs/ca/gen_ca.sh

make generate

docker compose up --build
