#!/bin/bash

grep "$(date -d yesterday "+%b %e")" /var/log/auth.log > /tmp/auth.log.tmp

if [ ! -d "/var/log/illegal-attempts" ]
then
  mkdir /var/log/illegal-attempts
fi

while IFS= read -r line; do grep -oP 'Invalid user \K\w+(?= from)' <<< "$line" >> /var/log/illegal-attempts/users.log; done < /tmp/auth.log.tmp

while IFS= read -r line; do grep -oP '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?=.*preauth)' <<< "$line" >> /var/log/illegal-attempts/IPs.log; done < /tmp/auth.log.tmp