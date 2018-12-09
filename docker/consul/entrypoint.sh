#!/bin/bash

consul agent -dev -client 0.0.0.0 > /var/log/consul.log &

consulUrl=http://127.0.0.1:8500/v1/agent/checks
response=$(curl --write-out %{http_code} --silent --output /dev/null ${consulUrl})
while [[ "$response" -ne "200" ]]
do
    echo ${response}
    sleep 0.25
    response=$(curl --write-out %{http_code} --silent --output /dev/null ${consulUrl})
done

sleep 1

#consul kv put teamtrack/db/host localhost
#consul kv put teamtrack/db/port 5432
#consul kv put teamtrack/db/name teamtrack_db
#consul kv put teamtrack/db/user root
#consul kv put teamtrack/db/pass 123456

for file in /consul-initial/*; do
    while IFS='' read -r line || [[ -n "$line" ]]; do
        consul kv put "$line"
    done < "$file"
done

tail -f /dev/null
