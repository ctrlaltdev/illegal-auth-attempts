#!/bin/bash

python3 ./dedupe.py

rm src/unique.IPs.count.log
rm src/unique.users.count.log

while IFS= read -r line; do count=$(grep -c "$line" src/IPs.log); echo $line:$count >> src/unique.IPs.count.log; done < src/unique.IPs.log

while IFS= read -r line; do count=$(grep -Pc "\A$line$" src/users.log); echo $line:$count >> src/unique.users.count.log; done < src/unique.users.log

python3 ./sort.py

echo DONE