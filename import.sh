#!/bin/bash

sources=$(find import/sources/ | grep -P "auth.log|secure");

for source in $sources
do
  prefix=$(echo $source | grep -oP "sources/\K[\w\d-_]+");

  rm import/${prefix}.users.log;
  while IFS= read -r line; do grep -oP 'Invalid user \K\w+(?= from)' <<< "$line" >> import/${prefix}.users.log; done < $source

  rm import/${prefix}.IPs.log;
  while IFS= read -r line; do grep -oP '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?=.*preauth)' <<< "$line" >> import/${prefix}.IPs.log; done < $source
done

echo DONE