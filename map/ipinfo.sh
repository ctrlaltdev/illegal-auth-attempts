#!/usr/bin/env bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
  DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  $SOURCE
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE"
done
DIR="$( cd -P "$( dirname "$SOURCE" )" >/dev/null && pwd )"

. $DIR/.env

head -1000 $DIR/../src/unique.IPs.count.log > $DIR/IPs.tmp
IPs="$( cat $DIR/IPs.tmp )"

echo -e "[" > $DIR/ip.loc.json

for line in $IPs
do
  IP="$( echo $line | grep -oP '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?=:\d+)' )"
  count="$( echo $line | grep -oP '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\K\d+' )"

  loc="$( curl -s -H "Authorization: Bearer $IPINFO_TOKEN" https://ipinfo.io/$IP/loc )"

  echo -e "\t{\"IP\":\"$IP\", \"count\": \"$count\", \"loc\": \"$loc\"}," >> $DIR/ip.loc.json
done

echo -e "]" >> $DIR/ip.loc.json

rm $DIR/IPs.tmp

echo DONE
