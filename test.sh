#!/bin/bash
echo $CLIENT_ID
echo $CLIENT_SECRET
curl --output output.txt -X POST \
  'https://api.kroger.com/v1/connect/oauth2/token' \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  --compressed \
  -u "$CLIENT_ID:$CLIENT_SECRET" \
  -d 'grant_type=client_credentials'
