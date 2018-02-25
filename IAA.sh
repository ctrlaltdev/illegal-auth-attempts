#!/bin/bash

rm src/*.log

usersfiles=$(find import/ | grep -P "users.log");
for file in $usersfiles
do
  cat $file >> src/users.log
done

IPsfiles=$(find import/ | grep -P "IPs.log");
for file in $IPsfiles
do
  cat $file >> src/IPs.log
done

python3 ./dedupe.py

while IFS= read -r line; do count=$(grep -c "$line" src/IPs.log); echo $line:$count >> src/unique.IPs.count.log; done < src/unique.IPs.log

while IFS= read -r line; do count=$(grep -Pc "\A$line$" src/users.log); echo $line:$count >> src/unique.users.count.log; done < src/unique.users.log

python3 ./sort.py

echo DONE