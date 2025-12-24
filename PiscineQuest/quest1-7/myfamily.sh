#!/bin/bash

if [ -z "$HERO_ID" ]; then
  echo "Error: HERO_ID is not set."
  exit 1
fi

curl -s "https://01.tomorrow-school.ai/assets/superhero/all.json" |
jq -r --arg id "$HERO_ID" '
  .[] | select(.id == ($id | tonumber)) |
  .connections.relatives
' |
sed 's|"||g' |
sed $'s/\\n/\n/g'

