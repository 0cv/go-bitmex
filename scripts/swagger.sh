#!/usr/bin/env bash

wget -q https://raw.githubusercontent.com/BitMEX/api-connectors/master/swagger.json -O raw.json
sed '/"default": {}/d' ./raw.json > raw2.json
sed '/"application\/json"/d' ./raw2.json > raw3.json
cat raw3.json | jq 'del(.security)' > swagger.json

rm -Rf ./swagger
mkdir -p ./swagger
swagger generate client -f ./swagger.json -t ./swagger

rm *.json