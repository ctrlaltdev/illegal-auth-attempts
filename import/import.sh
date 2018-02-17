#!/bin/bash

while IFS= read -r line; do grep -oP 'Invalid user \K\w+(?= from)' <<< "$line" >> ../src/users.log; done < auth.log

while IFS= read -r line; do grep -oP '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?=.*preauth)' <<< "$line" >> ../src/IPs.log; done < auth.log
