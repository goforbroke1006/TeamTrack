#!/usr/bin/sh

set -e

consul agent -dev -client 0.0.0.0 > /var/log/consul.log &

response=$(curl --write-out %{http_code} --silent --output /dev/null http://127.0.0.1:8500/v1/agent/checks)
echo $response

#curl http://127.0.0.1:8500/v1/agent/checks

sleep 10

consul kv put teamtrack/db/host localhost
consul kv put teamtrack/db/port 5432
consul kv put teamtrack/db/name teamtrack_db
consul kv put teamtrack/db/user root
consul kv put teamtrack/db/pass 123456

tail -f /dev/null
