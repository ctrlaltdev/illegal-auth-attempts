#!/bin/bash

grep "$(date -d yesterday "+%b %e")" /var/log/auth.log > /tmp/auth.log.tmp

while IFS= read -r line; do grep -oP 'Invalid user \K\w+(?= from)' <<< "$line" >> ../import/${USER}_${HOSTNAME}.users.log; done < /tmp/auth.log.tmp

while IFS= read -r line; do grep -oP '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?=.*preauth)' <<< "$line" >> ../import/${USER}_${HOSTNAME}IPs.log; done < /tmp/auth.log.tmp