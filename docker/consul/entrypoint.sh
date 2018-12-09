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

while IFS='' read -r line || [[ -n "$line" ]]; do
    consul kv put "$line"
done < "/consul-initial/.consul"

tail -f /dev/null
