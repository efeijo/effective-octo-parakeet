#!/bin/bash

uuid=$(uuidgen)

curl \
    --data "$(jq --null-input --arg id "$uuid" --arg text "sender sends to the receiver and stores message in a redis set" '{"id": $id, "text": $text}')" \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/send